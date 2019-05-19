package utils

import "third-party/src/confluent-kafka-go/kafka"

//ProducerObject is object for kafka message producer
var ProducerObject *kafka.Producer

//ConsumerObject is object for kafka message Consumer
var ConsumerObject *kafka.Consumer

//StoragePath is the path where all albums and images will be stored
var StoragePath string

//MessageQueueAddr is the address of kafkabus
var MessageQueueAddr string
