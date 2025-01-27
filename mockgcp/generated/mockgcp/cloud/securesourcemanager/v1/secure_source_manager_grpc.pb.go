// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/securesourcemanager/v1/secure_source_manager.proto

package securesourcemanagerpb

import (
	iampb "cloud.google.com/go/iam/apiv1/iampb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecureSourceManagerClient is the client API for SecureSourceManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecureSourceManagerClient interface {
	// Lists Instances in a given project and location.
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Gets details of a single instance.
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
	// Creates a new instance in a given project and location.
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a single instance.
	DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Lists Repositories in a given project and location.
	//
	// **Host: Data Plane**
	ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
	// Gets metadata of a repository.
	//
	// **Host: Data Plane**
	GetRepository(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*Repository, error)
	// Creates a new repository in a given project and location.
	//
	// **Host: Data Plane**
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a Repository.
	//
	// **Host: Data Plane**
	DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Get IAM policy for a repository.
	GetIamPolicyRepo(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Set IAM policy on a repository.
	SetIamPolicyRepo(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error)
	// Test IAM permissions on a repository.
	// IAM permission checks are not required on this method.
	TestIamPermissionsRepo(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error)
	// CreateBranchRule creates a branch rule in a given repository.
	CreateBranchRule(ctx context.Context, in *CreateBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// ListBranchRules lists branch rules in a given repository.
	ListBranchRules(ctx context.Context, in *ListBranchRulesRequest, opts ...grpc.CallOption) (*ListBranchRulesResponse, error)
	// GetBranchRule gets a branch rule.
	GetBranchRule(ctx context.Context, in *GetBranchRuleRequest, opts ...grpc.CallOption) (*BranchRule, error)
	// UpdateBranchRule updates a branch rule.
	UpdateBranchRule(ctx context.Context, in *UpdateBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// DeleteBranchRule deletes a branch rule.
	DeleteBranchRule(ctx context.Context, in *DeleteBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type secureSourceManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSecureSourceManagerClient(cc grpc.ClientConnInterface) SecureSourceManagerClient {
	return &secureSourceManagerClient{cc}
}

func (c *secureSourceManagerClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) DeleteInstance(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) GetRepository(ctx context.Context, in *GetRepositoryRequest, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) DeleteRepository(ctx context.Context, in *DeleteRepositoryRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) GetIamPolicyRepo(ctx context.Context, in *iampb.GetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetIamPolicyRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) SetIamPolicyRepo(ctx context.Context, in *iampb.SetIamPolicyRequest, opts ...grpc.CallOption) (*iampb.Policy, error) {
	out := new(iampb.Policy)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/SetIamPolicyRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) TestIamPermissionsRepo(ctx context.Context, in *iampb.TestIamPermissionsRequest, opts ...grpc.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	out := new(iampb.TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/TestIamPermissionsRepo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) CreateBranchRule(ctx context.Context, in *CreateBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateBranchRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) ListBranchRules(ctx context.Context, in *ListBranchRulesRequest, opts ...grpc.CallOption) (*ListBranchRulesResponse, error) {
	out := new(ListBranchRulesResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListBranchRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) GetBranchRule(ctx context.Context, in *GetBranchRuleRequest, opts ...grpc.CallOption) (*BranchRule, error) {
	out := new(BranchRule)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetBranchRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) UpdateBranchRule(ctx context.Context, in *UpdateBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/UpdateBranchRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secureSourceManagerClient) DeleteBranchRule(ctx context.Context, in *DeleteBranchRuleRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteBranchRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecureSourceManagerServer is the server API for SecureSourceManager service.
// All implementations must embed UnimplementedSecureSourceManagerServer
// for forward compatibility
type SecureSourceManagerServer interface {
	// Lists Instances in a given project and location.
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Gets details of a single instance.
	GetInstance(context.Context, *GetInstanceRequest) (*Instance, error)
	// Creates a new instance in a given project and location.
	CreateInstance(context.Context, *CreateInstanceRequest) (*longrunningpb.Operation, error)
	// Deletes a single instance.
	DeleteInstance(context.Context, *DeleteInstanceRequest) (*longrunningpb.Operation, error)
	// Lists Repositories in a given project and location.
	//
	// **Host: Data Plane**
	ListRepositories(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
	// Gets metadata of a repository.
	//
	// **Host: Data Plane**
	GetRepository(context.Context, *GetRepositoryRequest) (*Repository, error)
	// Creates a new repository in a given project and location.
	//
	// **Host: Data Plane**
	CreateRepository(context.Context, *CreateRepositoryRequest) (*longrunningpb.Operation, error)
	// Deletes a Repository.
	//
	// **Host: Data Plane**
	DeleteRepository(context.Context, *DeleteRepositoryRequest) (*longrunningpb.Operation, error)
	// Get IAM policy for a repository.
	GetIamPolicyRepo(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error)
	// Set IAM policy on a repository.
	SetIamPolicyRepo(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error)
	// Test IAM permissions on a repository.
	// IAM permission checks are not required on this method.
	TestIamPermissionsRepo(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error)
	// CreateBranchRule creates a branch rule in a given repository.
	CreateBranchRule(context.Context, *CreateBranchRuleRequest) (*longrunningpb.Operation, error)
	// ListBranchRules lists branch rules in a given repository.
	ListBranchRules(context.Context, *ListBranchRulesRequest) (*ListBranchRulesResponse, error)
	// GetBranchRule gets a branch rule.
	GetBranchRule(context.Context, *GetBranchRuleRequest) (*BranchRule, error)
	// UpdateBranchRule updates a branch rule.
	UpdateBranchRule(context.Context, *UpdateBranchRuleRequest) (*longrunningpb.Operation, error)
	// DeleteBranchRule deletes a branch rule.
	DeleteBranchRule(context.Context, *DeleteBranchRuleRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedSecureSourceManagerServer()
}

// UnimplementedSecureSourceManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSecureSourceManagerServer struct {
}

func (UnimplementedSecureSourceManagerServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedSecureSourceManagerServer) GetInstance(context.Context, *GetInstanceRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedSecureSourceManagerServer) CreateInstance(context.Context, *CreateInstanceRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedSecureSourceManagerServer) DeleteInstance(context.Context, *DeleteInstanceRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstance not implemented")
}
func (UnimplementedSecureSourceManagerServer) ListRepositories(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}
func (UnimplementedSecureSourceManagerServer) GetRepository(context.Context, *GetRepositoryRequest) (*Repository, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepository not implemented")
}
func (UnimplementedSecureSourceManagerServer) CreateRepository(context.Context, *CreateRepositoryRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}
func (UnimplementedSecureSourceManagerServer) DeleteRepository(context.Context, *DeleteRepositoryRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRepository not implemented")
}
func (UnimplementedSecureSourceManagerServer) GetIamPolicyRepo(context.Context, *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicyRepo not implemented")
}
func (UnimplementedSecureSourceManagerServer) SetIamPolicyRepo(context.Context, *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicyRepo not implemented")
}
func (UnimplementedSecureSourceManagerServer) TestIamPermissionsRepo(context.Context, *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissionsRepo not implemented")
}
func (UnimplementedSecureSourceManagerServer) CreateBranchRule(context.Context, *CreateBranchRuleRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranchRule not implemented")
}
func (UnimplementedSecureSourceManagerServer) ListBranchRules(context.Context, *ListBranchRulesRequest) (*ListBranchRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBranchRules not implemented")
}
func (UnimplementedSecureSourceManagerServer) GetBranchRule(context.Context, *GetBranchRuleRequest) (*BranchRule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBranchRule not implemented")
}
func (UnimplementedSecureSourceManagerServer) UpdateBranchRule(context.Context, *UpdateBranchRuleRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBranchRule not implemented")
}
func (UnimplementedSecureSourceManagerServer) DeleteBranchRule(context.Context, *DeleteBranchRuleRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBranchRule not implemented")
}
func (UnimplementedSecureSourceManagerServer) mustEmbedUnimplementedSecureSourceManagerServer() {}

// UnsafeSecureSourceManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecureSourceManagerServer will
// result in compilation errors.
type UnsafeSecureSourceManagerServer interface {
	mustEmbedUnimplementedSecureSourceManagerServer()
}

func RegisterSecureSourceManagerServer(s grpc.ServiceRegistrar, srv SecureSourceManagerServer) {
	s.RegisterService(&SecureSourceManager_ServiceDesc, srv)
}

func _SecureSourceManager_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_DeleteInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).DeleteInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).DeleteInstance(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).ListRepositories(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).GetRepository(ctx, req.(*GetRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_DeleteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).DeleteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).DeleteRepository(ctx, req.(*DeleteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_GetIamPolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).GetIamPolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetIamPolicyRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).GetIamPolicyRepo(ctx, req.(*iampb.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_SetIamPolicyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).SetIamPolicyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/SetIamPolicyRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).SetIamPolicyRepo(ctx, req.(*iampb.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_TestIamPermissionsRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iampb.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).TestIamPermissionsRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/TestIamPermissionsRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).TestIamPermissionsRepo(ctx, req.(*iampb.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_CreateBranchRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranchRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).CreateBranchRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/CreateBranchRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).CreateBranchRule(ctx, req.(*CreateBranchRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_ListBranchRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBranchRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).ListBranchRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/ListBranchRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).ListBranchRules(ctx, req.(*ListBranchRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_GetBranchRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBranchRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).GetBranchRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/GetBranchRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).GetBranchRule(ctx, req.(*GetBranchRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_UpdateBranchRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranchRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).UpdateBranchRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/UpdateBranchRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).UpdateBranchRule(ctx, req.(*UpdateBranchRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecureSourceManager_DeleteBranchRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBranchRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecureSourceManagerServer).DeleteBranchRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.securesourcemanager.v1.SecureSourceManager/DeleteBranchRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecureSourceManagerServer).DeleteBranchRule(ctx, req.(*DeleteBranchRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecureSourceManager_ServiceDesc is the grpc.ServiceDesc for SecureSourceManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecureSourceManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.securesourcemanager.v1.SecureSourceManager",
	HandlerType: (*SecureSourceManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInstances",
			Handler:    _SecureSourceManager_ListInstances_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _SecureSourceManager_GetInstance_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _SecureSourceManager_CreateInstance_Handler,
		},
		{
			MethodName: "DeleteInstance",
			Handler:    _SecureSourceManager_DeleteInstance_Handler,
		},
		{
			MethodName: "ListRepositories",
			Handler:    _SecureSourceManager_ListRepositories_Handler,
		},
		{
			MethodName: "GetRepository",
			Handler:    _SecureSourceManager_GetRepository_Handler,
		},
		{
			MethodName: "CreateRepository",
			Handler:    _SecureSourceManager_CreateRepository_Handler,
		},
		{
			MethodName: "DeleteRepository",
			Handler:    _SecureSourceManager_DeleteRepository_Handler,
		},
		{
			MethodName: "GetIamPolicyRepo",
			Handler:    _SecureSourceManager_GetIamPolicyRepo_Handler,
		},
		{
			MethodName: "SetIamPolicyRepo",
			Handler:    _SecureSourceManager_SetIamPolicyRepo_Handler,
		},
		{
			MethodName: "TestIamPermissionsRepo",
			Handler:    _SecureSourceManager_TestIamPermissionsRepo_Handler,
		},
		{
			MethodName: "CreateBranchRule",
			Handler:    _SecureSourceManager_CreateBranchRule_Handler,
		},
		{
			MethodName: "ListBranchRules",
			Handler:    _SecureSourceManager_ListBranchRules_Handler,
		},
		{
			MethodName: "GetBranchRule",
			Handler:    _SecureSourceManager_GetBranchRule_Handler,
		},
		{
			MethodName: "UpdateBranchRule",
			Handler:    _SecureSourceManager_UpdateBranchRule_Handler,
		},
		{
			MethodName: "DeleteBranchRule",
			Handler:    _SecureSourceManager_DeleteBranchRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/securesourcemanager/v1/secure_source_manager.proto",
}
