package mondodb

import "go.mongodb.org/mongo-driver/mongo"

type Mongotore struct {
	Db *mongo.Client
	DbName string 

}


// func()