package place_order

type Request struct {
	UserId  string `validate:"required"`
	OrderNo string `validate:"required"`
}
