
package sqldb
type EventStoreDB struct {
	// DB *
}

func NewEventStoreDB() (*EventStoreDB, error) {
	// Connect to the "eventstoredb" database
	// 
	if err != nil {
		return nil, err
	}
	return &EventStoreDB{
		DB: dbEventStore,
	}, nil
}

func (eventstore *EventStoreDB) CreateEventStoreDBSchema() error {
	// Create the "events" table.
	
	return nil
}

func (eventstore *EventStoreDB) Close() {
	eventstore.DB.Close()
}

type  target struct {
	DB *sql.DB
}

func NewTarget() (*OrdersDB, error) {
	// Connect to the "ordersdb" database
	targetDB, err := .Open("")
	if err != nil {
		return nil, err
	}
	
}

//TODO CREATE TARGET SCHEMA
func (target *TargetUser) CreateTargetDBschema() error {
	// Create the "orders" table.
	
}
func (target *targetDB) Close() {
	target.DB.Close()
}