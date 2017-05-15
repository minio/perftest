# parallel

## Read and Write

Uploads requested concurrent number of objects to the server.

```
wget https://raw.githubusercontent.com/minio/perftest/master/parallel-upload-download/parallel-put.go
go build parallel-put.go
```

Now that you have built the code.

```
ACCESSKEY=minio SECRETKEY=minio123 ENDPOINT=http://147.75.193.69:9001 CONCURRENCY=500 BUCKET=parallel-put ./parallel-put
Elapsed time : 40.209136441s
Speed        :   29 objs/sec
Bandwidth    :  294 MBytes/sec
```

By default all objects uploaded are 10 MiB in size, to change the size to say 1 MiB. You can use `-size` specified in bytes.

Once you have successfully gathered the results for upload operation, now proceed to download the same uploaded objects.

```
wget https://raw.githubusercontent.com/minio/perftest/master/parallel-upload-download/parallel-get.go
go build parallel-get.go
```

Now that you have built the code, proceed to run.
```
ACCESSKEY=minio SECRETKEY=minio123 ENDPOINT=http://147.75.193.69:9001 CONCURRENCY=1000 BUCKET=parallel-put ./parallel-get
Elapsed time : 6.443437387s
Speed        :  155 objs/sec
Bandwidth    : 1552 MBytes/sec
```
