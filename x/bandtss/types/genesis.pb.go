// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: band/bandtss/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the bandtss module's genesis state.
type GenesisState struct {
	// params defines all the paramiters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// members is an array containing members information.
	Members []Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members"`
	// current_group is the current group information.
	CurrentGroup CurrentGroup `protobuf:"bytes,3,opt,name=current_group,json=currentGroup,proto3" json:"current_group"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fe6a4345855d3c9, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetMembers() []Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GenesisState) GetCurrentGroup() CurrentGroup {
	if m != nil {
		return m.CurrentGroup
	}
	return CurrentGroup{}
}

// Params defines the set of module parameters.
type Params struct {
	// active_duration is the duration where a member is active without interaction.
	ActiveDuration time.Duration `protobuf:"bytes,1,opt,name=active_duration,json=activeDuration,proto3,stdduration" json:"active_duration"`
	// reward_percentage is the percentage of block rewards allocated to active TSS members.
	// The reward proportion is calculated after being allocated to oracle rewards.
	RewardPercentage uint64 `protobuf:"varint,2,opt,name=reward_percentage,json=rewardPercentage,proto3" json:"reward_percentage,omitempty"`
	// inactive_penalty_duration is the duration where a member cannot activate back after being set to inactive.
	InactivePenaltyDuration time.Duration `protobuf:"bytes,3,opt,name=inactive_penalty_duration,json=inactivePenaltyDuration,proto3,stdduration" json:"inactive_penalty_duration"`
	// max_transition_duration is the maximum duration where the transition process waits
	// since the start of the process until an incoming group replaces a current group.
	MaxTransitionDuration time.Duration `protobuf:"bytes,4,opt,name=max_transition_duration,json=maxTransitionDuration,proto3,stdduration" json:"max_transition_duration"`
	// fee is the tokens that will be paid per signer.
	Fee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fe6a4345855d3c9, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetActiveDuration() time.Duration {
	if m != nil {
		return m.ActiveDuration
	}
	return 0
}

func (m *Params) GetRewardPercentage() uint64 {
	if m != nil {
		return m.RewardPercentage
	}
	return 0
}

func (m *Params) GetInactivePenaltyDuration() time.Duration {
	if m != nil {
		return m.InactivePenaltyDuration
	}
	return 0
}

func (m *Params) GetMaxTransitionDuration() time.Duration {
	if m != nil {
		return m.MaxTransitionDuration
	}
	return 0
}

func (m *Params) GetFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fee
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "band.bandtss.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "band.bandtss.v1beta1.Params")
}

func init() {
	proto.RegisterFile("band/bandtss/v1beta1/genesis.proto", fileDescriptor_3fe6a4345855d3c9)
}

var fileDescriptor_3fe6a4345855d3c9 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0x3a, 0xff, 0xfc, 0xd1, 0x50, 0xa0, 0x58, 0x41, 0x75, 0x2a, 0x64, 0x47, 0x59,
	0x65, 0xc3, 0x4c, 0x4b, 0x77, 0x15, 0x1b, 0x0c, 0x52, 0x25, 0x44, 0xa5, 0xc8, 0xb0, 0x02, 0x21,
	0x6b, 0xec, 0xdc, 0xba, 0x16, 0xf1, 0x8c, 0x35, 0x33, 0x0e, 0xe9, 0x5b, 0xb0, 0x64, 0xc9, 0x9a,
	0x67, 0xe0, 0x01, 0xba, 0xec, 0x0e, 0x56, 0x2d, 0x4a, 0x36, 0x3c, 0x06, 0x9a, 0xf1, 0x47, 0x2a,
	0x14, 0xa4, 0x6e, 0xfc, 0x31, 0xf7, 0x9c, 0xdf, 0x9c, 0xb9, 0x73, 0xd1, 0x28, 0xa6, 0x6c, 0x4a,
	0xf4, 0x43, 0x49, 0x49, 0xe6, 0x07, 0x31, 0x28, 0x7a, 0x40, 0x52, 0x60, 0x20, 0x33, 0x89, 0x0b,
	0xc1, 0x15, 0x77, 0xfa, 0xba, 0x8c, 0x6b, 0x0d, 0xae, 0x35, 0x7b, 0xfd, 0x94, 0xa7, 0xdc, 0x08,
	0x88, 0xfe, 0xaa, 0xb4, 0x7b, 0x5e, 0xca, 0x79, 0x3a, 0x03, 0x62, 0xfe, 0xe2, 0xf2, 0x94, 0x4c,
	0x4b, 0x41, 0x55, 0xc6, 0x59, 0x53, 0x4f, 0xb8, 0xcc, 0xb9, 0x24, 0x31, 0x95, 0xd0, 0x6e, 0x97,
	0xf0, 0xac, 0xa9, 0x6f, 0xce, 0xd3, 0xec, 0x6d, 0x34, 0xa3, 0x1f, 0x16, 0xda, 0x3e, 0xae, 0x12,
	0xbe, 0x51, 0x54, 0x81, 0x73, 0x84, 0x7a, 0x05, 0x15, 0x34, 0x97, 0xae, 0x35, 0xb4, 0xc6, 0x77,
	0x9f, 0x3e, 0xc6, 0x9b, 0x12, 0xe3, 0x89, 0xd1, 0x04, 0xdd, 0x8b, 0x2b, 0xbf, 0x13, 0xd6, 0x0e,
	0xe7, 0x19, 0xfa, 0x3f, 0x87, 0x3c, 0x06, 0x21, 0xdd, 0xad, 0xa1, 0xfd, 0x6f, 0xf3, 0x89, 0x11,
	0xd5, 0xe6, 0xc6, 0xe2, 0x9c, 0xa0, 0x7b, 0x49, 0x29, 0x04, 0x30, 0x15, 0xa5, 0x82, 0x97, 0x85,
	0x6b, 0x9b, 0x00, 0xa3, 0xcd, 0x8c, 0x17, 0x95, 0xf4, 0x58, 0x2b, 0x6b, 0xd2, 0x76, 0x72, 0x63,
	0x6d, 0xf4, 0xdd, 0x46, 0xbd, 0x2a, 0xa5, 0xf3, 0x1a, 0x3d, 0xa0, 0x89, 0xca, 0xe6, 0x10, 0x35,
	0x1d, 0xac, 0x0f, 0x37, 0xc0, 0x55, 0x8b, 0x71, 0xd3, 0x62, 0xfc, 0xb2, 0x16, 0x04, 0x77, 0x34,
	0xf2, 0xcb, 0xb5, 0x6f, 0x85, 0xf7, 0x2b, 0x6f, 0x53, 0x71, 0x9e, 0xa3, 0x87, 0x02, 0x3e, 0x51,
	0x31, 0x8d, 0x0a, 0x10, 0x09, 0x30, 0x45, 0x53, 0x70, 0xb7, 0x86, 0xd6, 0xb8, 0x1b, 0xf4, 0x97,
	0x57, 0xfe, 0x4e, 0x68, 0x8a, 0x93, 0xb6, 0x16, 0xee, 0x88, 0xbf, 0x56, 0x9c, 0x08, 0x0d, 0x32,
	0x56, 0x47, 0x2a, 0x80, 0xd1, 0x99, 0x3a, 0x5f, 0x47, 0xb3, 0x6f, 0x1f, 0x6d, 0xb7, 0xa1, 0x4c,
	0x2a, 0x48, 0x9b, 0xf1, 0x3d, 0xda, 0xcd, 0xe9, 0x22, 0x52, 0x82, 0x32, 0x99, 0xe9, 0x95, 0x35,
	0xbe, 0x7b, 0x7b, 0xfc, 0xa3, 0x9c, 0x2e, 0xde, 0xb6, 0x88, 0x16, 0xfe, 0x01, 0xd9, 0xa7, 0x00,
	0xee, 0x7f, 0xe6, 0x8a, 0x07, 0xb8, 0x9a, 0x42, 0xac, 0xa7, 0x70, 0x7d, 0x3b, 0x3c, 0x63, 0xc1,
	0xbe, 0x06, 0x7d, 0xbb, 0xf6, 0xc7, 0x69, 0xa6, 0xce, 0xca, 0x18, 0x27, 0x3c, 0x27, 0xf5, 0xc8,
	0x56, 0xaf, 0x27, 0x72, 0xfa, 0x91, 0xa8, 0xf3, 0x02, 0xa4, 0x31, 0xc8, 0x50, 0x73, 0x8f, 0xba,
	0xbf, 0xbf, 0xfa, 0x56, 0xf0, 0xea, 0x62, 0xe9, 0x59, 0x97, 0x4b, 0xcf, 0xfa, 0xb5, 0xf4, 0xac,
	0xcf, 0x2b, 0xaf, 0x73, 0xb9, 0xf2, 0x3a, 0x3f, 0x57, 0x5e, 0xe7, 0xdd, 0xfe, 0x0d, 0x9c, 0x9e,
	0x0a, 0x73, 0x82, 0x84, 0xcf, 0x48, 0x72, 0x46, 0x33, 0x46, 0xe6, 0x87, 0x64, 0xd1, 0x0e, 0xbd,
	0x81, 0xc7, 0x3d, 0x23, 0x39, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xb4, 0x5a, 0x1f, 0xa1,
	0x03, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ActiveDuration != that1.ActiveDuration {
		return false
	}
	if this.RewardPercentage != that1.RewardPercentage {
		return false
	}
	if this.InactivePenaltyDuration != that1.InactivePenaltyDuration {
		return false
	}
	if this.MaxTransitionDuration != that1.MaxTransitionDuration {
		return false
	}
	if len(this.Fee) != len(that1.Fee) {
		return false
	}
	for i := range this.Fee {
		if !this.Fee[i].Equal(&that1.Fee[i]) {
			return false
		}
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CurrentGroup.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxTransitionDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxTransitionDuration):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.InactivePenaltyDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactivePenaltyDuration):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	if m.RewardPercentage != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RewardPercentage))
		i--
		dAtA[i] = 0x10
	}
	n5, err5 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.ActiveDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintGenesis(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CurrentGroup.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.ActiveDuration)
	n += 1 + l + sovGenesis(uint64(l))
	if m.RewardPercentage != 0 {
		n += 1 + sovGenesis(uint64(m.RewardPercentage))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.InactivePenaltyDuration)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxTransitionDuration)
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentGroup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.ActiveDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPercentage", wireType)
			}
			m.RewardPercentage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RewardPercentage |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InactivePenaltyDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.InactivePenaltyDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTransitionDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxTransitionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
