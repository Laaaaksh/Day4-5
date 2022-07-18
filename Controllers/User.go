package Controllers

import (
	"Day4-5/Models"
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
	username,pass,ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Provide login credentials"})
		return
	}

	var product Models.Product
	err := Models.AuthUser(username,pass)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "add valid product details"})
		return
	} else {
		if err := Models.CreateProduct(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, product)
		}
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
	username,pass,ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Provide login credentials"})
		return
	}
	err := Models.AuthUser(username,pass)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	id := c.Params.ByName("id")
	err = Models.GetProductByID(&product, id)
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
func GetAllTransaction(c *gin.Context) {
	var transaction []Models.Transaction
	username,pass,ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Provide login credentials"})
		return
	}
	err := Models.AuthUser(username, pass)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	err = Models.GetAllTransaction(&transaction, username)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}

func CreateTransaction(c *gin.Context) {
	var transaction Models.Transaction

	username,pass,ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"please provide login credentials"})
		return
	}
	err := Models.AuthUser(username, pass)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.BindJSON(&transaction)


	err = Models.CreateTransaction(&transaction,username)
	if err != nil {
		//fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	} else {
		c.JSON(http.StatusOK, transaction)
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
