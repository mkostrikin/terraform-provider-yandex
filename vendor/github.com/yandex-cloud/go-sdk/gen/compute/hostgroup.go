// Code generated by sdkgen. DO NOT EDIT.

//nolint
package compute

import (
	"context"

	"google.golang.org/grpc"

	compute "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// HostGroupServiceClient is a compute.HostGroupServiceClient with
// lazy GRPC connection initialization.
type HostGroupServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) Create(ctx context.Context, in *compute.CreateHostGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) Delete(ctx context.Context, in *compute.DeleteHostGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) Get(ctx context.Context, in *compute.GetHostGroupRequest, opts ...grpc.CallOption) (*compute.HostGroup, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).Get(ctx, in, opts...)
}

// List implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) List(ctx context.Context, in *compute.ListHostGroupsRequest, opts ...grpc.CallOption) (*compute.ListHostGroupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).List(ctx, in, opts...)
}

type HostGroupIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *HostGroupServiceClient
	request *compute.ListHostGroupsRequest

	items []*compute.HostGroup
}

func (c *HostGroupServiceClient) HostGroupIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *HostGroupIterator {
	return &HostGroupIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListHostGroupsRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *HostGroupIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.HostGroups
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *HostGroupIterator) Value() *compute.HostGroup {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HostGroupIterator) Error() error {
	return it.err
}

// ListHosts implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) ListHosts(ctx context.Context, in *compute.ListHostGroupHostsRequest, opts ...grpc.CallOption) (*compute.ListHostGroupHostsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).ListHosts(ctx, in, opts...)
}

type HostGroupHostsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *HostGroupServiceClient
	request *compute.ListHostGroupHostsRequest

	items []*compute.Host
}

func (c *HostGroupServiceClient) HostGroupHostsIterator(ctx context.Context, hostGroupId string, opts ...grpc.CallOption) *HostGroupHostsIterator {
	return &HostGroupHostsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListHostGroupHostsRequest{
			HostGroupId: hostGroupId,
			PageSize:    1000,
		},
	}
}

func (it *HostGroupHostsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListHosts(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Hosts
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *HostGroupHostsIterator) Value() *compute.Host {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HostGroupHostsIterator) Error() error {
	return it.err
}

// ListInstances implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) ListInstances(ctx context.Context, in *compute.ListHostGroupInstancesRequest, opts ...grpc.CallOption) (*compute.ListHostGroupInstancesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).ListInstances(ctx, in, opts...)
}

type HostGroupInstancesIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *HostGroupServiceClient
	request *compute.ListHostGroupInstancesRequest

	items []*compute.Instance
}

func (c *HostGroupServiceClient) HostGroupInstancesIterator(ctx context.Context, hostGroupId string, opts ...grpc.CallOption) *HostGroupInstancesIterator {
	return &HostGroupInstancesIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListHostGroupInstancesRequest{
			HostGroupId: hostGroupId,
			PageSize:    1000,
		},
	}
}

func (it *HostGroupInstancesIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListInstances(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Instances
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *HostGroupInstancesIterator) Value() *compute.Instance {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HostGroupInstancesIterator) Error() error {
	return it.err
}

// ListOperations implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) ListOperations(ctx context.Context, in *compute.ListHostGroupOperationsRequest, opts ...grpc.CallOption) (*compute.ListHostGroupOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).ListOperations(ctx, in, opts...)
}

type HostGroupOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *HostGroupServiceClient
	request *compute.ListHostGroupOperationsRequest

	items []*operation.Operation
}

func (c *HostGroupServiceClient) HostGroupOperationsIterator(ctx context.Context, hostGroupId string, opts ...grpc.CallOption) *HostGroupOperationsIterator {
	return &HostGroupOperationsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &compute.ListHostGroupOperationsRequest{
			HostGroupId: hostGroupId,
			PageSize:    1000,
		},
	}
}

func (it *HostGroupOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *HostGroupOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *HostGroupOperationsIterator) Error() error {
	return it.err
}

// Update implements compute.HostGroupServiceClient
func (c *HostGroupServiceClient) Update(ctx context.Context, in *compute.UpdateHostGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return compute.NewHostGroupServiceClient(conn).Update(ctx, in, opts...)
}
