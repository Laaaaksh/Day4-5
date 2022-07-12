package Models
type Product struct {
	Id      uint   `json:"id"`
	ProdName    string `json:"prod_name"`
	ProdId string `json:"prod_id"`
	Price string `json:"price"`
	Quantity  	string `json:"quantity"`
}

type Transaction struct {
	p Product
	Id  uint `json:"id"`
	TransactionId uint `json:"transaction_id"`
	ProdId string `json:"prod_id"`
	ProdQuantity  uint `json:"prod_quantity"`
	AmountPaid uint `json:"amount_paid"`
}

func (b * Product) TableName() string{
	return "product"
}

func (c * Transaction) TableName1() string  {
	return "transaction"
}