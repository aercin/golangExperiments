package get_order

type Response struct {
	OrderId  int64 `copier:"Id"`
	OrderNo  string
	UserId   int64
	Status   int
	Products []OrderProduct `copier:"OrderProducts"`
}

type OrderProduct struct {
	ProductId int64
	Price     float64
	Quantity  int
}
