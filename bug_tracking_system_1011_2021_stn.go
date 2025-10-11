// 代码生成时间: 2025-10-11 20:21:38
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/recover"
    "net/http"
)

// Bug represents the structure of a bug report
type Bug struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Description string `json:"description"`
    Status    string `json:"status"`
}

// NewBug initializes a new Bug with default values
func NewBug(title, description string) *Bug {
    return &Bug{
        Title:     title,
        Description: description,
        Status:    "Open",
    }
}

// bugController handles the HTTP requests
type bugController struct{
    bugs []*Bug
}

// AddBug adds a new bug to the system
func (c *bugController) AddBug(ctx iris.Context) {
    var bug Bug
    if err := ctx.ReadJSON(&bug); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": fmt.Sprintf("Failed to read bug data: %s", err),
        })
        return
    }
    bug.ID = len(c.bugs) + 1
    c.bugs = append(c.bugs, &bug)
    ctx.StatusCode(http.StatusCreated)
    ctx.JSON(bug)
}

// GetBugs retrieves all bugs in the system
func (c *bugController) GetBugs(ctx iris.Context) {
    ctx.JSON(iris.Map{
        "bugs": c.bugs,
    })
}

// UpdateBug updates a bug by its ID
func (c *bugController) UpdateBug(ctx iris.Context) {
    bugID, _ := ctx.Params.GetInt("id")
    var bug Bug
    if err := ctx.ReadJSON(&bug); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": fmt.Sprintf("Failed to read bug data: %s", err),
        })
        return
    }
    for i, b := range c.bugs {
        if b.ID == bugID {
            c.bugs[i] = &bug
            ctx.JSON(iris.Map{
                "message": "Bug updated successfully",
                "bug": c.bugs[i],
            })
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.JSON(iris.Map{
        "error": fmt.Sprintf("Bug with ID %d not found", bugID),
    })
}

// StartBugTrackingSystem launches the bug tracking system
func StartBugTrackingSystem() {
    app := iris.Default()

    // Register middleware
    app.Use(recover.New())
    app.Use(logger.New())

    // Create a new bug controller
    bugCtrl := &bugController{
        bugs: make([]*Bug, 0),
    }

    // Define routes
    bugRoutes := app.Party("/bugs")
    {
        bugRoutes.Post("/add", bugCtrl.AddBug)
        bugRoutes.Get("/list", bugCtrl.GetBugs)
        bugRoutes.Put("/{id:int}", bugCtrl.UpdateBug)
    }

    // Start the server
    app.Listen(":8080")
}

func main() {
    StartBugTrackingSystem()
}
