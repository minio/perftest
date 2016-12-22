
# dicomimport 

## Introduction

This is a load testing tool for Minio (or any other S3 compatible server) with an ephasis on medical images in the DICOM format. It embeds images of various modalities that are modified on the fly to generate unique binary objects. Each image is then hashed before being uploaded to the server whereby the hash is used as a key name of the object.

## Downloading 

You can download the Windows executable for dicomimport from here: https://github.com/minio/perftest/releases/download/v0.1/dicomimport.exe

## Building from source

Make sure you have git and golang installed, and then run as follows:

```
go get github.com/minio/perftest/dicomimport
```

## Preparation

Make sure a bucket called dicom is available.

Using `mc` you can create it as follows:

```
mc mb myminio/dicom
```

## Configuration

In order to run `dicomimport` you first need to configure the access information to the Minio server. As such you will need to configure with the following three parameters:

- access key
- secret key 
- endpoint

Here are the command line statements to define the environment variables for this: 

Windows:
```
set ACCESSKEY=5D94Q9WPYAV26D068GIO
set SECRETKEY=GOgBwUsaKn3RmWwO25zq+ZyqLeuSK2aNGu7Z7GTA
set ENDPOINT=http://172.31.17.143:9000
```

## How to run

You can run dicomimport as follows:

```
dicomimport -m "CT" -w 50 -r 1000
```

The meaning of command line flags is as follows

- `-m`: modality (currently "CT" or "MR")
- `-w`: number of worker threads in parallel
- `-r`: total number of objects to upload
