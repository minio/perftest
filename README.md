# Performance Tests for Minio

This repository shows the results of some performance tests that were executed on several different server configurations of the Minio Object Storage server.

## Setup

<< describe test setup >>

<< describe setup.sh >>

## Installation

<< describe how to install a server >>

## Results

| Objects/sec | 4 node | 8 node | 12 node | 16 node |
| ----------- | ------:| ------:| -------:| -------:|
| 1M          |    220 |    353 |         |         |
| 3M          |    221 |    292 |         |         |
| 6M          |    224 |    294 |         |         |
| 12M         |    198 |    262 |         |         |

## Code

<< describe perftest.go >>
