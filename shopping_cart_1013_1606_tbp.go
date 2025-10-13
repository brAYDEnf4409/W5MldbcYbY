// 代码生成时间: 2025-10-13 16:06:45
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/kataras/iris/v12"
)

// ShoppingCart represents the structure of a shopping cart.
type ShoppingCart struct {
    Items map[string]int
}

// CartService is the service that handles shopping cart operations.
type CartService struct {
    carts map[string]ShoppingCart
}

// NewCartService returns a new instance of CartService.
func NewCartService() *CartService {
    return &CartService{
        carts: make(map[string]ShoppingCart),
    }
}

// AddItem adds an item to the shopping cart.
func (s *CartService) AddItem(cartID, itemID string, quantity int) error {
    if _, exists := s.carts[cartID]; !exists {
        s.carts[cartID] = ShoppingCart{Items: make(map[string]int)}
    }
    s.carts[cartID].Items[itemID] += quantity
    return nil
}

// RemoveItem removes an item from the shopping cart.
func (s *CartService) RemoveItem(cartID, itemID string) error {
    if _, exists := s.carts[cartID]; !exists {
        return fmt.Errorf("cart with ID %s does not exist", cartID)
    }
    if _, exists := s.carts[cartID].Items[itemID]; !exists {
        return fmt.Errorf("item with ID %s does not exist in cart", itemID)
    }
    delete(s.carts[cartID].Items, itemID)
    return nil
}

// GetCart retrieves the shopping cart by its ID.
func (s *CartService) GetCart(cartID string) (*ShoppingCart, error) {
    if cart, exists := s.carts[cartID]; exists {
        return &cart, nil
    }
    return nil, fmt.Errorf("cart with ID %s does not exist", cartID)
}

// hashString hashes a string using SHA-256 and returns a hex-encoded string.
func hashString(s string) string {
    h := sha256.Sum256([]byte(s))
    return hex.EncodeToString(h[:])
}

func main() {
    app := iris.New()
    cartService := NewCartService()

    // Define routes for shopping cart operations.
    app.Post("/cart/{cartID}", func(ctx iris.Context) {
        itemID := ctx.URLParam("itemID")
        quantity := ctx.URLParamDefault("quantity", "1")
        if err := cartService.AddItem(ctx.Param("cartID"), itemID, parseInt(quantity)); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(iris.Map{
            "message": "Item added to cart successfully",
        })
    })

    app.Delete("/cart/{cartID}", func(ctx iris.Context) {
        itemID := ctx.URLParam("itemID")
        if err := cartService.RemoveItem(ctx.Param("cartID\), itemID); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(iris.Map{
            "message": "Item removed from cart successfully",
        })
    })

    app.Get("/cart/{cartID}", func(ctx iris.Context) {
        cart, err := cartService.GetCart(ctx.Param("cartID"))
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(cart)
    })

    // Start the Iris server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// parseInt converts a string to an integer, with error handling.
func parseInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        log.Fatalf("Failed to parse integer from string: %v", err)
    }
    return i
}