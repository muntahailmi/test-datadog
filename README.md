# test-datadog
Testing sending metric to datadog

## Installation
- Clone this repository
- Copy `env.sample` into `.env`
	```
	cp env.sample .env
	```
- Modify the `.env` value to fit your environment
  - Fill `NAMESPACE` with your datadog preferred namespace
  - Fill `DD_AGENT_HOST` with address of your datadog-agent host
- Install go dependencies
	```
	go mod download
	```
- Run program
	```
	go run main.go
	```
- Check in your datadog whether there's any new value in metrics `<NAMESPACE>`.service_entity_counter
  - Using metrics explorer is somewhat recommended with link somewhat like this: `https://app.datadoghq.com/metric/explorer?live=true&exp_scope=service%3Ametric-test&exp_group=entity&exp_metric=<NAMESPACE>.service_entity_counter` with `<NAMESPACE>` filled with what you fill in `.env` previously
