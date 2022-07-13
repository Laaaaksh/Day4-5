package Models

type Product struct {
	Id       uint   `json:"id"`
	ProdName string `json:"prod_name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
//func NewProduct (id uint, ProdName string, Price string, Quantity string) Product{
//	newproduct := Product{}
//	newproduct.Id = id
//	newproduct.ProdName = ProdName
//	newproduct.Price = Price
//	newproduct.Quantity = Quantity
//	newproduct.ResponseMessage = "Order Added Successfully"
//	return newproduct
//}
type Transaction struct {
	Id            uint   `json:"id"`
	TransactionId uint   `json:"transaction_id"`
	ProdId        string `json:"prod_id"`
	ProdQuantity  uint   `json:"prod_quantity"`
	AmountPaid    uint   `json:"amount_paid"`
}

func (b *Product) TableName() string {
	return "product"
}

func (c *Transaction) TableName1() string {
	return "transaction"
}
//func OrderUpdate(){
//	    dynamicstruct.ExtendStruct(Product{}).
//		AddField("ResponseMessage", "string", `json:"response_message"`).
//		Build().
//		New()
//
//}

