// 代码生成时间: 2025-09-24 01:25:44
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/kataras/iris/v12"
)

// ErrorLogger 是一个中间件，用于记录请求中的错误
func ErrorLogger(ctx iris.Context) {
    next := ctx.Next
    ctx.Next = func(ctx iris.Context) {
        next(ctx)
        if ctx.GetStatusCode() >= 400 {
            logError(ctx)
        }
    }
}

// logError 记录错误日志到文件
func logError(ctx iris.Context) {
    request := ctx.Request()
    response := ctx.ResponseWriter()

    // 获取请求的基本信息
    method := request.Method
    path := request.URL.Path
    status := response.Status()

    // 获取当前时间
    timestamp := time.Now().Format("2006-01-02 15:04:05")

    // 创建日志信息
    logMsg := fmt.Sprintf("[%s] %s %s - %d
", timestamp, method, path, status)

    // 写入日志文件
    file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Printf("Failed to open error log file: %v", err)
        return
    }
    defer file.Close()
    if _, err := file.WriteString(logMsg); err != nil {
        log.Printf("Failed to write to error log file: %v", err)
    }
}

func main() {
    app := iris.New()
    app.Use(ErrorLogger)

    // 设置路由
    app.Get("/error", func(ctx iris.Context) {
        // 模拟一个错误
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Internal Server Error")
    })

    // 启动服务器
    app.Listen(":8080")
}
