# dateseq
 _dateseq_ is a simple golang command-line utility that prints a date/datetime sequence. It is really a "hello world" like toy **go** project but one might actually find it handy for shell scripting. See the usage examples below

#### Build
To build this utility you will need to:
1. have **go** installed, [golang.org/doc/install](https://golang.org/doc/install)
2. have your $GOPATH set, see [setting your GOPATH](https://github.com/golang/go/wiki/Setting-GOPATH)
3. install one pkg dependency


With your $GOPATH set:

```shell
$ go get github.com/HyperplaneOrg/go-strftime
$ git clone https://github.com/HyperplaneOrg/dateseq.git
$ cd dateseq
$ go build
```

#### Usage Examples
---

```shell
./dateseq 20010101 20010107
20010101
20010102
20010103
20010104
20010105
20010106
20010107
```
The above prints each day from the start date 20010101 to the end date 20010107.

```shell
$ dateseq -s 4 20010101 20010107
20010101
20010105
```
The above prints every forth day between the start date 20010101 to the end date 20010107.

```shell
$ dateseq -s 8 -f "PROD_ABC.PRC.%Y%m%d.hdf" 20010101 20010131
PROD_ABC.PRC.20010101.hdf
PROD_ABC.PRC.20010109.hdf
PROD_ABC.PRC.20010117.hdf
PROD_ABC.PRC.20010125.hdf
```

The above prints every eighth day between the start date 20010101 to the end date 20010131, the resultant dates are embedded in a file name (handy for scripting).


```shell
$ dateseq -s 8 -f "PROD.HRLY.%Y%m%d%H.hdf" 2001010100 2001010123
PROD.HRLY.2001010100.hdf
PROD.HRLY.2001010108.hdf
PROD.HRLY.2001010116.hdf
```

The above prints every eighth hour for the date 20010101, the resultant datestimes are embedded in a file name. 

