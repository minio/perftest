/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/minio/blake2b-simd"
	_ "io/ioutil"
	"os"
	"sync"
	"time"
)

var (
	workers  = flag.Int("w", 10, "Number of workers to use. Defaults to 10.")
	modality = flag.String("m", "", "Modality to use (CT, MR, etc.)")
	runs     = flag.Int("r", 10, "Number of runs to do. Defaults to 10.")
)

// modifyUid modifies the past part of the SOP Instance UID
func modifyUid(uid, modifier string) string {

	s := uid[:len(uid)-len(modifier)] + modifier
	//fmt.Println("modifier", s)

	return s
}

// generateBlob generates an image of a certain modality and makes sure
// that is had a unique SOP Instance UID
func genererateBlob(modifier, modality string) ([]byte, error) {

	data, err := Asset("data/"+modality+".dcm")
	if err != nil {
		return nil, err
	}

	offset := 0x1e0
	size := int(data[offset]) + int(data[offset+1])*0x100
	uid := string(data[offset+2 : offset+2+size])
	uid = modifyUid(uid, modifier)

	// write modified uid back
	copy(data[offset+2:], []byte(uid))

	return data, nil
}

// hashBlob returns the BLAKE2 hash of the blob
func hashBlob(data []byte) (string, error) {

	h := blake2b.New512()
	h.Reset()
	h.Write(data)
	sum := h.Sum(nil)

	return fmt.Sprintf("%x", sum), nil
}

// uploadBlob does an upload to the S3/Minio server
func uploadBlob(data []byte, hash string) error {

	//err := ioutil.WriteFile(hash, data, os.ModePerm)
	//fmt.Println(err)

	credsUp := credentials.NewStaticCredentials("", "", "")
	sessUp := session.New(aws.NewConfig().WithCredentials(credsUp).WithRegion("us-east-1").WithEndpoint("http://127.0.0.1:9000").WithS3ForcePathStyle(true))

	// split key at 2nd character to force creation of directory
	key := hash[0:2] + "/" + hash[2:]

	uploader := s3manager.NewUploader(sessUp)
	var err error
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String("dicom"),
		Key:    aws.String(key),
	})

	return err
}

// Worker routine for uploading an image
func putWorker(imageCh <-chan imageDescriptor, outCh chan<- int) {

	for i := range imageCh {

		data, err := genererateBlob(i.instUID, i.modality)
		if err != nil {
			fmt.Println("Exiting out due to error from genererateBlob:", err)
			return
		}
		hash, err := hashBlob(data)
		if err != nil {
			fmt.Println("Exiting out due to error from hashBlob:", err)
			return
		}
		err = uploadBlob(data, hash)
		if err != nil {
			fmt.Println("Exiting out due to error from uploadBlob:", err)
			return
		}

		outCh <- len(data)
	}
}

type imageDescriptor struct {
	instUID  string
	modality string
}

func main() {
	flag.Parse()
	if *modality == "" {
		fmt.Println("Bad arguments")
		return
	}

	var wg sync.WaitGroup
	imageCh := make(chan imageDescriptor)
	outCh := make(chan int)

	// Start worker go routines
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			putWorker(imageCh, outCh)
		}()
	}

	pid := os.Getpid()

	start := time.Now()

	// Push onto input channel
	go func() {
		for i := 0; i < *runs; i++ {
			t := fmt.Sprintf("%v", time.Now().UnixNano())
			modifier := fmt.Sprintf("%d.%v.%d", pid, t[len(t)-7:], i)
			imageCh <- imageDescriptor{instUID: modifier, modality: "CT"}
		}

		// Close input channel
		close(imageCh)
	}()

	// Wait for workers to complete
	go func() {
		wg.Wait()
		close(outCh) // Close output channel
	}()

	// compute total size of bytes uploaded
	totalSize := 0
	for o := range outCh {
		totalSize += o
	}

	fmt.Println("Total size   :", totalSize, "bytes")
	elapsed := time.Since(start)
	fmt.Println("Elapsed time :", elapsed)
	seconds := float64(elapsed) / float64(time.Second)
	fmt.Printf("Speed        : %4.0f objs/sec\n", float64(*runs) / seconds)
	fmt.Printf("Bandwidth    : %4.0f MBit/sec\n", 8 * float64(totalSize) / seconds / 1024 / 1024)

	//fmt.Println("Number of objects:", len(list))
}
