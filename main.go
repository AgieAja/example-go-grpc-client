package main

import (
	"context"
	"log"

	pb "github.com/AgieAja/example-go-grpc/examplepb/proto"
	pbProduct "github.com/AgieAja/example-go-grpc/productpb/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//connect to grpc
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		return
	}

	// create a new gin server
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/greet", greet(conn))
	r.Run(":8080")
}

// handler greeting
func greet(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := pb.NewGreetingServiceClient(conn)
		ctx := context.Background()
		in := &pb.GreetRequest{Name: "Agie"}
		result, err := client.Greet(ctx, in)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": result.GetMessage()})
	}
}

func getProduct(conn *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		client := pbProduct.NewProductServiceClient(conn)
		ctx := context.Background()
		result, err := client.GetProduct(ctx, nil)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"product": result})
	}
}
