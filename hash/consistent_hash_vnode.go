package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"stathat.com/c/consistent"
)

var (
	keyPtr    = flag.Int("keys", 10000000, "key number")
	nodesPtr  = flag.Int("nodes", 3, "node number of old cluster")
	vnodesPtr = flag.Int("new-nodes", 4, "node number of new cluster")
)

func ratio(v1, v2 int) string {
	r := float64(v1) / float64(v2)
	return fmt.Sprintf("%f%%", r*100)
}

func main() {
	flag.Parse()
	var keys, nodes, vnodes, nodeStr = *keyPtr, *nodesPtr, *vnodesPtr, ""

	c := consistent.New()
	for i := 0; i < nodes; i++ {
		nodeStr = fmt.Sprintf("node-%d", i)
		c.Add(nodeStr)
	}
	vnodeC := consistent.New()
	for i := 0; i < nodes; i++ {
		for j := 0; j < vnodes; j++ {
			nodeStr = fmt.Sprintf("node-%d-vnode-%d", i, j)
			vnodeC.Add(nodeStr)
		}
	}

	node0, node1, node2 := 0, 0, 0
	for i := 0; i < keys; i++ {
		server, err := c.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatalf("aaa: %v", err)
		}

		if strings.Compare(server, "node-0") == 0 {
			node0++
		} else if strings.Compare(server, "node-1") == 0 {
			node1++
		} else if strings.Compare(server, "node-2") == 0 {
			node2++
		} else {
			fmt.Println("unknown server:", server)
		}
	}

	fmt.Println("normal mode: node0", ratio(node0, keys/3), ", node1", ratio(node1, keys/3), ", node2", ratio(node2, keys/3))

	node0, node1, node2 = 0, 0, 0
	for i := 0; i < keys; i++ {
		server, err := vnodeC.Get(strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(server, "node-0") {
			node0++
		} else if strings.HasPrefix(server, "node-1") {
			node1++
		} else if strings.HasPrefix(server, "node-2") {
			node2++
		} else {
			fmt.Println("unknown server:", server)
		}
	}

	fmt.Println("vnode mode: node0", ratio(node0, keys/3), ", node1", ratio(node1, keys/3), ratio(node2, keys/3))
}
