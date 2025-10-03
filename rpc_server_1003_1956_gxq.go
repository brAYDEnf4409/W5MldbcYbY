// 代码生成时间: 2025-10-03 19:56:50
package main

import (
    "crypto/tls"
    "encoding/json"
    "fmt"
    "log"
    "net"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

// 定义服务接口
type RPCService interface {
    Call(data string) (string, error)
}

// 实现服务接口
type MyRPCService struct {}

func (s *MyRPCService) Call(data string) (string, error) {
    // 示例逻辑：返回传入数据的反向字符串
    reverse := reverseString(data)
    return reverse, nil
}

// 字符串反转函数
func reverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func main() {
    // 设置 gRPC 服务器
    grpcServer := grpc.NewServer()
    // 注册服务
    rpcService := &MyRPCService{}
    RPCServiceServer := NewRPCServiceServer(rpcService)
    grpc.RegisterRPCServiceServer(grpcServer, RPCServiceServer)

    // 设置 HTTPS 证书
    creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
    if err != nil {
        log.Fatalf("Failed to create server credentials: %v", err)
    }
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer.Serve(lis)

    // 设置 Iris 服务器
    irisApp := iris.New()
    irisApp.Get("/rpc", func(ctx iris.Context) {
        // 解析请求体
        var request struct {
            Data string `json:"data"`
        }
        if err := ctx.ReadJSON(&request); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.Writef("Error reading JSON: %s", err)
            return
        }

        // 调用 gRPC 服务
        conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(creds))
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.Writef("Failed to connect to gRPC server: %s", err)
            return
        }
        defer conn.Close()
        c := NewRPCServiceClient(conn)
        r, err := c.Call(context.Background(), &RPCServiceCallRequest{Data: request.Data})
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.Writef("gRPC call failed: %s", err)
            return
        }
        ctx.JSON(r)
    })

    // 启动 Iris 服务器
    if err := irisApp.Listen(":8080", iris.WithCharset("UTF-8")); err != nil {
        log.Fatalf("Failed to start Iris server: %v", err)
    }
}

// RPCServiceServer 是 RPCService 的服务器实现
type RPCServiceServer struct {
    RPCService.UnimplementedRPCServiceServer
    service RPCService
}

func NewRPCServiceServer(s RPCService) *RPCServiceServer {
    return &RPCServiceServer{service: s}
}

func (s *RPCServiceServer) Call(ctx context.Context, r *RPCServiceCallRequest) (*RPCServiceCallResponse, error) {
    response, err := s.service.Call(r.Data)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "call failed: %v", err)
    }
    return &RPCServiceCallResponse{Response: response}, nil
}

// RPCServiceCallRequest 是调用请求的消息
type RPCServiceCallRequest struct {
    Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
}

// RPCServiceCallResponse 是调用响应的消息
type RPCServiceCallResponse struct {
    Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response"`
}
