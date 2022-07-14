# calculator-grpc

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

- Run:
```
make gen-proto
make run-server
make run-client
```
