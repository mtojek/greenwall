# GreenWall

![Logo of GreenWall]
(http://icons.iconarchive.com/icons/iconka/easter-egg-bunny/128/green-cute-icon.png)

Status: **Done** (waiting for feedback)

[![Build Status](https://travis-ci.org/mtojek/greenwall.svg?branch=master)](https://travis-ci.org/mtojek/greenwall)

## Description

GreenWall is a tiny service health dashboard written in Go (with frontend prepared in Bootstrap). The aim of this project is to develop a small web application that can be run as a live dashboard, which presents health statuses of specified server nodes. 

The app can be installed in a couple of seconds thus do not hesitate to run this on your operational wall screens!

### Screenshots

#### Desktop view

<img src="https://github.com/mtojek/greenwall/blob/master/screenshot-1.png" alt="Screenshot Desktop" width="640px" />

#### Mobile view

<img src="https://github.com/mtojek/greenwall/blob/master/screenshot-2.png" alt="Screenshot Mobile" width="256px" />

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

Prepare a YAML file (```config.yaml```) with definitions of monitored hosts:
```yaml
---
general:
  healthcheckEvery: 15s
  httpClientTimeout: 5s
  refreshDashboardEvery: 10s
groups:
  - name: Frontend Nodes (us-east-1)
    nodes:
      - name: front-1
        endpoint: https://www.example.com/
        expectedPattern: Example
      - name: front-2
        endpoint: https://www.example.com/
        expectedPattern: WillNotFindThis
  - name: Middleware Nodes (us-west-2)
    nodes:
      - name: middleware-1 with a really long name
        endpoint: https://www.example.com/
      - name: middleware-2
        endpoint: https://www.example.com/
  - name: Backend Nodes (us-west-2)
    nodes:
      - name: backend-1
        endpoint: https://www.example.com/
      - name: backend-2
        endpoint: https://www.example.com/
      - name: backend-3
        endpoint: https://www.example.com/
      - name: backend-4
        endpoint: https://1234567890.example.com/
      - name: backend-5
        endpoint: https://www.example.com/
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

Then copy the ```dist.zip``` file to the target host and unzip it. Please remember also to provide a ```config.yaml``` file before running the ```./greenwall``` binary.

## Contact

Please feel free to leave any comment or feedback by opening a new issue or contacting me directly via [email](mailto:marcin@tojek.pl). Thank you.

## License

MIT License, see [LICENSE](https://github.com/mtojek/greenwall/blob/master/LICENSE) file.
