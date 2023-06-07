package de

import (
	"fmt"

	"github.com/bandprotocol/chain/v2/cylinder"
	"github.com/bandprotocol/chain/v2/cylinder/client"
	"github.com/bandprotocol/chain/v2/cylinder/store"
	"github.com/bandprotocol/chain/v2/pkg/logger"
	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// DE is a worker responsible for generating own nonce (DE) of signing process
type DE struct {
	context *cylinder.Context
	logger  *logger.Logger
	client  *client.Client
	eventCh <-chan ctypes.ResultEvent
}

var _ cylinder.Worker = &DE{}

// New creates a new instance of the DE worker.
// It initializes the necessary components and returns the created DE instance or an error if initialization fails.
func New(ctx *cylinder.Context) (*DE, error) {
	cli, err := client.New(ctx.Config, ctx.Keyring)
	if err != nil {
		return nil, err
	}

	return &DE{
		context: ctx,
		logger:  ctx.Logger.With("worker", "DE"),
		client:  cli,
	}, nil
}

// subscribe subscribes to request_sign events and initializes the event channel for receiving events.
// It returns an error if the subscription fails.
func (de *DE) subscribe() (err error) {
	subscriptionQuery := fmt.Sprintf(
		"tm.event = 'Tx' AND %s.%s = '%s'",
		types.EventTypeRequestSign,
		types.AttributeKeyMember,
		de.context.Config.Granter,
	)
	de.eventCh, err = de.client.Subscribe("DE", subscriptionQuery, 1000)
	return
}

// updateDE updates DE if the remaining DE is too low.
func (de *DE) updateDE() {
	// Query DE information
	deRes, err := de.client.QueryDE(de.context.Config.Granter)
	if err != nil {
		de.logger.Error(":cold_sweat: Failed to query DE information: %s", err)
		return
	}

	// Check remaining DE, ignore if it's more than min-DE
	remaining := deRes.GetRemaining()
	if remaining >= de.context.Config.MinDE {
		return
	}

	// Log
	de.logger.Info(":delivery_truck: Updating DE")

	// Generate new DE pairs
	privDEs, pubDEs, err := generateDEPairs(de.context.Config.MinDE)
	if err != nil {
		de.logger.Error(":cold_sweat: Failed to generate new DE pairs: %s", err)
		return
	}

	// Store all DEs in the store
	for i, privDE := range privDEs {
		err := de.context.Store.SetDE(pubDEs[i], privDE)
		if err != nil {
			de.logger.Error(":cold_sweat: Failed to set new DE in the store: %s", err)
			return
		}
	}

	// Send MsgDE
	de.context.MsgCh <- &types.MsgSubmitDEs{
		DEs:    pubDEs,
		Member: de.context.Config.Granter,
	}
}

// Start starts the DE worker.
// It subscribes to events and starts processing incoming events.
func (de *DE) Start() {
	de.logger.Info("start")

	err := de.subscribe()
	if err != nil {
		de.context.ErrCh <- err
		return
	}

	// Update one time when starting worker first time.
	de.updateDE()

	for range de.eventCh {
		go de.updateDE()
	}
}

// Stop stops the DE worker.
func (de *DE) Stop() {
	de.logger.Info("stop")
	de.client.Stop()
}

// generateDEPairs generates n pairs of DE
func generateDEPairs(n uint64) (privDEs []store.DE, pubDEs []types.DE, err error) {
	for i := uint64(1); i <= n; i++ {
		de, err := tss.GenerateKeyPairs(2)
		if err != nil {
			return nil, nil, err
		}

		privDEs = append(privDEs, store.DE{
			PrivD: de[0].PrivateKey,
			PrivE: de[1].PrivateKey,
		})

		pubDEs = append(pubDEs, types.DE{
			PubD: de[0].PublicKey,
			PubE: de[1].PublicKey,
		})
	}

	return privDEs, pubDEs, nil
}
