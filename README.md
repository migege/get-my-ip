# get-my-ip

Get-my-ip is an extremely simple TCP service to tell your real public IP address.

## Prerequisites

[GoLang](https://golang.org/dl/)

## Installation

Type the command below to compile from source.

```
go build main.go
```

Then, run it directly on a KNOWN host or setup it as a service by using some tools (eg. supervisord).

Get-my-ip will listen on port <mark>6666</mark>.

```
./main
```

## Usage

Telnet to 6666 of the KNOWN host, eg. migege.com:6666

```
telnet migege.com 6666
```
