package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	prefixFlag   = flag.String("p", "", "prefix for search")
	digitsFlag   = flag.Int("d", 0, "digits for search")
	bucketFlag   = flag.String("b", "", "bucket for search")
	regionFlag   = flag.String("r", "", "region for search")
	endpointFlag = flag.String("e", "", "endpoint for bucket")
	accessFlag   = flag.String("a", "", "access key")
	secretFlag   = flag.String("s", "", "secret key")
	cpu          = flag.Int("cpus", runtime.NumCPU(), "Number of CPUs to use. Defaults to number of processors.")
)

func listS3(wg *sync.WaitGroup, index int, results chan<- []string) {

	var creds *credentials.Credentials
	if *accessFlag != "" && *secretFlag != "" {
		creds = credentials.NewStaticCredentials(*accessFlag, *secretFlag, "")
	} else {
		creds = credentials.AnonymousCredentials
	}
	sess := session.New(aws.NewConfig().WithCredentials(creds).WithRegion(*regionFlag).WithEndpoint(*endpointFlag).WithS3ForcePathStyle(true))

	prefix := fmt.Sprintf("%s%x", *prefixFlag, index)
	prefixMax := ""
	if *digitsFlag != 0 && *digitsFlag <= 0xf {
		prefixMax = fmt.Sprintf("%s%x%x", *prefixFlag, index, *digitsFlag)
	}

	svc := s3.New(sess)
	inputparams := &s3.ListObjectsInput{
		Bucket: aws.String(*bucketFlag),
		Prefix: aws.String(prefix),
	}

	result := make([]string, 0, 1000)

	svc.ListObjectsPages(inputparams, func(page *s3.ListObjectsOutput, lastPage bool) bool {

		prefixMaxReached := false
		for _, value := range page.Contents {
			if prefixMax != "" && (*value.Key)[:len(prefixMax)] == prefixMax {
				prefixMaxReached = true
				break
			}
			copyObject(*value.Key)
			result = append(result, *value.Key)
		}

		if prefixMaxReached || lastPage {
			results <- result
			wg.Done()
			return false
		} else {
			return true
		}
	})
}

func listPrefixes() (map[string]bool, error) {

	var wg sync.WaitGroup
	var results = make(chan []string)

	for i := 0x0; i <= 0xf; i++ {
		wg.Add(1)

		go func(index int) {
			listS3(&wg, index, results)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	prefixHash := make(map[string]bool)
	for result := range results {
		for _, r := range result {
			prefixHash[r] = true
		}
	}

	return prefixHash, nil
}

func copyObject(k string) {

	// Following credentials have restricted access to just GetObject for lifedrive-100m-usw2
	accessKey100mRestrictedPolicy := "AKIAJDXB2JULIRQQVHZQ"
	secretKey100mRestrictedPolicy := "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"

	credsDown := credentials.NewStaticCredentials(accessKey100mRestrictedPolicy, secretKey100mRestrictedPolicy, "")
	sessDown := session.New(aws.NewConfig().WithCredentials(credsDown).WithRegion("us-west-2").WithEndpoint("https://s3-us-west-2.amazonaws.com").WithS3ForcePathStyle(true))

	credsUp := credentials.NewStaticCredentials("9OE9RNWW2PMU5X5A3WHH", "XVjlSlQ/JeLOA7k4Y2zwgNOhnTflirIm++bqgZHb", "")
	sessUp := session.New(aws.NewConfig().WithCredentials(credsUp).WithRegion("us-east-1").WithEndpoint("http://127.0.0.1:9000").WithS3ForcePathStyle(true))

	{
		file, err := os.Create(k)
		if err != nil {
			log.Fatal("Failed to create file", err)
		}
		defer file.Close()

		downloader := s3manager.NewDownloader(sessDown)
		numBytes, err := downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String("lifedrive-100m-usw2"),
				Key:    aws.String(k),
			})
		if err != nil {
			fmt.Println("Failed to download file", err, numBytes)
			return
		}
	}

	for attempt := 0; ; attempt++ {

		fileUp, err := os.Open(k)
		if err != nil {
			log.Fatal("Failed to open file", err)
		}
		defer fileUp.Close()

		uploader := s3manager.NewUploader(sessUp)
		_, err = uploader.Upload(&s3manager.UploadInput{
			Body:   fileUp,
			Bucket: aws.String("bucket100m"),
			Key:    aws.String(k[0:2] + "/" + k[2:]),
		})
		if err != nil {
			if attempt < 3 {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				// abort after three failed attempts
				log.Fatalln("Failed to upload", err)
			}
		}

		fmt.Println("Up:", k[:10])
		break
	}

	os.Remove(k)
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*cpu)
	if *prefixFlag == "" || *bucketFlag == "" || *regionFlag == "" || *endpointFlag == "" {
		fmt.Println("Bad arguments")
		return
	}

	var list map[string]bool
	list, _ = listPrefixes()

	fmt.Println("Number of objects:", len(list))
}
