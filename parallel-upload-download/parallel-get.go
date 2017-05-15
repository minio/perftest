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
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type devNull int

func (devNull) WriteAt(p []byte, off int64) (int, error) {
	return len(p), nil
}

// Discard is an io.WriterAt on which
// all WriteAt calls succeed without
// doing anything.
var Discard io.WriterAt = devNull(0)

// Downloads all object names in parallel.
func parallelDownloads(objectNames []string) {
	var wg sync.WaitGroup
	for _, objectName := range objectNames {
		wg.Add(1)
		go func(objectName string) {
			defer wg.Done()
			if err := downloadBlob(objectName); err != nil {
				panic(err)
			}
		}(objectName)
	}
	wg.Wait()
}

// downloadBlob does an upload to the S3/Minio server
func downloadBlob(objectName string) error {
	credsUp := credentials.NewStaticCredentials(os.Getenv("ACCESSKEY"), os.Getenv("SECRETKEY"), "")
	sessUp := session.New(aws.NewConfig().
		WithCredentials(credsUp).
		WithRegion("us-east-1").
		WithEndpoint(os.Getenv("ENDPOINT")).
		WithS3ForcePathStyle(true))

	downloader := s3manager.NewDownloader(sessUp, func(u *s3manager.Downloader) {
		u.PartSize = 64 * 1024 * 1024 // 64MB per part
	})

	var err error
	_, err = downloader.Download(Discard, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET")),
		Key:    aws.String(objectName),
	})

	return err
}

func main() {
	concurrency := os.Getenv("CONCURRENCY")
	conc, err := strconv.Atoi(concurrency)
	if err != nil {
		log.Fatalln(err)
	}

	var objectNames []string
	for i := 0; i < conc; i++ {
		objectNames = append(objectNames, fmt.Sprintf("object%d", i+1))
	}

	start := time.Now().UTC()
	parallelDownloads(objectNames)
	totalSize := conc * 10485760
	elapsed := time.Since(start)
	fmt.Println("Elapsed time :", elapsed)
	seconds := float64(elapsed) / float64(time.Second)
	fmt.Printf("Speed        : %4.0f objs/sec\n", float64(conc)/seconds)
	fmt.Printf("Bandwidth    : %4.0f MBit/sec\n", float64(totalSize)/seconds/1024/1024)
}
