package natsutils

import (
	"errors"
	"fmt"
	"sync"

	nats "github.com/nats-io/nats.go"
)

//NATSComponent is contained reusable logic realated to handling of
// the connection to NATS in the system

type NATSComponent struct {
	//cmu is  the lock from the component
	cmu sync.Mutex

	//nc is the connection of the component
	nc *nats.Conn

	//name is the name of the component
	name string
}

//NewNATSComponent creates a streamComponent
func NewNATSComponent(name string) *NATSComponent {
	return &NATSComponent{
		name: name,
	}
}

//ConnectToServer connects to NATS Server
func (c *NATSComponent) ConnectToServer(url string, options ...nats.Option) error {
	c.cmu.Lock()
	defer c.cmu.Unlock()

	//Connects to NATS with Cluster id , client id and  cutsomized options
	nc, err := nats.Connect(url, options...)
	if err != nil {
		return err
	}
	c.nc = nc
	fmt.Println("connection  to nats successful")
	return nil
}

func (c *NATSComponent) NATS() *nats.Conn {
	c.cmu.Lock()
	defer c.cmu.Unlock()
	return c.nc
}

//JestStreamContext returns the a JetsStreamContext
//for messaging and  stream managament

func (c *NATSComponent) JetsStreamContext(opts ...nats.JSOpt) (nats.JetStreamContext, error) {
	c.cmu.Lock()
	defer c.cmu.Unlock()

	jcontext, err := c.nc.JetStream(opts...)

	if err != nil {
		return nil, errors.New("error inconsistent option provided")
	}
	return jcontext, nil
}

func (c *NATSComponent) Name() string {
	c.cmu.Lock()
	defer c.cmu.Unlock()

	return c.name
}

//ShutDown makes the Component go away
func (c *NATSComponent) ShutDown() error {
	c.NATS().Close()
	return nil
}
