package main

import (
        "bytes"
        "fmt"
        "io"
        "io/ioutil"
        "log"
        "math/rand"
        "os"
        "strconv"
        "sync"
        "time"

        minio "github.com/minio/minio-go"
)

var data = bytes.Repeat([]byte("0123456789abcdef"), 5*1024*1024)

func parallelUploads(clients []*minio.Client, bucketName string, objectNames []string) {
        var wg sync.WaitGroup
        for _, clnt := range clients {
                for _, objectName := range objectNames {
                        wg.Add(1)
                        go func(clnt *minio.Client, objectName string) {
                                defer wg.Done()
                                clnt.PutObject(bucketName, objectName, bytes.NewReader(data), "")
                        }(clnt, objectName)
                }
        }
        wg.Wait()
}

func parallelDownloads(clients []*minio.Client, bucketName string, objectNames []string) {
        var wg sync.WaitGroup
       for _, clnt := range clients {
                for _, objectName := range objectNames {
                        wg.Add(1)
                        go func(clnt *minio.Client, objectName string) {
                                defer wg.Done()
                                obj, _ := clnt.GetObject(bucketName, objectName)
                                if obj != nil {
                                        io.Copy(ioutil.Discard, obj)
                                }
                        }(clnt, objectName)
                }
        }
        wg.Wait()
}

var src = rand.NewSource(time.Now().UTC().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyz01234569"
const (
        letterIdxBits = 6                    // 6 bits to represent a letter index
        letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
        letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Function to generate random string for bucket/object names.
func randString(n int) string {
        b := make([]byte, n)
        // A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
        for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
               if remain == 0 {
                        cache, remain = src.Int63(), letterIdxMax
                }
                if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
                        b[i] = letterBytes[idx]
                        i--
                }
                cache >>= letterIdxBits
                remain--
        }
        return string(b)
}

// generate random object name.
func getRandomObjectName() string {
        return randString(16)
}

func main() {
        var s3Clients []*minio.Client
        for _, arg := range os.Args[1:] {
                client, err := minio.New(arg, os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_KEY"), os.Getenv("SECURE") == "1")
                if err != nil {
                        log.Fatalln(err)
                }
                s3Clients = append(s3Clients, client)
        }
        bucketName := os.Getenv("BUCKET")
        concurrency := os.Getenv("CONCURRENCY")
        conc, err := strconv.Atoi(concurrency)
        if err != nil {
                log.Fatalln(err)
        }
        var objectNames []string
        for i := 0; i < conc; i++ {
             objectNames = append(objectNames, getRandomObjectName())
        }
        t1 := time.Now().UTC()
        parallelUploads(s3Clients, bucketName, objectNames)
        t2 := time.Now().UTC()
        fmt.Println("Parallel uploads took %s", t2.Sub(t1))
        t1 = time.Now().UTC()
        parallelDownloads(s3Clients, bucketName, objectNames)
        t2 = time.Now().UTC()
        fmt.Println("Parallel downloads took %s", t2.Sub(t1))
}
