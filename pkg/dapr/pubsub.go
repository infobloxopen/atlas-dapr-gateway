package dapr

import (
	"context"
	"fmt"

	daprv1 "github.com/dapr/go-sdk/dapr/proto/runtime/v1"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PubsubConfig struct {
	Name  string
	Topic string
}

type PubSub struct {
	logger      *logrus.Logger
	client      daprv1.DaprClient
	Source      PubsubConfig
	Destination []PubsubConfig
}

func InitPubsub(source PubsubConfig, destination []PubsubConfig, appPort int, grpcPort int, log *logrus.Logger) error {
	if source.Name == "" || source.Topic == "" || grpcPort < 1 {
		return fmt.Errorf("publisher is not configured")
	}
	if len(destination) == 0 || appPort < 1 {
		return fmt.Errorf("subscriber is not configured")
	}

	var err error
	ps := &PubSub{
		logger:      log,
		Source:      source,
		Destination: destination,
	}

	ps.initSubscriber(appPort)

	if ps.client, err = ps.initPublisher(grpcPort); err != nil {
		return err
	}

	return nil
}

func (p *PubSub) initPublisher(port int) (daprv1.DaprClient, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to open atlas pubsub connection: %v", err)
	}
	return daprv1.NewDaprClient(conn), nil
}

func (p *PubSub) publish(msg []byte) error {
	if p.client == nil {
		return fmt.Errorf("client is not initialized")
	}

	for _, pubsub := range p.Destination {
		_, err := p.client.PublishEvent(context.Background(), &daprv1.PublishEventRequest{
			Topic:      pubsub.Topic,
			Data:       msg,
			PubsubName: pubsub.Name,
		})
		p.logger.Debugf("Publish to pubsub %q, topic %q. Result: %v", pubsub.Name, pubsub.Topic, err)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PubSub) initSubscriber(appPort int) {
	s, err := daprd.NewService(fmt.Sprintf(":%d", appPort))
	if err != nil {
		p.logger.Fatalf("failed to start the server: %v", err)
	}

	subscription := &common.Subscription{
		PubsubName: p.Source.Name,
		Topic:      p.Source.Topic,
	}
	if err := s.AddTopicEventHandler(subscription, p.eventHandler); err != nil {
		p.logger.Fatalf("error adding handler: %v", err)
	}

	// start the server to handle incoming events
	go func(service common.Service) {
		if err := service.Start(); err != nil {
			p.logger.Fatalf("server error: %v", err)
		}
	}(s)
}

func (p *PubSub) eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	p.logger.Debugf("Incoming message from pubsub %q, topic %q, data: %s", e.PubsubName, e.Topic, e.Data)

	return false, p.publish(e.Data.([]byte))
}
