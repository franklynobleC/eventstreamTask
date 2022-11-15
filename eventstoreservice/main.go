package main

// import "/"
import (
	"context"
	"errors"
	"log"

	service "github.com/franklynobleC/eventstreamTask/eventstore"
	"github.com/franklynobleC/eventstreamTask/natsutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//Errors

var (
	InternalServerError = errors.New("internal server error")
	EventsNotFound      = errors.New("events not found")
	ConnectionFailed    = errors.New("can not connect to server.")
)

//Server is used to implement evenstore.EventStoreServer interface
type server struct {
	service.UnimplementedEventStoreServer
	repository service.Repository
	nats       natsutils.NATSComponent
}

//publishEvent publishes an  event via Nats JetStream server ''
func publishEvent(component *natsutils.NATSComponent, event *service.Event) {
	//Creates JetStreamContext to  pulish messages into Jetstream Stream
	jetStreamContext, _ := component.JetsStreamContext()

	subJect := event.EventType
	eventMsg := []byte(event.EventData)
	//Pusblish message on subject message on subject
	jetStreamContext.Publish(subJect, eventMsg)
	log.Println("published message on Subject" + subJect)

}

func main() {

}

//CODE TO BE REFACTORED, NOT SURE  I KNOW WHAT'S GOING ON
func (s *server) CreateEvent(ctx context.Context, createEventRequest *service.CreateEventRequest) (*service.CreateEventResponse, error) {
	err := s.repository.CreateEvent(ctx, createEventRequest.Event)

	if err != nil {
		return nil, status.Error(codes.Internal, InternalServerError.Error())
	}
	//publishes the event on nats-server
	go publishEvent(&s.nats, createEventRequest.Event)
	return &service.CreateEventResponse{Is_Success: true, Error: ""}, nil

}

func (s *server) GetEvent(ctx context.Context, filter *service.GetEventsRequest) (*service.GetEventResponse, error) {

	events, err := s.repository.GetEvents(ctx, filter)

	if err != nil {
		return nil, status.Error(codes.NotFound, EventsNotFound.Error())
	}

	return &service.GetEventResponse{Event: events}, nil
}
