// 代码生成时间: 2025-10-29 16:49:07
@author: 你的名字
@date: 2023-06-08
*/

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os/exec"
    "path/filepath"
    "sort"
    "strings"

    "github.com/kataras/iris/v12"
)

// AutoGrader 自动批改工具结构体
type AutoGrader struct {
    // 存储测试用例
    TestCases []Testcase
}

// Testcase 测试用例结构体
type Testcase struct {
    Input  string `json:"input"`  // 输入数据
    Output string `json:"output"` // 期望输出
    Command string `json:"command"` // 要执行的命令
}

// Response 响应结构体
type Response struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Points  int    `json:"points"`
}

// NewAutoGrader 创建自动批改工具实例
func NewAutoGrader() *AutoGrader {
    return &AutoGrader{
        TestCases: []Testcase{
            // 添加测试用例
        },
    }
}

// RunTest 运行测试用例
func (ag *AutoGrader) RunTest(codePath string) ([]Response, error) {
    var responses []Response
    for _, tc := range ag.TestCases {
        // 编译代码
        output, err := exec.Command("gcc", codePath).Output()
        if err != nil {
            responses = append(responses, Response{
                Success: false,
                Message: fmt.Sprintf("编译错误: %s", err),
            })
            continue
        }

        // 运行代码
        command := exec.Command("./a.out", tc.Input)
        output, err = command.CombinedOutput()
        if err != nil {
            responses = append(responses, Response{
                Success: false,
                Message: fmt.Sprintf("运行错误: %s", err),
            })
            continue
        }

        // 比较输出
        actualOutput := strings.TrimSpace(string(output))
        if actualOutput != tc.Output {
            responses = append(responses, Response{
                Success: false,
                Message: fmt.Sprintf("输出不匹配，期望: %s, 实际: %s", tc.Output, actualOutput),
            })
            continue
        }

        // 测试通过
        responses = append(responses, Response{
            Success: true,
            Message: "测试通过",
            Points:  10,
        })
    }
    return responses, nil
}

func main() {
    app := iris.New()
    ag := NewAutoGrader()

    // 上传代码并获取评分
    app.Post("/submit", func(ctx iris.Context) {
        file, _, err := ctx.FormFile("code")
        if err != nil {
            ctx.JSON(http.StatusBadRequest, Response{
                Success: false,
                Message: "文件上传失败",
            })
            return
        }
        defer file.Close()

        // 保存文件
        filename := filepath.Base(file.Filename)
        savePath := fmt.Sprintf("./uploads/%s", filename)
        if err := ctx.SaveFormFile(*file, savePath); err != nil {
            ctx.JSON(http.StatusInternalServerError, Response{
                Success: false,
                Message: "文件保存失败",
            })
            return
        }

        // 运行测试
        responses, err := ag.RunTest(savePath)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, Response{
                Success: false,
                Message: fmt.Sprintf("测试运行失败: %s", err),
            })
            return
        }

        // 返回评分结果
        var points int
        for _, resp := range responses {
            if resp.Success {
                points += resp.Points
            }
        }
        ctx.JSON(http.StatusOK, Response{
            Success: true,
            Message: fmt.Sprintf("总得分: %d", points),
            Points:  points,
        })
    })

    // 启动服务器
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatal(err)
    }
}