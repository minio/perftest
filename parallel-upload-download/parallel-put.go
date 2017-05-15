/*
 * Minio Cloud Storage (C) 2017 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Change this value to test with a different object size.
const defaultObjectSize = 10 * 1024 * 1024

// Uploads all the inputs objects in parallel, upon any error this function panics.
func parallelUploads(objectNames []string, data []byte) {
	var wg sync.WaitGroup
	for _, objectName := range objectNames {
		wg.Add(1)
		go func(objectName string) {
			defer wg.Done()
			if err := uploadBlob(data, objectName); err != nil {
				panic(err)
			}
		}(objectName)
	}
	wg.Wait()
}

// uploadBlob does an upload to the S3/Minio server
func uploadBlob(data []byte, objectName string) error {
	credsUp := credentials.NewStaticCredentials(os.Getenv("ACCESSKEY"), os.Getenv("SECRETKEY"), "")
	sessUp := session.New(aws.NewConfig().
		WithCredentials(credsUp).
		WithRegion("us-east-1").
		WithEndpoint(os.Getenv("ENDPOINT")).
		WithS3ForcePathStyle(true))

	uploader := s3manager.NewUploader(sessUp, func(u *s3manager.Uploader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})
	var err error
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(objectName),
	})

	return err
}

var (
	objectSize = flag.Int("size", defaultObjectSize, "Size of the object to upload.")
)

func main() {
	flag.Parse()

	concurrency := os.Getenv("CONCURRENCY")
	conc, err := strconv.Atoi(concurrency)
	if err != nil {
		log.Fatalln(err)
	}

	var objectNames []string
	for i := 0; i < conc; i++ {
		objectNames = append(objectNames, fmt.Sprintf("object%d", i+1))
	}

	var data = bytes.Repeat([]byte("a"), *objectSize)

	start := time.Now().UTC()
	parallelUploads(objectNames, data)

	totalSize := conc * *objectSize
	elapsed := time.Since(start)
	fmt.Println("Elapsed time :", elapsed)
	seconds := float64(elapsed) / float64(time.Second)
	fmt.Printf("Speed        : %4.0f objs/sec\n", float64(conc)/seconds)
	fmt.Printf("Bandwidth    : %4.0f MBit/sec\n", float64(totalSize)/seconds/1024/1024)
}
