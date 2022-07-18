package Models

import (
	"Day4-5/Config"
	"fmt"
	"time"
)

// Authentication Of User

func AuthUser(username string, pass string) error{
	var user User
	if err:= Config.DB.Where("user_name = ?", username). Find(&user).Error; err != nil{
		return fmt.Errorf("user not found")
	}
	if pass == user.Password {
		return nil
	} else {
		return fmt.Errorf("incorrect password")
	}
}
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
func GetAllTransaction(transaction *[]Transaction, username string) (err error) {
	if err = Config.DB.Find(transaction).Error; err != nil {
		return err
	}
	return nil
}

func CreateTransaction(transaction *Transaction, username string) (err error) {

	var Prod Product
	var PrevTran Transaction

	Config.DB.Where( "user_name", username).Last(&PrevTran)

	if PrevTran.Id != 0{
		currentTime := time.Now().Unix()
		if (currentTime - PrevTran.OrderTime) < 300 {
			err := fmt.Errorf("please try again after %d minutes", 5 - (currentTime - PrevTran.OrderTime)/60)
			//fmt.Println(err)
			return err
		}
	}
	if err = Config.DB.Create(transaction).Error; err != nil {
		return err
	}
	if Prod.Quantity < transaction.Quantity {
		Config.DB.Model(transaction).Updates(Transaction{Status: "Failed", Amount: Prod.Price * float32(transaction.Quantity), OrderTime: time.Now().Unix(), UserName: username})
		return fmt.Errorf("out of stock")
	}
		Config.DB.Model(Prod).Update("Quantity", Prod.Quantity-transaction.Quantity)
		Config.DB.Model(transaction).Update(Transaction{Status: "Placed", Amount: Prod.Price * float32(transaction.Quantity), OrderTime: time.Now().Unix(), UserName: username})

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