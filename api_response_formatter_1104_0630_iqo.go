// 代码生成时间: 2025-11-04 06:30:05
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// ApiResponseFormatter is a structure that holds the formatted API response data.
type ApiResponseFormatter struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// NewApiResponseFormatter creates a new instance of ApiResponseFormatter.
func NewApiResponseFormatter(code int, message string, data interface{}) ApiResponseFormatter {
    return ApiResponseFormatter{
        Code:    code,
        Message: message,
        Data:    data,
    }
}

// SuccessResponse creates a success response with a message and data.
func (formatter *ApiResponseFormatter) SuccessResponse(message string, data interface{}) *ApiResponseFormatter {
    return &ApiResponseFormatter{
        Code:    http.StatusOK,
        Message: message,
        Data:    data,
    }
}

// ErrorResponse creates an error response with a message.
func (formatter *ApiResponseFormatter) ErrorResponse(message string) *ApiResponseFormatter {
    return &ApiResponseFormatter{
        Code:    http.StatusInternalServerError,
        Message: message,
        Data:    nil,
    }
}

func main() {
    app := iris.New()
    app.Logger().SetLevel("debug")

    // Success API endpoint
    app.Get("/success", func(ctx iris.Context) {
        // Simulate data retrieval
        data := map[string]string{"key": "value"}
        response := NewApiResponseFormatter(http.StatusOK, "Success", data)
        ctx.JSON(response)
    })

    // Error API endpoint
    app.Get("/error", func(ctx iris.Context) {
        response := NewApiResponseFormatter(http.StatusInternalServerError, "An error occurred", nil)
        ctx.JSON(response)
    })

    // Start the Iris server
    app.Listen(fmt.Sprintf(":%d", 8080))
}
