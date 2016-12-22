
# dicomimport 

## Downloading 

You can download the Windows executable for dicomimport from here:

## Configuration

In order to run `dicomimport` you first need to configure the access information to the Minio server. As such you will need to configure with the following three parameters:

- access key
- secret key 
- endpoint

Here are the command line statements to define the environment variables for this: 

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

Meaning of command-line flags

- `-m`: modality
- `-w`: workers
- `-r`: runs
