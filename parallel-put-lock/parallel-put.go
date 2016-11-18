// Start the minio servers in another terminal.
// --------------------
// #!/bin/bash
//
// for i in $(seq 1 6); do
//      minio server --address localhost:900${i} http://localhost:9001/tmp/disk1 http://localhost:9002/tmp/disk2 \
//            http://localhost:9003/tmp/disk3 http://localhost:9004/tmp/disk4 http://localhost:9005/tmp/disk5 \
//            http://localhost:9006/tmp/disk6 &
// done
// ---------------------
//
// This starts 6 disk distributed XL setup locally.
//
// On another terminal compile the code.
//
//   go build parallel-put.go
//
// Grab accessKey and secretKey from minio servers and set them as env
// MINIO_ACCESS_KEY and MINIO_SECRET_KEY respectively.
//
// Proceed to run the test on all the 6 nodes.
//
//  ./parallel-put localhost:900{1..6}
//
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	minio "github.com/minio/minio-go"
)

func getMinioClients(minioNodes []string) ([]*minio.Client, error) {
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	clnts := make([]*minio.Client, len(minioNodes))
	for i, minioNode := range minioNodes {
		client, err := minio.New(minioNode, accessKey, secretKey, false)
		if err != nil {
			return nil, err
		}
		clnts[i] = client
	}
	return clnts, nil
}

func main() {
	clnts, err := getMinioClients(os.Args[1:])
	if err != nil {
		log.Fatalln(err)
	}

	// Data to be uploaded.
	data := make([]string, len(clnts))
	for i := range data {
		data[i] = strings.Repeat(fmt.Sprintf("Hello, World - %d", i), 10)
	}

	// Continously write to all nodes.
	j := 0
	for {
		j++
		fmt.Println("Running: ", j)
		wg := &sync.WaitGroup{}
		for i, d := range data {
			wg.Add(1)
			go func(i int, d string) {
				defer wg.Done()
				_, perr := clnts[i].PutObject("test", "testobject", strings.NewReader(d), "")
				if perr != nil {
					log.Println(perr)
				}
			}(i, d)
		}
		wg.Wait()
	}
}
