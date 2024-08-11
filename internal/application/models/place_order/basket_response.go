package place_order

type BasketResponse struct {
	Data      []basketItem
	IsSuccess bool
}

type basketItem struct {
	ProductId string
	Price     float64
	Quantity  int
}
