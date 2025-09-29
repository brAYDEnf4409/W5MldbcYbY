// 代码生成时间: 2025-09-30 03:16:19
package main

import (
    "fmt"
    "math"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// SalaryCalculator 结构体用于存储基础薪资和税率
type SalaryCalculator struct {
    BaseSalary float64
    TaxRate    float64
}

// CalculateSalary 计算税后薪资
func (sc *SalaryCalculator) CalculateSalary() float64 {
    // 计算税前薪资
    grossSalary := sc.BaseSalary
    // 计算税款
    taxAmount := grossSalary * (sc.TaxRate / 100)
    // 计算税后薪资
    netSalary := grossSalary - taxAmount
    return netSalary
}

// NewSalaryCalculator 创建SalaryCalculator实例
func NewSalaryCalculator(baseSalary, taxRate float64) *SalaryCalculator {
    return &SalaryCalculator{
        BaseSalary: baseSalary,
        TaxRate:    taxRate,
    }
}

// handlerSalaryCalculator 处理薪资计算请求
func handlerSalaryCalculator(ctx iris.Context) {
    // 获取请求参数
    baseSalary := ctx.URLParamFloat64("baseSalary")
    taxRate := ctx.URLParamFloat64("taxRate")

    // 参数校验
    if baseSalary <= 0 || taxRate <= 0 || taxRate > 100 {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": "Invalid parameters, baseSalary and taxRate must be greater than 0 and taxRate should not exceed 100.",
        })
        return
    }

    // 创建薪资计算器实例
    calculator := NewSalaryCalculator(baseSalary, taxRate)

    // 计算薪资
    netSalary := calculator.CalculateSalary()

    // 返回结果
    ctx.JSON(iris.Map{
        "netSalary": netSalary,
    })
}

func main() {
    app := iris.New()

    // 注册薪资计算器路由
    app.Get("/salary", handlerSalaryCalculator)

    // 启动服务器
    app.Run(iris.Addr(":8080"))
}
