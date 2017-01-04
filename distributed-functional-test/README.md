# Exhasutive functional test for Minio distributed.
  `server_test.go` from the server side suite test is made generic so that it could now be run against an external
   distributed server instance.

   Since the suite test is exhaustive and covers most of the functionalities, its a good validator for most of the server side functionalities.

   Facilities to run individual tests concurrently and under chaos test will be done.

# How to run.

- Set ENDPOINT.

  ```sh
  $ export S3_ENDPOINT=http://xxx.xx.xxx.xxx:<PORT>
  ```

- Set ACCESS_KEY.

  ```sh
  $ export ACCESS_KEY=xxxx
  ```

- Set SECRET_KEY.

  ```sh
  $ export SECRET_KEY=xxxx
  ```
- Run the test.

  ```sh
  $ go test -v
  ```
     OR

  ```sh
  $ go test -run=<TestName>
  ```

- Here is the list of supported tests.
```
  TestBucket   
  TestBucketMultipartList      
  TestBucketPolicy      
  TestBucketSQSNotification     
  TestContentTypePersists      
  TestCopyObject       
  TestDeleteBucket     
  TestDeleteBucketNotEmpty     
  TestDeleteMultipleObjects    
  TestDeleteObject     
  TestEmptyObject      
  TestGetObjectErrors  
  TestGetObjectLarge10MiB      
  TestGetObjectLarge11MiB      
  TestGetObjectRangeErrors     
  TestGetPartialObjectLarge10MiB       
  TestGetPartialObjectLarge11MiB       
  TestGetPartialObjectMisAligned       
  TestHeadOnBucket     
  TestHeadOnObjectLastModified
  TestHeader   
  TestListBuckets      
  TestListObjectsHandler       
  TestListObjectsHandlerErrors
  TestListenBucketNotificationHandler
  TestMultipleObjects  
  TestNonExistentBucket        
  TestNotBeAbleToCreateObjectInNonexistentBucket       
  TestNotImplemented   
  TestObjectGet       
  TestObjectGetAnonymous   
  TestObjectMultipart
  TestObjectMultipartAbort  
  TestObjectMultipartListError
  TestObjectValidMD5   
  TestPartialContent  
  TestPutBucket      
  TestPutBucketErrors
  TestPutObject       
  TestPutObjectLongName    
  TestSHA256Mismatch   
  TestValidateObjectMultipartUploadID  
  TestValidateSignature        
```  
