package integrationevents

type StockReportedEvent struct {
	MessageId      string                    `json:"messageId"`
	ConversationId string                    `json:"conversationId"`
	MessageType    []string                  `json:"messageType"`
	Message        StockReportedEventMessage `json:"message"`
}

type StockReportedEventMessage struct {
	OrderNo string `json:"orderId"`
}
