# exec-concurent tests. 

- Runs Minio-go functional tests concurrently on the directed Minio instances generating load.

# Instructions to run.

- Set the access Key.

  ```sh
  export ACCESS_KEY=xxxxxx
  ```

- Set the Secret Key.

  ```sh
  export SECRET_KEY=xxxxxxx
  ```

- Set the Minio server endpoint.

  ```sh
  export ENDPOINT=xxx.xx.xxx.xxx:<PORT>
  ```

- Set concurrent level of the load.

  ```sh
  export CONCURRENCY=100
  ```

- Build the api_functional_v4_test.go

  ```sh
  $ go test -c api_functional_v4_test.go 
  ```

- Build exec-concurrent.go

  ```sh
  $ go build exec-concurrent.go
  ```

- Run the test.

  ```sh
  $ ./exec-concurrent <TestName>
  ```

- Here are the list of supported   TestNAMES. 

  ```
  TestMakeBucketError
  TestMakeBucketRegions
  TestPutObjectReadAt
  TestListPartiallyUploaded
  TestGetOjectSeekEnd
  TestGetObjectClosedTwice
  TestRemovePartiallyUploaded
  TestResumablePutObject
  TestResumableFPutObject
  TestFPutObjectMultipart
  TestFPutObject
  TestGetObjectReadSeekFunctional
  TestGetObjectReadAtFunctional
  TestPresignedPostPolicy
  TestCopyObject
  TestFunctional
  ```  
  

     
