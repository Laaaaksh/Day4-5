package Controllers

import (
	"Day4-5/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProduct(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//Models.OrderUpdate()
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
func GetTransaction(c *gin.Context) {
	var transaction []Models.Transaction
	err := Models.GetAllTransaction(&transaction)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}

func CreateTransaction(c *gin.Context) {
	var transaction Models.Transaction
	c.BindJSON(&transaction)
	err := Models.CreateTransaction(&transaction)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
	var product Models.Product
	idProd := c.Params.ByName("id")
	errProd := Models.GetProductByID(&product, idProd)
	if errProd != nil {
		c.JSON(http.StatusNotFound, product)
	}
	c.BindJSON(&product)
	errProd = Models.UpdateProduct(&product, idProd)
	if errProd != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetTransactionByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var transaction Models.Transaction
	err := Models.GetTransactionByID(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}

func UpdateTransaction(c *gin.Context) {
	var transaction Models.Transaction

	id := c.Params.ByName("id")
	err := Models.GetTransactionByID(&transaction, id)
	if err != nil {
		c.JSON(http.StatusNotFound, transaction)
	}
	c.BindJSON(&transaction)
	err = Models.UpdateTransaction(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}

func DeleteTransaction(c *gin.Context) {
	var transaction Models.Transaction
	id := c.Params.ByName("id")
	err := Models.DeleteTransaction(&transaction, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
