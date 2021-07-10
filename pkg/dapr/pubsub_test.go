package dapr

import (
	"context"
	"reflect"
	"testing"

	daprv1 "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	"github.com/sirupsen/logrus"
)

func TestInitPubsub(t *testing.T) {
	type args struct {
		source      PubsubConfig
		destination []PubsubConfig
		appPort     int
		grpcPort    int
		log         *logrus.Logger
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitPubsub(tt.args.source, tt.args.destination, tt.args.appPort, tt.args.grpcPort, tt.args.log); (err != nil) != tt.wantErr {
				t.Errorf("InitPubsub() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPubSub_eventHandler(t *testing.T) {
	type fields struct {
		logger      *logrus.Logger
		client      daprv1.DaprClient
		Source      PubsubConfig
		Destination []PubsubConfig
	}
	type args struct {
		ctx context.Context
		e   *common.TopicEvent
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantRetry bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
			gotRetry, err := p.eventHandler(tt.args.ctx, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("eventHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRetry != tt.wantRetry {
				t.Errorf("eventHandler() gotRetry = %v, want %v", gotRetry, tt.wantRetry)
			}
		})
	}
}

func TestPubSub_initPublisher(t *testing.T) {
	type fields struct {
		logger      *logrus.Logger
		client      daprv1.DaprClient
		Source      PubsubConfig
		Destination []PubsubConfig
	}
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    daprv1.DaprClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
			got, err := p.initPublisher(tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("initPublisher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initPublisher() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPubSub_initSubscriber(t *testing.T) {
	type fields struct {
		logger      *logrus.Logger
		client      daprv1.DaprClient
		Source      PubsubConfig
		Destination []PubsubConfig
	}
	type args struct {
		appPort int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
		})
	}
}

func TestPubSub_publish(t *testing.T) {
	type fields struct {
		logger      *logrus.Logger
		client      daprv1.DaprClient
		Source      PubsubConfig
		Destination []PubsubConfig
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
			if err := p.publish(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}