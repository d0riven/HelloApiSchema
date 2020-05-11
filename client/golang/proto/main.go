package main

import (
	"context"
	"time"

	"github.com/d0riven/HelloApiSchema/client/golang/proto/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:5555", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer con.Close()

	client := pb.NewUserServiceClient(con)

	// Get
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		getOut, err := client.Get(ctx, &pb.GetUserInput{
			Id: 1,
		})
		if err != nil {
			panic(err)
		}
		println(getOut.String())
	}()

	// Create
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		createOut, err := client.Create(ctx, &pb.CreateUserInput{
			EmailAddress: "create@example.com",
			LastName:     "新規",
			FirstName:    "作成",
			Birthday:     "0001-01-01",
			Address:      "address",
		})
		if err != nil {
			panic(err)
		}
		println(createOut.String())
	}()

	// Update
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		createOut, err := client.Update(ctx, &pb.UpdateUserInput{
			EmailAddress: "update@example.com",
			LastName:     "既存",
			FirstName:    "更新",
			Birthday:     "0001-01-01",
			Address:      "address",
		})
		if err != nil {
			panic(err)
		}
		println(createOut.String())
	}()
}
