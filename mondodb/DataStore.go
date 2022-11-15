package mondodb

import (
	"context"

	// model "github.com/99designs/gqlgen/example/scalars/model"
	"github.com/franklynobleC/eventstreamTask/model"
)


//use this interface implementation for  the Database interaction
type DataStore interface {
	   acquireTargetsFromDb(context.Context, *model.TargetUser)(model.TargetUser, error)
       listTargethFromDb(context.Context, model.TargetUser)([] model.TargetUser, error)
    
	}
