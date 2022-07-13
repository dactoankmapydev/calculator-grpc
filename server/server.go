package main

import (
	"calculator/calculator/calculatorpb"
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	log.Println("sum called...")
	log.Printf("receive request %v and %v", req.GetNumber1(), req.GetNumber2())
	resp := &calculatorpb.SumResponse{
		Result: req.GetNumber1() + req.GetNumber2(),
	}
	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PNDRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called...")
	k := int32(2)
	number := req.GetNumber()
	log.Printf("receive request %v", number)
	for number > 1 {
		if number%k == 0 {
			number = number / k

			// send to client
			stream.Send(&calculatorpb.PNDResponse{
				Result: k,
			})
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
}

func (*server) Average(stream calculatorpb.CalculatorService_AverageServer) error {
	log.Println("Average called...")
	var total float32
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &calculatorpb.AverageResponse{
				Result: total / float32(count),
			}

			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatalf("err while receive average %v", err)
			return err
		}
		log.Printf("receive request %v", req)
		total += req.GetNumber()
		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while server %v", err)
	}
}
