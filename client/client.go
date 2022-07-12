package main

import (
	"calculator/calculator/calculatorpb"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()

	client := calculatorpb.NewCalculatorServiceClient(cc)
	callSum(client)
}

func callSum(c calculatorpb.CalculatorServiceClient) {
	log.Println("calling sum api")
	resp, err := c.Sum(context.Background(), &calculatorpb.SumRequest{
		Number1: 5,
		Number2: 9,
	})
	if err != nil {
		log.Fatalf("call sum api err %v", err)
	}
	log.Printf("sum api response %v\n", resp.GetResult())
}
