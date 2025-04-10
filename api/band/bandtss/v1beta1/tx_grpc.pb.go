// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: band/bandtss/v1beta1/tx.proto

package bandtssv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_RequestSignature_FullMethodName     = "/band.bandtss.v1beta1.Msg/RequestSignature"
	Msg_Activate_FullMethodName             = "/band.bandtss.v1beta1.Msg/Activate"
	Msg_UpdateParams_FullMethodName         = "/band.bandtss.v1beta1.Msg/UpdateParams"
	Msg_TransitionGroup_FullMethodName      = "/band.bandtss.v1beta1.Msg/TransitionGroup"
	Msg_ForceTransitionGroup_FullMethodName = "/band.bandtss.v1beta1.Msg/ForceTransitionGroup"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// RequestSignature submits a general message to be signed by a specific group.
	RequestSignature(ctx context.Context, in *MsgRequestSignature, opts ...grpc.CallOption) (*MsgRequestSignatureResponse, error)
	// Activate activates the status of the sender.
	Activate(ctx context.Context, in *MsgActivate, opts ...grpc.CallOption) (*MsgActivateResponse, error)
	// UpdateParams updates the x/bandtss parameters.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// TransitionGroup creates a request for creating a new group and replacing current group.
	TransitionGroup(ctx context.Context, in *MsgTransitionGroup, opts ...grpc.CallOption) (*MsgTransitionGroupResponse, error)
	// ForceTransitionGroup sets the given group to the incoming group without the signature of a transition
	// message from a current group.
	ForceTransitionGroup(ctx context.Context, in *MsgForceTransitionGroup, opts ...grpc.CallOption) (*MsgForceTransitionGroupResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestSignature(ctx context.Context, in *MsgRequestSignature, opts ...grpc.CallOption) (*MsgRequestSignatureResponse, error) {
	out := new(MsgRequestSignatureResponse)
	err := c.cc.Invoke(ctx, Msg_RequestSignature_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Activate(ctx context.Context, in *MsgActivate, opts ...grpc.CallOption) (*MsgActivateResponse, error) {
	out := new(MsgActivateResponse)
	err := c.cc.Invoke(ctx, Msg_Activate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransitionGroup(ctx context.Context, in *MsgTransitionGroup, opts ...grpc.CallOption) (*MsgTransitionGroupResponse, error) {
	out := new(MsgTransitionGroupResponse)
	err := c.cc.Invoke(ctx, Msg_TransitionGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ForceTransitionGroup(ctx context.Context, in *MsgForceTransitionGroup, opts ...grpc.CallOption) (*MsgForceTransitionGroupResponse, error) {
	out := new(MsgForceTransitionGroupResponse)
	err := c.cc.Invoke(ctx, Msg_ForceTransitionGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// RequestSignature submits a general message to be signed by a specific group.
	RequestSignature(context.Context, *MsgRequestSignature) (*MsgRequestSignatureResponse, error)
	// Activate activates the status of the sender.
	Activate(context.Context, *MsgActivate) (*MsgActivateResponse, error)
	// UpdateParams updates the x/bandtss parameters.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// TransitionGroup creates a request for creating a new group and replacing current group.
	TransitionGroup(context.Context, *MsgTransitionGroup) (*MsgTransitionGroupResponse, error)
	// ForceTransitionGroup sets the given group to the incoming group without the signature of a transition
	// message from a current group.
	ForceTransitionGroup(context.Context, *MsgForceTransitionGroup) (*MsgForceTransitionGroupResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RequestSignature(context.Context, *MsgRequestSignature) (*MsgRequestSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSignature not implemented")
}
func (UnimplementedMsgServer) Activate(context.Context, *MsgActivate) (*MsgActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) TransitionGroup(context.Context, *MsgTransitionGroup) (*MsgTransitionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransitionGroup not implemented")
}
func (UnimplementedMsgServer) ForceTransitionGroup(context.Context, *MsgForceTransitionGroup) (*MsgForceTransitionGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceTransitionGroup not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_RequestSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestSignature)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestSignature_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestSignature(ctx, req.(*MsgRequestSignature))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgActivate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Activate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Activate(ctx, req.(*MsgActivate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransitionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransitionGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransitionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransitionGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransitionGroup(ctx, req.(*MsgTransitionGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ForceTransitionGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgForceTransitionGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ForceTransitionGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ForceTransitionGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ForceTransitionGroup(ctx, req.(*MsgForceTransitionGroup))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "band.bandtss.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestSignature",
			Handler:    _Msg_RequestSignature_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _Msg_Activate_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "TransitionGroup",
			Handler:    _Msg_TransitionGroup_Handler,
		},
		{
			MethodName: "ForceTransitionGroup",
			Handler:    _Msg_ForceTransitionGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/bandtss/v1beta1/tx.proto",
}
