package Routes
import (
	"Day4-5/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/product-api")
	{
		grp1.GET("product", Controllers.GetProduct)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.PUT("product/:id", Controllers.UpdateProduct)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
	}
	grp2 := r.Group("/transaction-api")
	{
		grp2.GET("transaction", Controllers.GetTransaction)
		grp2.POST("transaction", Controllers.CreateTransaction)
		grp2.GET("transaction/:id", Controllers.GetTransactionByID)
		grp2.PUT("transaction/:id", Controllers.UpdateTransaction)
		grp2.DELETE("transaction/:id", Controllers.DeleteTransaction)
	}
	return r
}

