# GreenWall

![Logo of GreenWall]
(http://icons.iconarchive.com/icons/iconka/easter-egg-bunny/128/green-cute-icon.png)

Status: **In progress** 

[![Build Status](https://travis-ci.org/mtojek/greenwall.svg?branch=master)](https://travis-ci.org/mtojek/greenwall)

## Description

GreenWall is a tiny service health dashboard written in Go (with frontend prepared in Bootstrap). The aim of this project is to develop a small web application that can be run as a live dashboard, which presents health statuses of specified server nodes. 

The app can be installed in a couple of seconds thus do not hesitate to run this on your operational wall screens!

### Screenshots

TODO Desktop

TODO Mobile

## Features

* web live dashboard based on Bootstrap
* easily resizeable dashboard (wall, desktop, mobile screens)
* definition of monitored hosts in a YAML file
* use HTTP endpoints as source of health information
* search for "healthy" phrases in HTTP responses
* install and run in a few seconds!

## Quickstart

Download and install GreenWall:
```bash
go get github.com/mtojek/greenwall
```

Prepare a YAML file (```config.yaml```) with definition of monitored hosts:
```yaml
TODO
```

Run the application:
```bash
greenwall -staticDir $GOPATH/src/github.com/mtojek/greenwall/frontend
```

Go to the live dashboard:

[http://localhost:9001](http://localhost:9001)

## Dist

It is possible to build a GreenWall distribution (```dist.zip```), which can be easily installed on the target host. Firstly, prepare a distribution:

```bash
make dist
```

Then copy the ```dist.zip``` file to the target host and unzip it. Please remember also to provide a ```config.yaml``` file before runnning the ```./greenwall``` binary.

## License

MIT License, see [LICENSE](https://github.com/mtojek/greenwall/blob/master/LICENSE) file.
