package service

import context "context"


type Repository interface {
 	CreateEvent(ctx context.Context, event *Event) error
 	GetEvents(ctx context.Context, filter *GetEventsRequest) (*Event, error)
 }
