package messaging

import (
	"fmt"
	"os"
	"third-party/src/confluent-kafka-go/kafka"
)

var ProducerObject *kafka.Producer
var ConsumerObject *kafka.Consumer

//InitProducer creates and returns the Producer object
func InitProducer(broker string) {
	var err error
	ProducerObject, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Producer %v\n", ProducerObject)

}

//InitConsumer creates and returns the Consumer object
func InitConsumer(broker, group string) {
	var err error
	//Note change the default offset
	ConsumerObject, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		// Enable generation of PartitionEOF when the
		// end of a partition is reached.
		"enable.partition.eof": true,
		"auto.offset.reset":    "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %T\n", ConsumerObject)
}

func WriteMessage(message string, topic string) {
	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)
	//Note
	if err := ProducerObject.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, deliveryChan); err != nil {
		fmt.Println("error producing data=== ", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}

//ReadMessage is used to read messages form kafka topic specified and
//return error in case of failure
func ReadMessage(topic string) []string {
	var returnMessage []string

	if err := ConsumerObject.Subscribe(topic, nil); err != nil {
		fmt.Println("error subscribing to topic")
	}

	run := true
	for run == true {
		select {
		case ev := <-ConsumerObject.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				ConsumerObject.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				ConsumerObject.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				returnMessage = append(returnMessage, string(e.Value))
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
				run = false
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}
	if err := ConsumerObject.Unsubscribe(); err != nil {
		fmt.Println("error Un-subscribing to topic")
	}
	fmt.Printf("Closing consumer\n")
	//ConsumerObject.Close()
	return returnMessage
}

func DisconnectConsumer() {
	ConsumerObject.Close()
}

func DisconnectProducer() {
	ProducerObject.Close()
}
