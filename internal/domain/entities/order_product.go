package entities

type OrderProduct struct {
	Id        int64   `gorm:"column:id;primaryKey"`
	OrderId   int64   `gorm:"column:order_id"`
	ProductId string  `gorm:"column:product_id"`
	Price     float64 `gorm:"column:price"`
	Quantity  int     `gorm:"column:quantity"`
}
