// 代码生成时间: 2025-10-23 03:28:02
package main

import (
    "fmt"
    "math"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// TableSortFilter 用于定义一个表格排序和过滤的请求结构体
type TableSortFilter struct {
    Page     int    `json:"page"`     // 页码
    PageSize int    `json:"pageSize"` // 每页大小
    SortBy   string `json:"sortBy"`   // 排序字段
    SortOrder string `json:"sortOrder"` // 排序顺序，asc或desc
    Filter   string `json:"filter"`   // 过滤条件
}

// Item 用于模拟表格中的数据项
type Item struct {
    ID        int       `json:"id"`         // 唯一标识
    Name      string    `json:"name"`       // 名称
    Age       int       `json:"age"`        // 年龄
    CreatedAt time.Time `json:"createdAt"` // 创建时间
}

// resultPage 用于返回分页请求的结果
type resultPage struct {
    Items      []Item `json:"items"`      // 本页数据项
    TotalItems int    `json:"totalItems"` // 数据项总数
    TotalPages int    `json:"totalPages"` // 总页数
    Page       int    `json:"page"`       // 当前页码
    PageSize   int    `json:"pageSize"`   // 每页大小
}

// setupRouter 设置路由和处理函数
func setupRouter(app *iris.Application) {
    app.Get("/items", func(ctx iris.Context) {
        filter := ctx.URLParam("filter")
        page := ctx.URLParamDefault("page", "1")
        pageSize := ctx.URLParamDefault("pageSize\, 10")
        sortBy := ctx.URLParamDefault("sortBy", "ID")
        sortOrder := ctx.URLParamDefault("sortOrder\, "asc")

        // 解析页码和每页大小
        pageInt, _ := strconv.Atoi(page)
        pageSizeInt, _ := strconv.Atoi(pageSize)

        // 创建表格排序过滤器实例
        sortFilter := TableSortFilter{
            Page:     pageInt,
            PageSize: pageSizeInt,
            SortBy:   sortBy,
            SortOrder: sortOrder,
            Filter:   filter,
        }

        // 获取数据并分页
        items, totalItems := getData(sortFilter)

        // 计算总页数
        totalPages := int(math.Ceil(float64(totalItems) / float64(sortFilter.PageSize)))

        // 返回分页结果
        ctx.JSON(resultPage{
            Items:      items,
            TotalItems: totalItems,
            totalPages: totalPages,
            Page:       sortFilter.Page,
            PageSize:   sortFilter.PageSize,
        })
    })
}

// getData 模拟从数据库获取数据
func getData(sortFilter TableSortFilter) ([]Item, int) {
    // 这里只是一个示例，实际应用中应该从数据库获取数据
    var items []Item
    for i := 1; i <= 100; i++ {
        item := Item{
            ID:        i,
            Name:      fmt.Sprintf("Item #%d", i),
            Age:       i % 10 + 20, // 年龄在20到29之间
            CreatedAt: time.Now().Add(-24 * time.Hour * time.Duration(i)),
        }
        if sortFilter.Filter == "" || strings.Contains(item.Name, sortFilter.Filter) {
            items = append(items, item)
        }
    }

    // 根据排序字段和顺序对数据进行排序
    sort.Slice(items, func(i, j int) bool {
        switch sortFilter.SortBy {
        case "ID":
            return items[i].ID < items[j].ID
        case "Name":
            return items[i].Name < items[j].Name
        case "Age":
            return items[i].Age < items[j].Age
        case "CreatedAt":
            return items[i].CreatedAt.Before(items[j].CreatedAt)
        default:
            return items[i].ID < items[j].ID
        }
    })

    if sortFilter.SortOrder == "desc" {
        items = reverse(items)
    }

    // 计算总项数
    totalItems := len(items)

    // 分页
    items = items[(sortFilter.Page-1)*sortFilter.PageSize : sortFilter.Page*sortFilter.PageSize]

    return items, totalItems
}

// reverse 反转切片
func reverse(items []Item) []Item {
    for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
        items[i], items[j] = items[j], items[i]
    }
    return items
}

func main() {
    app := iris.New()
    setupRouter(app)
    // 启动服务
    app.Listen(":8080")
}
