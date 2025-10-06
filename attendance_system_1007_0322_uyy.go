// 代码生成时间: 2025-10-07 03:22:19
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "log"
    "time"
)

// AttendanceRecord represents a record of an attendance punch.
type AttendanceRecord struct {
    ID        string    `json:"id"`
    Timestamp time.Time `json:"timestamp"`
    Status    string    `json:"status"`
}

// NewAttendanceRecord creates a new attendance record.
func NewAttendanceRecord(id string) *AttendanceRecord {
    return &AttendanceRecord{
        ID: id,
        Timestamp: time.Now(),
        Status: "punched",
    }
}

func main() {
    // Initialize Iris
    app := iris.New()

    // Define routes
    app.Get("/health", healthCheck)
    app.Post("/punch", punchIn)

    // Start the Iris server
    log.Fatal(app.Listen(":8080"))
}

// healthCheck is a simple endpoint to check the server's health.
func healthCheck(ctx iris.Context) {
    ctx.JSON(iris.StatusOK, iris.Map{ "status": "ok" })
}

// punchIn handles the punch-in request.
func punchIn(ctx iris.Context) {
    id := ctx.URLParam("id")
    if id == "" {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{ "error": "employee ID is required" })
        return
    }

    record := NewAttendanceRecord(id)
    ctx.JSON(iris.StatusOK, record)
}
