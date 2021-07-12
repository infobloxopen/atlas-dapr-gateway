package testing

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"

	commonv1pb "github.com/dapr/go-sdk/dapr/proto/common/v1"
	daprv1 "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"google.golang.org/grpc"

	intg "github.com/infobloxopen/atlas-app-toolkit/integration"
)

const (
	portUpperBound = 60000
)

type ProtectedPort struct {
	sync.Mutex
	PortLowerBound int
}

var Mx = ProtectedPort{
	sync.Mutex{},
	30000,
}

func GetOpenPort() (int, error) {
	Mx.Lock()
	defer Mx.Unlock()
	port, err := intg.GetOpenPortInRange(Mx.PortLowerBound, portUpperBound)
	Mx.PortLowerBound = port + 1
	return port, err
}

type message struct {
	destination string
	message     string
}

type DaprClientMock struct {
	processedMessages []message
}

func (d *DaprClientMock) InvokeService(ctx context.Context, in *daprv1.InvokeServiceRequest, opts ...grpc.CallOption) (*commonv1pb.InvokeResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetState(ctx context.Context, in *daprv1.GetStateRequest, opts ...grpc.CallOption) (*daprv1.GetStateResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetBulkState(ctx context.Context, in *daprv1.GetBulkStateRequest, opts ...grpc.CallOption) (*daprv1.GetBulkStateResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) SaveState(ctx context.Context, in *daprv1.SaveStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) DeleteState(ctx context.Context, in *daprv1.DeleteStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) DeleteBulkState(ctx context.Context, in *daprv1.DeleteBulkStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) ExecuteStateTransaction(ctx context.Context, in *daprv1.ExecuteStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) PublishEvent(ctx context.Context, in *daprv1.PublishEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	d.processedMessages = append(d.processedMessages, message{
		destination: fmt.Sprintf("%s:%s", in.PubsubName, in.Topic),
		message:     string(in.Data),
	})
	return &empty.Empty{}, nil
}

func (d *DaprClientMock) InvokeBinding(ctx context.Context, in *daprv1.InvokeBindingRequest, opts ...grpc.CallOption) (*daprv1.InvokeBindingResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetSecret(ctx context.Context, in *daprv1.GetSecretRequest, opts ...grpc.CallOption) (*daprv1.GetSecretResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetBulkSecret(ctx context.Context, in *daprv1.GetBulkSecretRequest, opts ...grpc.CallOption) (*daprv1.GetBulkSecretResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) RegisterActorTimer(ctx context.Context, in *daprv1.RegisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) UnregisterActorTimer(ctx context.Context, in *daprv1.UnregisterActorTimerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) RegisterActorReminder(ctx context.Context, in *daprv1.RegisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) UnregisterActorReminder(ctx context.Context, in *daprv1.UnregisterActorReminderRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetActorState(ctx context.Context, in *daprv1.GetActorStateRequest, opts ...grpc.CallOption) (*daprv1.GetActorStateResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) ExecuteActorStateTransaction(ctx context.Context, in *daprv1.ExecuteActorStateTransactionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) InvokeActor(ctx context.Context, in *daprv1.InvokeActorRequest, opts ...grpc.CallOption) (*daprv1.InvokeActorResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) GetMetadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*daprv1.GetMetadataResponse, error) {
	panic("implement me")
}

func (d *DaprClientMock) SetMetadata(ctx context.Context, in *daprv1.SetMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}

func (d *DaprClientMock) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	panic("implement me")
}
