# Kratos datadog plugin

## Requirements
* install datadog agent https://docs.datadoghq.com/agent/
* install datadog go library:
  ```shell
  go get github.com/DataDog/datadog-go/statsd
  ```
* set datadog environment variables
  ```shell
  DD_AGENT_HOST=127.0.0.1
  DD_DOGSTATSD_PORT=8125
  ```

## Usage

### Server

```go
// grpc sever
grpcSrv := grpc.NewServer(
    grpc.Address(":9000"),
    grpc.Middleware(
        metrics.Server(
          metrics.WithSeconds(statsd.NewTiming("name", statsd.WithLabels("kind", "operation"), statsd.WithClient(client))),
          metrics.WithRequests(statsd.NewCounter("name", statsd.WithLabels("kind", "operation", "code", "reason"), statsd.WithClient(client))),
        )
    ),
)

// http server
httpSrv := http.NewServer(
    http.Address(":8000"),
    http.Middleware(
        metrics.Server(
          metrics.WithSeconds(statsd.NewTiming("name", statsd.WithLabels("kind", "operation"), statsd.WithClient(client))),
          metrics.WithRequests(statsd.NewCounter("name", statsd.WithLabels("kind", "operation", "code", "reason"), statsd.WithClient(client))),
        )
    ),
)
```

### Client

```go
// grpc client
conn, err := grpc.DialInsecure(
    context.Background(),
    grpc.WithEndpoint("127.0.0.1:9000"),
    grpc.WithMiddleware(
        metrics.Client(
            metrics.WithSeconds(statsd.NewTiming("name", statsd.WithLabels("kind", "operation"), statsd.WithClient(client))),
            metrics.WithRequests(statsd.NewCounter("name", statsd.WithLabels("kind", "operation", "code", "reason"), statsd.WithClient(client))),
        ),
    ),
)

// http client
conn, err := http.NewClient(
    context.Background(),
    http.WithEndpoint("127.0.0.1:8000"),
    http.WithMiddleware(
        metrics.Client(
            metrics.WithSeconds(statsd.NewTiming("name", statsd.WithLabels("kind", "operation"), statsd.WithClient(client))),
            metrics.WithRequests(statsd.NewCounter("name", statsd.WithLabels("kind", "operation", "code", "reason"), statsd.WithClient(client))),
        ),
    ),
)
```