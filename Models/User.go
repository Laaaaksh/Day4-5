package Models
import (
	"Day4-5/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllProduct(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {

	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(product)
	return nil
}

// Transactions Begin Here
func GetAllTransaction(transaction *[]Transaction) (err error) {
	if err = Config.DB.Find(transaction).Error; err != nil {
		return err
	}
	return nil
}

func CreateTransaction(transaction *Transaction) (err error) {
	if err = Config.DB.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func GetTransactionByID(transaction *Transaction, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(transaction *Transaction, id string) (err error) {
	fmt.Println(transaction)
	Config.DB.Save(transaction)
	return nil
}

func DeleteTransaction(transaction *Transaction, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(transaction)
	return nil
}