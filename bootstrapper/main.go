package main

import (
	"log"

	"github.com/franklynobleC/eventstreamTask/natsutils"
	"github.com/nats-io/nats.go"
)


func main() {

}



//createJetStream  creates a stream by using JetStreamContext
func createJestream(streamName string, streamSubject string) error {
	natsComponent := natsutils.NewNATSComponent("boostrapper")
	err := natsComponent.ConnectToServer(nats.DefaultURL)

	 if err != nil {
		log.Fatal(err)
	 }
   var js nats.JetStreamContext 

   js, err = natsComponent.JetsStreamContext()

   if err != nil {
	log.Fatalln(err)
   }
   //check if the target stream already exist; if not create it.
   stream, err := js.StreamInfo(streamName)

   if err != nil {
	log.Println(err)
   }
   if stream == nil {
	log.Printf("Creating stream %q and Subject %q", streamName, streamSubject)
	_, err = js.AddStream(&nats.StreamConfig{
     Name: streamName,
	 Subjects: []string{streamSubject},   
   })
   if err != nil {
	return  err 
   }
     
}
return nil 


}

//TODO: Create an EventStore DB 
func createEventStoreDb() {
}

// TODO: Create an EventStore DB
func CreateTargetDB() {

} 