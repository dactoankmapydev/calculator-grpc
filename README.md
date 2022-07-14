# calculator-grpc

- Structure
```
.
├── client
│   └── client.go
├── go.mod
├── go.sum
├── Makefile
├── pb
│   ├── calculator.pb.go
│   ├── service_calculator_grpc.pb.go
│   ├── service_calculator.pb.go
│   └── service_calculator.pb.gw.go
├── proto
│   ├── calculator.proto
│   ├── google
│   │   └── api
│   │       ├── annotations.proto
│   │       ├── field_behavior.proto
│   │       ├── httpbody.proto
│   │       └── http.proto
│   └── service_calculator.proto
├── proxy
│   └── proxy.go
├── README.md
├── server
│   └── server.go
└── tools.go
```

- Service
```
service CalculatorService {
    rpc Sum (SumRequest) returns (SumResponse) {
        option (google.api.http) = {
            get: "/v1/sum"
        };
    }
    rpc SumWithDeadline(SumRequest) returns (SumResponse) {}
    rpc PrimeNumberDecomposition (PNDRequest) returns (stream PNDResponse) {}
    rpc Average(stream AverageRequest) returns (AverageResponse) {}
    rpc FindMax(stream FindMaxRequest) returns (stream FindMaxResponse) {}
}
```

- Generate code:
```
make gen-proto
```

- Run:
```
make run-server
make run-client
```

- Run API:
```
make run-server
make run-proxy
```

- Request API
```
http://localhost:8081/v1/sum?number1=1&number2=21
```

- Response API
```
{
"result": 22
}
```
