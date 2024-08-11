package change_order_status

type Request struct {
	MessageId   string
	OrderNo     string
	OrderStatus int
}
