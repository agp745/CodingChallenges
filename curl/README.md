# GoCurl

GoCurl, similar to [Curl](https://curl.se/), is a command line tool that allows you to send http(s) requests to a desired URL

## Usage

./bin/gocurl \[options] \<URL> \[Flags]

## Options

### -v

**Verbose:**  Will print out headers

### -I

**HEAD:** Will only return Response Headers without the body

### -K

**Keep-Alive:** Sets connection header to Keep-Alive.

### -X \<method>

**Method:** Set the request method. Set to **GET** by default

## Flags

### -d \<data>

**Data:** Set request data

### -H \<headers>

**Headers:** Set any additional request headers

## Installation

Clone the repo and build the cli using the provided makefile

```bash
git clone git@github.com:agp745/CodingChallenges.git
cd CodingChallenges/curl
make build
```

The cli is located within the bin directory

```bash
./bin/gocurl https://example.com
```

**NOTE:** There are example commands in the [makefile](makefile) to test the different features

```bash
make get
make post
make head
```
