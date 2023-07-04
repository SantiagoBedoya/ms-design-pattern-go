package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductCommand struct {
	Name  string
	Price float64
}

type GetProductQuery struct {
	ProductID string
}

type CommandHandler interface {
	Handle(command interface{})
}

type QueryHandler interface {
	Handle(query interface{}) interface{}
}

type ProductCommandHandler struct {
}

func (h *ProductCommandHandler) Handle(command interface{}) {
	switch command.(type) {
	case *CreateProductCommand:
		fmt.Println("Creating product")
	default:
		fmt.Println("Undefined command")
	}
}

type ProductQueryHandler struct{}

func (h *ProductQueryHandler) Handle(query interface{}) interface{} {
	switch query.(type) {
	case *GetProductQuery:
		fmt.Println("Getting product")
		result := map[string]string{
			"ID":   "1",
			"Name": "Product 1",
		}
		return result

	default:
		fmt.Println("Undefined query")
		return nil
	}
}

func setupRoutes() {
	commandHandler := &ProductCommandHandler{}
	queryHandler := &ProductQueryHandler{}

	router := gin.Default()
	router.POST("/products", func(ctx *gin.Context) {
		var command CreateProductCommand
		if err := ctx.BindJSON(&command); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		commandHandler.Handle(&command)
		ctx.Status(http.StatusCreated)
	})

	router.GET("/products/:id", func(ctx *gin.Context) {
		productID := ctx.Param("id")
		query := &GetProductQuery{ProductID: productID}
		result := queryHandler.Handle(query)

		if result != nil {
			ctx.JSON(http.StatusOK, result)
		} else {
			ctx.Status(http.StatusNotFound)
		}
	})

	router.Run()
}

func main() {
	setupRoutes()
}
