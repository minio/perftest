# Performance Tests for Minio

This repository shows the results of some performance tests that were executed on several different server configurations of the Minio Object Storage server.

First we present the results as that is probably what most people are interested in. The 

## Results

| Objects/sec | 4 node | 8 node | 12 node | 16 node |
| ----------- | ------:| ------:| -------:| -------:|
| 1M          |    220 |    353 |         |         |
| 3M          |    221 |    292 |         |         |
| 6M          |    224 |    294 |         |         |
| 12M         |    198 |    262 |         |         |

## Setup

<< describe test setup >>

<< describe setup.sh >>

## Installation

<< describe how to install a server >>

Log into the machine

### Prepare disk

```
wget https://raw.githubusercontent.com/minio/perftest/master/raid_ephemeral.sh
chmod +x raid_ephemeral.sh 
sudo ./raid_ephemeral.sh
sudo chmod 0777 /mnt
mkdir /mnt/distr
```

### Install Golang

```
sudo yum install git
```

```
wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
tar -C ${HOME} -xzf go1.7.3.linux-amd64.tar.gz
echo export GOROOT=${HOME}/go >> ~/.bashrc
echo export GOPATH=${HOME}/work >> ~/.bashrc
echo export PATH=$PATH:${HOME}/go/bin:${HOME}/work/bin >> ~/.bashrc
source ~/.bashrc
```

```
go version
```

### Install minio 
```
go get -u github.com/minio/minio
```


## Start Minio Server

```
minio server 172.31.13.67:/mnt/distr 172.31.13.66:/mnt/distr 172.31.13.69:/mnt/distr 172.31.13.68:/mnt/distr 172.31.14.165:/mnt/distr 172.31.14.164:/mnt/distr 172.31.14.163:/mnt/distr 172.31.14.162:/mnt/distr
```

## Running Performance Tests

```
time ./perftest -p "12" -d 2 -b "lifedrive-100m-usw2" -r "us-west-2" -e "https://s3-us-west-2.amazonaws.com" -a "AKIAI2LOW75FPFTZ5VZA" -s "9RjNUCGY2c+zibVbW3Up8sfx68uFmvsL+4lbwAfE"
```


## Code

<< describe perftest.go >>
