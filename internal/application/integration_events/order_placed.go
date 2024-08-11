package integrationevents

type OrderPlacedEvent struct {
	MessageId      string   `json:"messageId"`
	ConversationId string   `json:"conversationId"`
	MessageType    []string `json:"messageType"`
	Message        Message  `json:"message"`
}

type Message struct {
	OrderNo string      `json:"orderId"`
	Items   []OrderItem `copier:"OrderProducts" json:"items"`
}

type OrderItem struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
