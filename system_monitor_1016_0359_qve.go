// 代码生成时间: 2025-10-16 03:59:22
package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    "strings"
    "time"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/mem"
    "gopkg.in/olahol/melody.v1"
)

// SystemMonitor 结构体包含websocket连接
type SystemMonitor struct {
    m *melody.Melody
}

// NewSystemMonitor 创建一个新的SystemMonitor实例
func NewSystemMonitor() *SystemMonitor {
    return &SystemMonitor{m: melody.New()}
}

// MonitorResources 监控系统资源并实时通过websocket发送数据
func (sm *SystemMonitor) MonitorResources() {
    sm.m.Handle("/monitor", func(s *melody.Session) {
        ticker := time.NewTicker(5 * time.Second)
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                err := sm.sendResourceData(s)
                if err != nil {
                    log.Printf("Error sending resource data: %s
", err)
                    return
                }
            case <-s.Close:
                return
            }
        }
    })
}

// sendResourceData 发送系统资源数据
func (sm *SystemMonitor) sendResourceData(s *melody.Session) error {
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return err
    }

    load, err := load.Avg()
    if err != nil {
        return err
    }

    mem, err := mem.VirtualMemory()
    if err != nil {
        return err
    }

    return s.Send(map[string]interface{}{
        "cpu":     cpuPercent[0],
        "load":    load.Avg15,
        "ram":     fmt.Sprintf("%d%%", mem.UsedPercent),
        "swap":    "not implemented",
    })
}

// StartServer 启动监控服务器
func StartServer() {
    irisApp := iris.Default()
    monitor := NewSystemMonitor()
    monitor.MonitorResources()
    irisApp.Any("/ws/*", monitor.m.Handle)
    irisApp.Listen(":8080")
}

func main() {
    StartServer()
}
