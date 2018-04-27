package main

import (
    "os"
    "os/signal"
    "log"
    "github.com/Shopify/sarama"
)

func main() {
    consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := consumer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetOldest)
    //partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := partitionConsumer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    consumed := 0
    ConsumerLoop:
    for {
        select {
        case msg := <-partitionConsumer.Messages():
            log.Printf("Consumed message %s offset %d\n", msg.Value, msg.Offset)
            consumed++
        case <-signals:
            break ConsumerLoop
        }
    }

    log.Printf("consumed: %d\n", consumed)
}