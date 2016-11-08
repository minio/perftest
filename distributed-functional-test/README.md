# Exhasutive functional test for Minio distributed.
  `server_test.go` from the server side suite test is made generic so that it could now be run against an external
   distributed server instance. 

   Since the suite test is exhastive and covers most of the functionalities, its a good validator till a server side distributed test suite is developed.

   Facilities to run initidividual tests concurrently and under chaos test will be done.

# How to run.

- Set ENDPOINT. 

  ```sh
  $ export ENDPOINT=http://xxx.xx.xxx.xxx:<PORT>
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
  $ go test -check.v
  ```
     OR 

  ```sh
  $ go test -check.f <TestName>

- Here is the list of supported tests.
```
  TestSuiteCommon.TestBucket   
  TestSuiteCommon.TestBucketMultipartList      
  TestSuiteCommon.TestBucketPolicy      
  TestSuiteCommon.TestBucketSQSNotification     
  TestSuiteCommon.TestContentTypePersists      
  TestSuiteCommon.TestCopyObject       
  TestSuiteCommon.TestDeleteBucket     
  TestSuiteCommon.TestDeleteBucketNotEmpty     
  TestSuiteCommon.TestDeleteMultipleObjects    
  TestSuiteCommon.TestDeleteObject     
  TestSuiteCommon.TestEmptyObject      
  TestSuiteCommon.TestGetObjectErrors  
  TestSuiteCommon.TestGetObjectLarge10MiB      
  TestSuiteCommon.TestGetObjectLarge11MiB      
  TestSuiteCommon.TestGetObjectRangeErrors     
  TestSuiteCommon.TestGetPartialObjectLarge10MiB       
  TestSuiteCommon.TestGetPartialObjectLarge11MiB       
  TestSuiteCommon.TestGetPartialObjectMisAligned       
  TestSuiteCommon.TestHeadOnBucket     
  TestSuiteCommon.TestHeadOnObjectLastModified 
  TestSuiteCommon.TestHeader   
  TestSuiteCommon.TestListBuckets      
  TestSuiteCommon.TestListObjectsHandler       
  TestSuiteCommon.TestListObjectsHandlerErrors 
  TestSuiteCommon.TestListenBucketNotificationHandler 
  TestSuiteCommon.TestMultipleObjects  
  TestSuiteCommon.TestNonExistentBucket        
  TestSuiteCommon.TestNotBeAbleToCreateObjectInNonexistentBucket       
  TestSuiteCommon.TestNotImplemented   
  TestSuiteCommon.TestObjectGet       
  TestSuiteCommon.TestObjectGetAnonymous   
  TestSuiteCommon.TestObjectMultipart 
  TestSuiteCommon.TestObjectMultipartAbort  
  TestSuiteCommon.TestObjectMultipartListError 
  TestSuiteCommon.TestObjectValidMD5   
  TestSuiteCommon.TestPartialContent  
  TestSuiteCommon.TestPutBucket      
  TestSuiteCommon.TestPutBucketErrors 
  TestSuiteCommon.TestPutObject       
  TestSuiteCommon.TestPutObjectLongName    
  TestSuiteCommon.TestSHA256Mismatch   
  TestSuiteCommon.TestValidateObjectMultipartUploadID  
  TestSuiteCommon.TestValidateSignature        
```  
