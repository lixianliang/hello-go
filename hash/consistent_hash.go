package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"stathat.com/c/consistent"
)

var (
	keyPtr      = flag.Int("keys", 10000000, "key number")
	nodesPtr    = flag.Int("nodes", 3, "node number of old cluster")
	newNodesPtr = flag.Int("new-nodes", 4, "node number of new cluster")
)

func hash(key, nodes int) int {
	return key % nodes
}

func main() {
	flag.Parse()
	var keys, nodes, newNodes = *keyPtr, *nodesPtr, *newNodesPtr

	c := consistent.New()
	for i := 0; i < nodes; i++ {
		c.Add(strconv.Itoa(i))
	}

	newC := consistent.New()
	for i := 0; i < newNodes; i++ {
		newC.Add(strconv.Itoa(i))
	}

	migrate := 0
	for i := 0; i < keys; i++ {
		server, err := c.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatalf("aaa: %v", err)
		}

		newServer, err := newC.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatalf("bbb: %v", err)
		}

		if server != newServer {
			migrate++
		}
	}

	migrateRatio := float64(migrate) / float64(keys)
	fmt.Printf("%f%%\n", migrateRatio*100)
}
