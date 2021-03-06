/*
Copyright 2018 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
/*
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

/*
Package backend is a generated protocol buffer package.

It is generated from these files:
	backend.proto

It has these top-level messages:
	Profile
	MatchObject
	Result
	Roster
	ConnectionInfo
	Assignments
*/
package backend

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Data structure for a profile to pass to the matchmaking function.
type Profile struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties string `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

// Data structure for all the properties of a match.
type MatchObject struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Properties string `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *MatchObject) Reset()                    { *m = MatchObject{} }
func (m *MatchObject) String() string            { return proto.CompactTextString(m) }
func (*MatchObject) ProtoMessage()               {}
func (*MatchObject) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MatchObject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MatchObject) GetProperties() string {
	if m != nil {
		return m.Properties
	}
	return ""
}

// Simple message to return success/failure and error status.
type Result struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// Data structure to hold a list of players in a match.
type Roster struct {
	PlayerIds string `protobuf:"bytes,1,opt,name=player_ids,json=playerIds" json:"player_ids,omitempty"`
}

func (m *Roster) Reset()                    { *m = Roster{} }
func (m *Roster) String() string            { return proto.CompactTextString(m) }
func (*Roster) ProtoMessage()               {}
func (*Roster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Roster) GetPlayerIds() string {
	if m != nil {
		return m.PlayerIds
	}
	return ""
}

// Simple message used to pass the connection string for the DGS to the player.
type ConnectionInfo struct {
	ConnectionString string `protobuf:"bytes,1,opt,name=connection_string,json=connectionString" json:"connection_string,omitempty"`
}

func (m *ConnectionInfo) Reset()                    { *m = ConnectionInfo{} }
func (m *ConnectionInfo) String() string            { return proto.CompactTextString(m) }
func (*ConnectionInfo) ProtoMessage()               {}
func (*ConnectionInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConnectionInfo) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

type Assignments struct {
	Roster         *Roster         `protobuf:"bytes,1,opt,name=roster" json:"roster,omitempty"`
	ConnectionInfo *ConnectionInfo `protobuf:"bytes,2,opt,name=connection_info,json=connectionInfo" json:"connection_info,omitempty"`
}

func (m *Assignments) Reset()                    { *m = Assignments{} }
func (m *Assignments) String() string            { return proto.CompactTextString(m) }
func (*Assignments) ProtoMessage()               {}
func (*Assignments) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Assignments) GetRoster() *Roster {
	if m != nil {
		return m.Roster
	}
	return nil
}

func (m *Assignments) GetConnectionInfo() *ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*MatchObject)(nil), "MatchObject")
	proto.RegisterType((*Result)(nil), "Result")
	proto.RegisterType((*Roster)(nil), "Roster")
	proto.RegisterType((*ConnectionInfo)(nil), "ConnectionInfo")
	proto.RegisterType((*Assignments)(nil), "Assignments")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	// Calls to ask the matchmaker to run a matchmaking function.
	//
	// Run MMF once.  Return a matchobject that fits this profile.
	CreateMatch(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*MatchObject, error)
	// Continually run MMF and stream matchobjects that fit this profile until
	// client closes the connection.
	ListMatches(ctx context.Context, in *Profile, opts ...grpc.CallOption) (API_ListMatchesClient, error)
	// Delete a matchobject from state storage manually. (Matchobjects in state
	// storage will also automatically expire after a while)
	DeleteMatch(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error)
	// Call that manage communication of  DGS connection info to players.
	//
	// Write the DGS connection info for the list of players in the
	// Assignments.roster to state storage, so that info can be read by the game
	// client(s).
	CreateAssignments(ctx context.Context, in *Assignments, opts ...grpc.CallOption) (*Result, error)
	// Remove DGS connection info for the list of players in the Roster from
	// state storage.
	DeleteAssignments(ctx context.Context, in *Roster, opts ...grpc.CallOption) (*Result, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) CreateMatch(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*MatchObject, error) {
	out := new(MatchObject)
	err := grpc.Invoke(ctx, "/API/CreateMatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) ListMatches(ctx context.Context, in *Profile, opts ...grpc.CallOption) (API_ListMatchesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_API_serviceDesc.Streams[0], c.cc, "/API/ListMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIListMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_ListMatchesClient interface {
	Recv() (*MatchObject, error)
	grpc.ClientStream
}

type aPIListMatchesClient struct {
	grpc.ClientStream
}

func (x *aPIListMatchesClient) Recv() (*MatchObject, error) {
	m := new(MatchObject)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) DeleteMatch(ctx context.Context, in *MatchObject, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/API/DeleteMatch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) CreateAssignments(ctx context.Context, in *Assignments, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/API/CreateAssignments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) DeleteAssignments(ctx context.Context, in *Roster, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/API/DeleteAssignments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	// Calls to ask the matchmaker to run a matchmaking function.
	//
	// Run MMF once.  Return a matchobject that fits this profile.
	CreateMatch(context.Context, *Profile) (*MatchObject, error)
	// Continually run MMF and stream matchobjects that fit this profile until
	// client closes the connection.
	ListMatches(*Profile, API_ListMatchesServer) error
	// Delete a matchobject from state storage manually. (Matchobjects in state
	// storage will also automatically expire after a while)
	DeleteMatch(context.Context, *MatchObject) (*Result, error)
	// Call that manage communication of  DGS connection info to players.
	//
	// Write the DGS connection info for the list of players in the
	// Assignments.roster to state storage, so that info can be read by the game
	// client(s).
	CreateAssignments(context.Context, *Assignments) (*Result, error)
	// Remove DGS connection info for the list of players in the Roster from
	// state storage.
	DeleteAssignments(context.Context, *Roster) (*Result, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_CreateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/CreateMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateMatch(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_ListMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Profile)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).ListMatches(m, &aPIListMatchesServer{stream})
}

type API_ListMatchesServer interface {
	Send(*MatchObject) error
	grpc.ServerStream
}

type aPIListMatchesServer struct {
	grpc.ServerStream
}

func (x *aPIListMatchesServer) Send(m *MatchObject) error {
	return x.ServerStream.SendMsg(m)
}

func _API_DeleteMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/DeleteMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteMatch(ctx, req.(*MatchObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_CreateAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Assignments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).CreateAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/CreateAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).CreateAssignments(ctx, req.(*Assignments))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_DeleteAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Roster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).DeleteAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/DeleteAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).DeleteAssignments(ctx, req.(*Roster))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatch",
			Handler:    _API_CreateMatch_Handler,
		},
		{
			MethodName: "DeleteMatch",
			Handler:    _API_DeleteMatch_Handler,
		},
		{
			MethodName: "CreateAssignments",
			Handler:    _API_CreateAssignments_Handler,
		},
		{
			MethodName: "DeleteAssignments",
			Handler:    _API_DeleteAssignments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMatches",
			Handler:       _API_ListMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "backend.proto",
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x29, 0xc6, 0x16, 0x66, 0x11, 0x64, 0xe3, 0x81, 0x90, 0xf8, 0x27, 0x3d, 0x88, 0x46,
	0xb3, 0x31, 0x78, 0xc1, 0x83, 0x07, 0x82, 0x17, 0x12, 0x8d, 0xa4, 0x3e, 0x00, 0x29, 0xdb, 0x01,
	0x56, 0xeb, 0x6e, 0xb3, 0xbb, 0x1c, 0x7c, 0x53, 0x1f, 0xc7, 0xb8, 0x2d, 0xba, 0x1c, 0x3c, 0x78,
	0x9c, 0x5f, 0xbf, 0x6f, 0xe6, 0xeb, 0xcc, 0xc2, 0xc1, 0x22, 0xe5, 0x6f, 0x28, 0x33, 0x56, 0x68,
	0x65, 0x55, 0x7c, 0x07, 0xd1, 0x4c, 0xab, 0xa5, 0xc8, 0x91, 0xb6, 0xa1, 0x2e, 0xb2, 0x5e, 0x70,
	0x16, 0x5c, 0x34, 0x93, 0xba, 0xc8, 0xe8, 0x09, 0x40, 0xa1, 0x55, 0x81, 0xda, 0x0a, 0x34, 0xbd,
	0xba, 0xe3, 0x1e, 0x89, 0xef, 0x81, 0x3c, 0xa5, 0x96, 0xaf, 0x9f, 0x17, 0xaf, 0xc8, 0xed, 0xbf,
	0xed, 0x23, 0x08, 0x13, 0x34, 0x9b, 0xdc, 0xd2, 0x1e, 0x44, 0x66, 0xc3, 0x39, 0x1a, 0xe3, 0xec,
	0x8d, 0x64, 0x5b, 0xd2, 0x23, 0xd8, 0x47, 0xad, 0x95, 0xae, 0xec, 0x65, 0x11, 0x0f, 0x20, 0x4c,
	0x94, 0xb1, 0xa8, 0xe9, 0x31, 0x40, 0x91, 0xa7, 0x1f, 0xa8, 0xe7, 0x22, 0x33, 0xd5, 0xec, 0x66,
	0x49, 0xa6, 0xd9, 0x77, 0xc2, 0xf6, 0x44, 0x49, 0x89, 0xdc, 0x0a, 0x25, 0xa7, 0x72, 0xa9, 0xe8,
	0x15, 0x74, 0xf9, 0x0f, 0x99, 0x1b, 0xab, 0x85, 0x5c, 0x55, 0xbe, 0xc3, 0xdf, 0x0f, 0x2f, 0x8e,
	0xc7, 0x6b, 0x20, 0x63, 0x63, 0xc4, 0x4a, 0xbe, 0xa3, 0xb4, 0x86, 0x9e, 0x42, 0xa8, 0xdd, 0x58,
	0x67, 0x20, 0xc3, 0x88, 0x95, 0x29, 0x92, 0x0a, 0xd3, 0x11, 0x74, 0xbc, 0xe6, 0x42, 0x2e, 0x95,
	0xcb, 0x4d, 0x86, 0x1d, 0xb6, 0x1b, 0x23, 0x69, 0xf3, 0x9d, 0x7a, 0xf8, 0x19, 0xc0, 0xde, 0x78,
	0x36, 0xa5, 0x03, 0x20, 0x13, 0x8d, 0xa9, 0x45, 0xb7, 0x58, 0xda, 0x60, 0xd5, 0x6d, 0xfa, 0x2d,
	0xe6, 0xad, 0x3a, 0xae, 0xd1, 0x4b, 0x20, 0x8f, 0xc2, 0x58, 0x07, 0xd1, 0xfc, 0x2d, 0xbc, 0x09,
	0xe8, 0x39, 0x90, 0x07, 0xcc, 0x71, 0xdb, 0x73, 0x47, 0xd0, 0x8f, 0x58, 0x79, 0x83, 0xb8, 0x46,
	0xaf, 0xa1, 0x5b, 0xce, 0xf6, 0xff, 0xb9, 0xc5, 0xbc, 0xca, 0x57, 0x0f, 0xa0, 0x5b, 0x76, 0xf5,
	0xd5, 0xdb, 0x8d, 0x78, 0xc2, 0x45, 0xe8, 0xde, 0xd9, 0xed, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0x23, 0x14, 0x36, 0x78, 0x02, 0x00, 0x00,
}
