package api

import (
	"io/ioutil"
	"net/http"
	"strconv"

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

	lib.Logger.Info("Route Add Product Received Product: " + productBody.ToString())
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

// @Summary Delete a Product item by Id
// @Description Save new product into DynamoDb
// @Accept json
// @Produce json
// @Param id path int true "ID of the product to find"
// @Success 201 {object} entities.ProductResponse "Successful response"
// @Router /products/delete/{id} [delete]
func (app *ManagerProduct) routeDelete(c *gin.Context) {
	var productBody = &entities.ProductBody{}
	var id string = c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	productBody.Id = id2
	fakeResponse := entities.ProductResponse{}
	lib.Logger.Info("Route Delete Product Received Product: " + productBody.ToString())
	err2 := app.controlProduct.DeleteProduct(productBody)
	if err2 != nil {
		fakeResponse.StatusCode = http.StatusInternalServerError
		fakeResponse.Message = err2.Error()
		c.JSON(http.StatusInternalServerError, fakeResponse)
		return
	}
	fakeResponse.StatusCode = http.StatusOK
	fakeResponse.Message = "Product deleted with success!"
	c.JSON(http.StatusOK, fakeResponse)
}

// @Summary Get all products
// @Description Select all products from DynamoDb
// @Accept json
// @Produce json
// @Success 200 {object} entities.ProductListResponse "Successful response"
// @Router /products/list [get]
func (app *ManagerProduct) routeList(c *gin.Context) {

	fakeResponse := &entities.ProductListResponse{}
	product, err2 := app.controlProduct.ListAllProducts()
	if err2 != nil {
		fakeResponse.StatusCode = http.StatusInternalServerError
		fakeResponse.Message = err2.Error()
		c.JSON(http.StatusInternalServerError, fakeResponse)
		return
	}

	fakeResponse.Message = "Items from Product Db"
	fakeResponse.StatusCode = http.StatusOK
	fakeResponse.ProductList = product
	c.JSON(http.StatusOK, fakeResponse)
}

// @Summary Find a Product item
// @Description Select new product into DynamoDb
// @Accept json
// @Produce json
// @Param id path int true "ID of the product to find"
// @Success 200 {object} entities.ProductResponse "Successful response"
// @Router /products/find/{id} [get]
func (app *ManagerProduct) routeFind(c *gin.Context) {
	var productBody = &entities.ProductBody{}
	var id string = c.Param("id")
	id2, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	productBody.Id = id2
	fakeResponse := entities.ProductListResponse{}
	lib.Logger.Info("Route Delete Product Received Product: " + productBody.ToString())
	products, err2 := app.controlProduct.FindProduct(productBody)
	if err2 != nil {
		fakeResponse.StatusCode = http.StatusInternalServerError
		fakeResponse.Message = err2.Error()
		c.JSON(http.StatusInternalServerError, fakeResponse)
		return
	}

	fakeResponse.Message = "Items from Product Db"
	fakeResponse.StatusCode = http.StatusOK
	fakeResponse.ProductList = products
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
			blocklist.DELETE("/delete/:id", app.routeDelete)
			blocklist.GET("/list", app.routeList)
			blocklist.GET("/find/:id", app.routeFind)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8001")
}
