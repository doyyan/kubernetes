package adapter

import (
	"net/http"

	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql/model"
	"github.com/doyyan/kubernetes/internal/app/adapter/repository"
	"github.com/doyyan/kubernetes/internal/app/application/usecase"
	"github.com/doyyan/kubernetes/internal/app/domain/valueObject"
	"github.com/gin-gonic/gin"
)

var (
	cusotmerRepository = repository.Customer{}
	productRepository  = repository.Product{}
	orderRepository    = repository.Order{}
)

type Controller struct {
}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	r.POST("/order", ctrl.createOrder)
	r.POST("/customer", ctrl.createCustomer)
	return r
}

func (ctrl Controller) createOrder(c *gin.Context) {
	customerID := c.Query("customerId")
	productID := c.Query("productId")
	args := usecase.CreateOrderArgs{
		CustomerID:         customerID,
		ProductID:          productID,
		CustomerRepository: cusotmerRepository,
		ProductRepository:  productRepository,
		OrderRepository:    orderRepository,
	}
	order := usecase.CreateOrder(args)
	c.JSON(200, order)
}

func (ctrl Controller) createCustomer(c *gin.Context) {
	customer := model.Customer{}
	if err := c.BindJSON(&customer); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	args := usecase.CreateCustomerArgs{
		Customer: valueObject.Customer{
			Name: customer.Name,
		},
		CustomerRepository: cusotmerRepository,
	}
	usecase.CreateCustomer(args)
	c.JSON(200, customer)
}
