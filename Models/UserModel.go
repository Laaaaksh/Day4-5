package Models


type Product struct {
	Id       uint   `json:"id"`
	KeyId string `json:"key_id" binding:"required" gorm:"unique"`
	ProdName string `json:"prod_name" binding:"required"`
	Price    float32 `json:"price" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type User struct{
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required" gorm:"unique"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Transaction struct {
	Id            uint   `json:"id"`
	ProdId        string `json:"prod_id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
	ProdQuantity  uint   `json:"prod_quantity"`
	UserName string `json:"user_name"`
	Amount float32 `json:"amount"`
	OrderTime int64 `json:"order_time"`
	Status string `json:"status"`
}

func (b *Product) TableName() string {
	return "product"
}

func (b * User) TableName() string{
	return "user"
}
func (b *Transaction) TableName() string {
	return "transaction"
}




