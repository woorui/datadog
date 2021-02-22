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