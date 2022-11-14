package main

// import "/"
import (
	"context"
	"errors"

	pb "github.com/franklynobleC/eventstreamTask/pb"
)

//var  Error

var (
	NoEVentToRetrieve = errors.New("no event to  retrieve.")
	NoEVentStream     = errors.New("no Streams found.")
    ConnectionFailed = errors.New("can not connect to server.")
)

type NewEventStore struct {
	 pb.UnimplementedEventStoreServer
}

func main() {

}
//CODE TO BE REFACTORED, NOT SURE  I KNOW WHAT'S GOING ON
func (n *NewEventStore) GetEvents(context.Context, *pb.GetEventsRequest) (*pb.GetEventResponse, error) {
	//TODO: imeplement this to event Store

	return nil, errors.New("")
}

func (n *NewEventStore) GetEventsStream(context.Context, *pb.EventStore_GetEventsStreamServer) error {
	//  TODO : to Implement this to get eventStream

	return errors.New("")
}
