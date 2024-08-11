package integrationevents

//this service interact with .net services those uses masstransit
//thats why our message format must be support masstransit json schema.
const (
	OrderPlacedEventMessageType       string = "urn:message:core_messages.IntegrationEvents:OrderPlacedEvent"
	IntegrationEventBaseMessageType   string = "urn:message:core_messages:IntegrationEventBase"
	StockNotDecreasedEventMessageType string = "urn:message:core_messages.IntegrationEvents:StockNotDecreasedEvent"
	StockDecreasedEventMessageType    string = "urn:message:core_messages.IntegrationEvents:StockDecreasedEvent"
)
