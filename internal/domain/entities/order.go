package entities

type Order struct {
	Id            int64          `gorm:"column:id;primaryKey"`
	OrderNo       string         `gorm:"column:order_no"`
	UserId        string         `gorm:"column:user_id"`
	Status        int            `gorm:"column:status"`
	OrderProducts []OrderProduct `gorm:"foreignKey:OrderId;references:Id"`
}
