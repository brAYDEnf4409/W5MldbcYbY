// 代码生成时间: 2025-10-06 16:33:48
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// HealthCheckHandler is a function that handles health check requests.
// It returns a simple message indicating the service is healthy.
func HealthCheckHandler(ctx iris.Context) {
    ctx.JSON(http.StatusOK, iris.Map{
        "status": "ok",
        "timestamp": time.Now().String(),
    })
}

func main() {
    app := iris.New()

    // Register the health check endpoint
    app.Get("/health", HealthCheckHandler)
    
    // Start the server
    log.Printf("Health check service is running on port :8080")
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
    }
}
