package dapr

import (
	"context"
	"fmt"
	"github.com/dapr/go-sdk/service/common"
	"log"
	"testing"

	daprv1 "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/sirupsen/logrus"

	mocks "github.com/infobloxopen/atlas-dapr-gateway/testing"
)

type wants struct {
	destination string
	message     string
}

func remove(s []wants, i int) []wants {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

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
		{
			name: "Valid - one valid source, one valid destination",
			args: args{
				source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				destination: []PubsubConfig{
					{
						Name:  "DestinationPubsub",
						Topic: "DestinationTopic",
					},
				},
				appPort:  1250,
				grpcPort: 50001,
				log:      logrus.StandardLogger(),
			},
			wantErr: false,
		},
		{
			name: "Valid - one valid source, two valid destinations",
			args: args{
				source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				destination: []PubsubConfig{
					{
						Name:  "DestinationPubsub1",
						Topic: "DestinationTopic",
					},
					{
						Name:  "DestinationPubsub2",
						Topic: "DestinationTopic",
					},
				},
				appPort:  1250,
				grpcPort: 50001,
				log:      logrus.StandardLogger(),
			},
			wantErr: false,
		},
		{
			name: "Invalid - one valid source, one valid destination, invalid appPort",
			args: args{
				source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				destination: []PubsubConfig{
					{
						Name:  "DestinationPubsub",
						Topic: "DestinationTopic",
					},
				},
				appPort:  0,
				grpcPort: 50001,
				log:      logrus.StandardLogger(),
			},
			wantErr: true,
		},
		{
			name: "Invalid - one valid source, one valid destination, invalid grpcPort",
			args: args{
				source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				destination: []PubsubConfig{
					{
						Name:  "DestinationPubsub",
						Topic: "DestinationTopic",
					},
				},
				appPort:  1250,
				grpcPort: 0,
				log:      logrus.StandardLogger(),
			},
			wantErr: true,
		},
		{
			name: "Valid - no valid source, one valid destination",
			args: args{
				source: PubsubConfig{},
				destination: []PubsubConfig{
					{
						Name:  "DestinationPubsub",
						Topic: "DestinationTopic",
					},
				},
				appPort:  1250,
				grpcPort: 50001,
				log:      logrus.StandardLogger(),
			},
			wantErr: true,
		},
		{
			name: "Inalid - one valid source, no destination",
			args: args{
				source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				destination: []PubsubConfig{},
				appPort:     1250,
				grpcPort:    50001,
				log:         logrus.StandardLogger(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var e error
			if tt.args.appPort > 0 {
				tt.args.appPort, e = mocks.GetOpenPort()
				//= mocks.StartMockDaprPublisher()
				if e != nil {
					log.Fatalf("failed to find free port: %v", e)
				}
			}
			if tt.args.grpcPort > 0 {
				tt.args.grpcPort, e = mocks.GetOpenPort()
			}

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
		message     []byte
	}
	tests := []struct {
		name     string
		fields   fields
		err      bool
		response []wants
	}{
		{
			name: "Valid - initialized client",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: &mocks.DaprClientMock{},
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
				},
				message: []byte("test"),
			},
			err: false,
			response: []wants{
				{
					destination: "Destination1:Topic1",
					message:     "test",
				},
			},
		},
		{
			name: "Valid - multi-destination",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: &mocks.DaprClientMock{},
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
					{
						Name:  "Destination2",
						Topic: "Topic2",
					},
				},
				message: []byte("test"),
			},
			err: false,
			response: []wants{
				{
					destination: "Destination1:Topic1",
					message:     "test",
				},
				{
					destination: "Destination2:Topic2",
					message:     "test",
				},
			},
		},
		{
			name: "Invalid - client is not initialized",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: nil,
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
				},
				message: []byte("test"),
			},
			err:      true,
			response: []wants{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
			te := common.TopicEvent{
				Data:       tt.fields.message,
				Topic:      tt.fields.Source.Topic,
				PubsubName: tt.fields.Source.Name,
			}
			_, err := p.eventHandler(context.Background(), &te)
			if (err != nil) != tt.err {
				t.Fatalf("eventHandler() error = %v, wantErr %v", err, tt.err)
			} else {
				if err == nil {
					for _, req := range tt.fields.Destination {
						success := false
						for i, resp := range tt.response {
							if resp.destination == fmt.Sprintf("%s:%s", req.Name, req.Topic) {
								success = true
								tt.response = remove(tt.response, i)
								break
							}
						}
						if !success {
							t.Fatalf("%s:%s is expected in response, but not found", req.Topic, req.Name)
						}
					}
					if len(tt.response) > 0 {
						t.Fatalf("response has unexpected item(s): %v", tt.response)
					}
				}
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
		message     []byte
	}
	tests := []struct {
		name     string
		fields   fields
		err      bool
		response []wants
	}{
		{
			name: "Valid - initialized client",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: &mocks.DaprClientMock{},
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
				},
				message: []byte("test"),
			},
			err: false,
			response: []wants{
				{
					destination: "Destination1:Topic1",
					message:     "test",
				},
			},
		},
		{
			name: "Valid - multi-destination",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: &mocks.DaprClientMock{},
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
					{
						Name:  "Destination2",
						Topic: "Topic2",
					},
				},
				message: []byte("test"),
			},
			err: false,
			response: []wants{
				{
					destination: "Destination1:Topic1",
					message:     "test",
				},
				{
					destination: "Destination2:Topic2",
					message:     "test",
				},
			},
		},
		{
			name: "Invalid - client is not initialized",
			fields: fields{
				logger: logrus.StandardLogger(),
				client: nil,
				Source: PubsubConfig{
					Name:  "SourcePubsub",
					Topic: "SourceTopic",
				},
				Destination: []PubsubConfig{
					{
						Name:  "Destination1",
						Topic: "Topic1",
					},
				},
				message: []byte("test"),
			},
			err:      true,
			response: []wants{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PubSub{
				logger:      tt.fields.logger,
				client:      tt.fields.client,
				Source:      tt.fields.Source,
				Destination: tt.fields.Destination,
			}
			if err := p.publish(tt.fields.message); (err != nil) != tt.err {
				t.Fatalf("publish() error = %v, wantErr %v", err, tt.err)
			} else {
				if err == nil {
					for _, req := range tt.fields.Destination {
						success := false
						for i, resp := range tt.response {
							if resp.destination == fmt.Sprintf("%s:%s", req.Name, req.Topic) {
								success = true
								tt.response = remove(tt.response, i)
								break
							}
						}
						if !success {
							t.Fatalf("%s:%s is expected in response, but not found", req.Topic, req.Name)
						}
					}
					if len(tt.response) > 0 {
						t.Fatalf("response has unexpected item(s): %v", tt.response)
					}
				}
			}
		})
	}
}
