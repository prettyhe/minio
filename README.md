## Minio [![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/minio/minio?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Minio is a minimal cloud storage server for Micro Services & Magnetic Disks. Written in Golang and licensed under [Apache license v2](./LICENSE). Compatible with Amazon S3 APIs. 

## Minio Client

[Minio Client (mc)](https://github.com/minio/mc#minio-client-mc-) provides a modern alternative to Unix commands like ``ls``, ``cat``, ``cp``, ``sync``, and ``diff``. It supports POSIX compatible filesystems and Amazon S3 compatible cloud storage systems. It is entirely written in Golang.

## Amazon S3 Compatible Client Libraries
- [Golang Library](https://github.com/minio/minio-go)
- [Java Library](https://github.com/minio/minio-java)
- [Nodejs Library](https://github.com/minio/minio-js)
- [Python Library](https://github.com/minio/minio-py)
- [.Net Library](https://github.com/minio/minio-dotnet)

### Install [![Build Status](https://travis-ci.org/minio/minio.svg?branch=master)](https://travis-ci.org/minio/minio)[![Build status](https://ci.appveyor.com/api/projects/status/k61d0v3ritbwm2su?svg=true)](https://ci.appveyor.com/project/harshavardhana/minio)

#### Linux, OS X, Windows

~~~
$ go get github.com/minio/minio
~~~

### How to use Minio?

~~~
$ minio server
NAME:
  minio server - Start Minio cloud storage server.

USAGE:
  minio server PATH

EXAMPLES:
  1. Start minio server on Linux.
      $ minio server /home/shared

  2. Start minio server on Windows.
      $ minio server C:\MyShare

  3. Start minio server bound to a specific IP:PORT, when you have multiple network interfaces.
      $ minio --address 192.168.1.101:9000 /home/shared
~~~

~~~
$ minio server ~/Photos
AccessKey: G5GJRH51R2HSUWYPGIX5  SecretKey: uxhBC1Yscut3/u81l5L8Yp636ZUk32N4m/gFASuZ

To configure Minio Client.

	$ wget https://dl.minio.io:9000/updates/2015/Oct/darwin-amd64/mc
	$ chmod 755 mc
	$ ./mc config host add localhost:9000 G5GJRH51R2HSUWYPGIX5 uxhBC1Yscut3/u81l5L8Yp636ZUk32N4m/gFASuZ
	$ ./mc mb localhost/photobucket
	$ ./mc cp ~/Photos... localhost/photobucket

Starting minio server:
Listening on http://127.0.0.1:9000
Listening on http://172.30.2.17:9000
~~~

### Contribute to Minio Project
Please follow Minio [Contributor's Guide](./CONTRIBUTING.md)

