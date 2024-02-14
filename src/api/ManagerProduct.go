package api

import (
	"io/ioutil"
	"net/http"

	"github.com/crud/docs"
	"github.com/crud/entities"
	"github.com/crud/lib"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ManagerProductObj struct {
}

// @Summary Add a Product item
// @Description Save new product into DynamoDb
// @Accept json
// @Produce json
// @Param entityBody body entities.ProductBody true "Product item to add"
// @Success 201 {object} entities.ProductResponse "Successful response"
// @Router /products/add [post]
func (app *ManagerProductObj) routeAdd(c *gin.Context) {
	var productBody = &entities.ProductBody{}
	if err := c.ShouldBindJSON(productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lib.Logger.Info("Received Product: " + productBody.ToString())

	fakeResponse := entities.ProductResponse{
		Message:    "Product test, not save yet!",
		StatusCode: http.StatusCreated,
	}
	c.JSON(http.StatusOK, fakeResponse)
}

func (app *ManagerProductObj) Run() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Manager products"
	docs.SwaggerInfo.Description = "Create a basic crud"
	docs.SwaggerInfo.Version = "1.0"
	gin.DefaultWriter = ioutil.Discard
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		blocklist := v1.Group("/products")
		{
			blocklist.POST("/add", app.routeAdd)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8003")
}

var ManagerProduct *ManagerProductObj = &ManagerProductObj{}
