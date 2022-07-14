package main

import (
	"calculator/pb"
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer cc.Close()

	client := pb.NewCalculatorServiceClient(cc)
	callSum(client, 1, 10)
	callSumWithDeadline(client, 1*time.Second, 1, 2)
	callPND(client, 10)
	callAverage(client)
	callFindMax(client)
}

func callSum(c pb.CalculatorServiceClient, num1, num2 int32) {
	log.Println("calling sum api")
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		Number1: num1,
		Number2: num2,
	})
	if err != nil {
		log.Printf("call sum api err %v", err)
		if errStatus, ok := status.FromError(err); ok {
			log.Printf("err msg: %v\n", errStatus.Message())
			log.Printf("err code: %v\n", errStatus.Code())
			if errStatus.Code() == codes.InvalidArgument {
				log.Printf("invalid argument num %v, %v", num1, num2)
				return
			}
		}
	}
	log.Printf("sum api response %v\n", resp.GetResult())
}

func callSumWithDeadline(c pb.CalculatorServiceClient, timeout time.Duration, num1, num2 int32) {
	log.Println("calling sum with deadline api")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.SumWithDeadline(ctx, &pb.SumRequest{
		Number1: num1,
		Number2: num2,
	})

	if err != nil {
		if statusErr, ok := status.FromError(err); ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("calling sum with deadline DeadlineExceeded")
			} else {
				log.Printf("call sum with deadline api err %v", err)
			}
		} else {
			log.Fatalf("call sum with deadline unknown err %v", err)
		}
		return
	}

	log.Printf("sum with deadline api response %v\n", resp.GetResult())
}

func callPND(c pb.CalculatorServiceClient, num int32) {
	log.Println("calling pnd api")
	stream, err := c.PrimeNumberDecomposition(context.Background(), &pb.PNDRequest{
		Number: num,
	})

	if err != nil {
		log.Fatalf("callPND err %v", err)
	}

	for {
		resp, recvErr := stream.Recv()

		// stream finish
		if recvErr == io.EOF {
			log.Println("server finish streaming")
			return
		}

		log.Printf("prime number %v", resp.GetResult())
	}
}

func callAverage(c pb.CalculatorServiceClient) {
	log.Println("calling average api")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("call average err %v", err)
	}

	listReq := []pb.AverageRequest{
		{
			Number: 5,
		},
		{
			Number: 10,
		},
		{
			Number: 15.3,
		},
		{
			Number: 2.4,
		},
		{
			Number: 6.7,
		},
	}

	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatalf("send average request err %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("receive average response err %v", err)
	}

	log.Printf("average response %v", resp)
}

func callFindMax(c pb.CalculatorServiceClient) {
	log.Println("calling find max api")
	stream, err := c.FindMax(context.Background())
	if err != nil {
		log.Fatalf("call find max err %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		// send many requests
		listReq := []pb.FindMaxRequest{
			{
				Number: 5,
			},
			{
				Number: 10,
			},
			{
				Number: 15,
			},
			{
				Number: 2,
			},
			{
				Number: 7,
			},
		}
		for _, req := range listReq {
			err := stream.Send(&req)
			if err != nil {
				log.Fatalf("send finmax request err %v", err)
				break
			}
		}
		// client finish
		stream.CloseSend()
	}()

	go func() {
		// receive many requests
		for {
			resp, err := stream.Recv()

			// server finish
			if err == io.EOF {
				log.Println("ending find max api")
				break
			}
			if err != nil {
				log.Fatalf("receive find max err %v", err)
				break
			}
			log.Printf("max: %v\n", resp.GetMax())
		}
		close(waitc)
	}()

	<-waitc
}
