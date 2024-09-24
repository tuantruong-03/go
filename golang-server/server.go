package main

import (
	"gPRC/bootstrap"
	"gPRC/controllers"
	"gPRC/hello"
	middlewares "gPRC/middlewares"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/

func main() {
	defer createGrpcServer()
	appCtx := bootstrap.BuildApplicationContext()
	router := gin.Default()
	router.Use(middlewares.Recovery())
	controllers.RegisterController(router, appCtx)
	log.Fatal(router.Run(":9091"))

	// router.GET("/health", func(ctx *gin.Context) {
	// 	ctx.String(200, "health")
	// })
	// router.Run(":9091")

}

func createGrpcServer() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	helloServer := hello.Server{}
	hello.RegisterHelloServiceServer(grpcServer, &helloServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
