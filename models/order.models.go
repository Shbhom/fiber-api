package models

type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	ProdRefer int     `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProdRefer"`
	UserRefer int     `json:"user_id"`
	User      User    `gorm:"foreignKey:UserRefer"`
}
