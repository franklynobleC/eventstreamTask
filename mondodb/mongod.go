package mondodb

import "go.mongodb.org/mongo-driver/mongo"

type MonoStore struct {
	Db *mongo.Client
	DbName string 

}


// func()