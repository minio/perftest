# upload-perftest
Upload performance testing program for Minio Object Storage server.

The command line options are:

```shell
$ ./upload-perftest --help
Usage of ./upload-perftest:
  -bucket string
    	Bucket to use for uploads test (default "bucket")
  -c int
    	concurrency - number of parallel uploads (default 1)
  -h string
    	service endpoint host (default "localhost:9000")
  -m int
    	Maximum amount of disk usage in GBs (default 80)
  -s	Set if endpoint requires https
  -seed int
    	random seed (default 42)

```

Credentials are passed via the environment variables `ACCESS_KEY` and
`SECRET_KEY`.

After the options, a positional parameter for the size of objects to
upload is required. This can be specified with units like `1MiB` or
`1GB`.

The program generates objects of the given size using a fast,
in-memory, partially-random data generator for object content.

The concurrency options sets the number of parallel uploader threads
and simulates multiple uploaders opening separate connections to the
Minio server endpoint. Each thread sequentially performs uploads of
the given size.

The program exits on any kind of upload error with non-zero exit
status. On a successful run, the program exits when uploads have been
continuosly performed for at least 15 minutes and at least 10 objects
have been uploaded.

Every 10 seconds, the program reports the number of objects uploaded,
the average data bandwidth achieved since the start (total object
bytes sent/duration of the test), the average number of objects
uploaded per second since the start, and the total amount of object
data uploaded.

To not overflow disk capacity of the server, the `-m` options takes
the number of GBs of maximum disk space to use in the test. If the
given amount of data is written, the program randomly overwrites
previously written objects.

A sample run looks like the following:

``` shell
$ ./upload-perftest   -h moslb:80 -m 80 -c 32 10MiB
Generating names for objects...
done.
At 10.01: Avg data b/w: 143.91 MiBps. Avg obj/s: 14.39. Data Written: 1440.00 MiB in 144 objects.
At 20.01: Avg data b/w: 148.45 MiBps. Avg obj/s: 14.85. Data Written: 2970.00 MiB in 297 objects.
At 30.01: Avg data b/w: 150.30 MiBps. Avg obj/s: 15.03. Data Written: 4510.00 MiB in 451 objects.
At 40.01: Avg data b/w: 154.22 MiBps. Avg obj/s: 15.42. Data Written: 6170.00 MiB in 617 objects.
At 50.01: Avg data b/w: 156.18 MiBps. Avg obj/s: 15.62. Data Written: 7810.00 MiB in 781 objects.
...

```
