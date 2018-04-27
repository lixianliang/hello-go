package main

import (
    //"fmt"
    "log"
    "github.com/Shopify/sarama"
)

func main() {
    producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
    if err != nil {
        log.Fatalln(err)
    }
    defer func() {
        if err := producer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    msg := &sarama.ProducerMessage{
        Topic: "my_topic", 
        Value: sarama.StringEncoder("testing 123456789"),
    }

    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        log.Printf("failed to send message: %s\n", err)
    } else {
        log.Printf("> message send to partition %d at offset %d\n", partition, offset)
    }
}
