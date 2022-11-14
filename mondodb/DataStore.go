package mondodb

import "context"


type DataStore interface {
	 //TODO : add the target and  get all Targets
	  acquireTargets(context.Context,)

}