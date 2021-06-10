# GO-TAMBOON

This project is a part of [golang challenge][3]

## Goal

* Decrypt the file using a simple [ROT-128][2] algorithm.
* Make donations by creating a Charge via the [Charge API][0] for each row in the
decrypted CSV.
* Produce a brief summary at the end.
* Handle errors gracefully, without stopping the entire process.
* Writes readable and maintainable code.

### Bonus

* Have a good Go package structure.
* Be a good internet citizen and throttles the API call if we hit rate limit.
* Run as fast as possible on a multi-core CPU.
* Allocate as little memory as possible.
* Complete the entire process without leaving large trace of Credit Card numbers
in memory, or on disk.
* Ensure reproducible builds on your workspace.

## How to run

I provide `run.sh` script for execute. The meaning of each flag is

``` bash
-f  string  # path to encrypt file
-pk string  # public key
-sk string  # secret key
-n num      # number of task to run in parallel
-t num      # show top num in summary
-d boolean  # if true the will print log on terminal
```

You can get **public key** and **secret key** from [omise][4] website

## Result

The result in following is ran on 06/2021

Because more of credit card data may expired in the future so the success rate may not be the same

``` bash
$ ./run.sh #  -n 3 -d false

Performing donations on ../data/fng.1000.csv.rot128
Done.

Summary:

total received           : 26863951.03  THB
successfully donated     : 13094005.50  THB
faulty donation          : 13769945.53  THB

all                      : 1000 times
success rate             : 48.20 %
average per person       : 13094.01     THB

Top 3 Donors:
 1. Mrs. Mimosa R Tûk              :     50750.24 THB
 2. Mr. Falco S Bracegirdle        :     50744.57 THB
 3. Mrs. Pimpernel C Headstrong    :     50684.38 THB

Executed time: 3m6.979439124s
```

### Performance (at 06/2021)

| n | time <br /> (m:s:ms) | Success rate <br /> (of all card) |Success rate <br /> (of only valid <br />or non expired card) |
| ----------- | ----------- | ----------- | ----------- |
| 1 | 10:47:96 | 48.20% | 100.00% |
| 2 | 04:34:71 | 48.20% | 100.00% |
| 3 | 03:06:98 | 48.20% | 100.00% |
| 4 | 02:21:42 | 48.20% | 100.00% |
| 5 | 01:54:04 | 47.50% | 98.55% |
| 6 | 01:07:20 | 30.80% | 63.90% |
| 7 | 01:05:43 | 29.00% | 60.17% |
| 8 | 00:48:11 | 23.20% | 48.13% |
| 9 | 00:31:85 | 15.00% | 31.12% |
| 10 | 00:27:20 | 12.90% | 26.76% |

the reason that the success rate of only valid or not expired card become lower when degree of parallelism is increase is because it reach the [Charge API][0] limit rate

 [0]: https://www.omise.co/charges-api
 [1]: https://en.wikipedia.org/wiki/Caesar_cipher
 [2]: https://play.golang.org/p/dCWYyWPHwj4
 [3]: https://github.com/opn-ooo/challenges/tree/master/challenge-go
 [4]: https://www.omise.co
