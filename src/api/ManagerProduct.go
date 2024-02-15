package api

import (
	"io/ioutil"
	"net/http"

	"github.com/crud/docs"
	"github.com/crud/entities"
	"github.com/crud/lib"
	"github.com/crud/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ManagerProduct struct {
	controlProduct *service.ControlProduct
}

func (app *ManagerProduct) Init() {
	app.controlProduct = service.NewControlProduct()
}

func NewManagerProduct() *ManagerProduct {
	managerProduct := &ManagerProduct{}
	managerProduct.Init()
	return managerProduct
}

// @Summary Add a Product item
// @Description Save new product into DynamoDb
// @Accept json
// @Produce json
// @Param entityBody body entities.ProductBody true "Product item to add"
// @Success 201 {object} entities.ProductResponse "Successful response"
// @Router /products/add [post]
func (app *ManagerProduct) routeAdd(c *gin.Context) {
	var productBody = &entities.ProductBody{}
	if err := c.ShouldBindJSON(productBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lib.Logger.Info("Received Product: " + productBody.ToString())
	fakeResponse := entities.ProductResponse{}
	err2 := app.controlProduct.SaveProduct(productBody)
	if err2 != nil {
		fakeResponse.StatusCode = http.StatusInternalServerError
		fakeResponse.Message = err2.Error()
		c.JSON(http.StatusInternalServerError, fakeResponse)
		return
	}

	fakeResponse.StatusCode = http.StatusOK
	fakeResponse.Message = "Product saved with success!"
	c.JSON(http.StatusOK, fakeResponse)
}

func (app *ManagerProduct) Run() {

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

	r.Run(":8001")
}
