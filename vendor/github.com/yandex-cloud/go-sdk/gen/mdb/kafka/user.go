// Code generated by sdkgen. DO NOT EDIT.

//nolint
package kafka

import (
	"context"

	"google.golang.org/grpc"

	kafka "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// UserServiceClient is a kafka.UserServiceClient with
// lazy GRPC connection initialization.
type UserServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements kafka.UserServiceClient
func (c *UserServiceClient) Create(ctx context.Context, in *kafka.CreateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements kafka.UserServiceClient
func (c *UserServiceClient) Delete(ctx context.Context, in *kafka.DeleteUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements kafka.UserServiceClient
func (c *UserServiceClient) Get(ctx context.Context, in *kafka.GetUserRequest, opts ...grpc.CallOption) (*kafka.User, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).Get(ctx, in, opts...)
}

// GrantPermission implements kafka.UserServiceClient
func (c *UserServiceClient) GrantPermission(ctx context.Context, in *kafka.GrantUserPermissionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).GrantPermission(ctx, in, opts...)
}

// List implements kafka.UserServiceClient
func (c *UserServiceClient) List(ctx context.Context, in *kafka.ListUsersRequest, opts ...grpc.CallOption) (*kafka.ListUsersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).List(ctx, in, opts...)
}

type UserIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *UserServiceClient
	request *kafka.ListUsersRequest

	items []*kafka.User
}

func (c *UserServiceClient) UserIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *UserIterator {
	return &UserIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &kafka.ListUsersRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *UserIterator) Next() bool {
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

	it.items = response.Users
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *UserIterator) Value() *kafka.User {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *UserIterator) Error() error {
	return it.err
}

// RevokePermission implements kafka.UserServiceClient
func (c *UserServiceClient) RevokePermission(ctx context.Context, in *kafka.RevokeUserPermissionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).RevokePermission(ctx, in, opts...)
}

// Update implements kafka.UserServiceClient
func (c *UserServiceClient) Update(ctx context.Context, in *kafka.UpdateUserRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewUserServiceClient(conn).Update(ctx, in, opts...)
}
