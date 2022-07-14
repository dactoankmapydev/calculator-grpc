package main

import (
	"calculator/pb"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("sum called...")
	log.Printf("receive request %v and %v", req.GetNumber1(), req.GetNumber2())
	resp := &pb.SumResponse{
		Result: req.GetNumber1() + req.GetNumber2(),
	}
	return resp, nil
}

func (*server) SumWithDeadline(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("sum with deadline called...")
	log.Printf("receive request %v and %v", req.GetNumber1(), req.GetNumber2())

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			log.Println("context.Canceled...")
			return nil, status.Error(codes.Canceled, "client canceled req")
		}
		time.Sleep(1 * time.Second)
	}
	resp := &pb.SumResponse{
		Result: req.GetNumber1() + req.GetNumber2(),
	}
	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *pb.PNDRequest, stream pb.CalculatorService_PrimeNumberDecompositionServer) error {
	log.Println("PrimeNumberDecomposition called...")
	k := int32(2)
	number := req.GetNumber()
	log.Printf("receive request %v", number)
	for number > 1 {
		if number%k == 0 {
			number = number / k

			// send to client
			stream.Send(&pb.PNDResponse{
				Result: k,
			})
		} else {
			k++
			log.Printf("k increase to %v", k)
		}
	}
	return nil
}

func (*server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average called...")
	var total float32
	var count int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &pb.AverageResponse{
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

func (*server) FindMax(stream pb.CalculatorService_FindMaxServer) error {
	log.Println("Find max called...")
	max := int32(0)
	for {
		req, err := stream.Recv()

		// client finish
		if err == io.EOF {
			log.Println("EOF...")
			return nil
		}
		if err != nil {
			log.Fatalf("err while receive find max %v", err)
			return err
		}
		number := req.GetNumber()
		log.Printf("receive number %v\n", number)
		if number > max {
			max = number
		}

		// send to client
		err = stream.Send(&pb.FindMaxResponse{
			Max: max,
		})
		if err != nil {
			log.Fatalf("send max err %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &server{})

	fmt.Println("calculator is running...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while server %v", err)
	}
}
