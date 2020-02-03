package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/DataDog/datadog-go/statsd"
)

func main() {
	ddaddr := os.Getenv("DD_AGENT_HOST")
	if ddaddr == "" {
		handleError(errors.New("Empty Address"), "Error on Reading Addr")
	}
	ddaddr += ":8125"
	namespace := os.Getenv("NAMESPACE") + "."

	client, err := statsd.New(ddaddr,
		statsd.WithNamespace(namespace),
		// option below will buffer up commands and send them when the buffer is reached or after 100msec.
		statsd.WithMaxMessagesPerPayload(1), // sets the maximum number of buffer
		statsd.WithTags([]string{"env:" + os.Getenv("ENV")}),
	)
	handleError(err, "Error on statsd.New")

	fmt.Println("Start")
	err = incrCounter(client, "metric-test", "total-test-start", "ok", "test", "start")
	handleError(err, "Error on incrCounter(start)")

	// simulating processing a job
	time.Sleep(time.Second)

	fmt.Println("Perform End")
	err = incrCounter(client, "metric-test", "total-test-perform", "ok", "test", "perform")
	handleError(err, "Error on incrCounter(perform)")
	err = incrCounter(client, "metric-test", "total-test-finished", "ok", "test", "finish")
	handleError(err, "Error on incrCounter(end)")
	fmt.Println("Success")
	err = client.Flush()
	handleError(err, "Error on Flush")
}

func incrCounter(client *statsd.Client, service, entity, status string, tags ...string) error {
	sendTags := []string{
		fmt.Sprintf("service:%s", service),
		fmt.Sprintf("entity:%s", entity),
		fmt.Sprintf("status:%s", status),
	}
	for idx, tag := range tags {
		sendTags = append(sendTags, fmt.Sprintf("tag%d:%s", idx, tag))
	}
	return client.Incr("service_entity_counter", sendTags, 1)
}

func handleError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf(msg+": %v", err))
	}
}
