// 代码生成时间: 2025-10-15 01:42:22
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
)

// Product represents an item in the e-commerce platform.
type Product struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
    Price float64 `json:"price"`
}

// ProductController handles CRUD operations for products.
type ProductController struct {
    // no fields needed for this example
}

// BeforeActivation sets up the controller route.
func (ctrl *ProductController) BeforeActivation(b mvc.BeforeActivation) {
    b.Handle("GET", "/products", ctrl.ListProducts)
    b.Handle("POST", "/products", ctrl.CreateProduct)
}

// ListProducts returns a list of products.
func (ctrl *ProductController) ListProducts(ctx iris.Context) {
    // In a real application, this would query a database.
    products := []Product{
        {ID: 1, Name: "Laptop", Price: 999.99},
        {ID: 2, Name: "Smartphone", Price: 599.99},
    }
    ctx.JSON(iris.StatusOK, products)
}

// CreateProduct adds a new product to the e-commerce platform.
func (ctrl *ProductController) CreateProduct(ctx iris.Context) {
    var product Product
    if err := ctx.ReadJSON(&product); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": fmt.Sprintf("Error reading product data: %s", err.Error()),
        })
        return
    }
    // In a real application, this would save the product to a database.
    ctx.JSON(iris.StatusCreated, product)
}

func main() {
    app := iris.New()
    mvc.New(app.Party("/api")).Register(ProductController{})

    // Start the server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("Failed to start server: %s
", err)
    }
}
