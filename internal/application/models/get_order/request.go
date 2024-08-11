package get_order

type Request struct {
	OrderNo string `validate:"required"`
}
