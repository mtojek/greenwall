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
* HTTP endpoints can be used as source of health information
* search for "healthy" phrases in HTTP responses
* ICMP ping
* check expired SSL certificates with TLS health check
* pluggable health checks (waiting for TCP, DNS, REST, SOAP and others!)
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
        type: http_check
        parameters:
            expectedPattern: Example
      - name: front-2
        endpoint: https://www.example.com/
        type: http_check
        parameters:
            expectedPattern: WillNotFindThis
  - name: Middleware Nodes (us-west-2)
    nodes:
      - name: middleware-1 with a really long name
        endpoint: https://www.example.com/
        type: http_check
      - name: middleware-2
        endpoint: https://www.example.com/
        type: http_check
  - name: Backend Nodes (us-west-2)
    nodes:
      - name: backend-1
        endpoint: https://www.example.com/
        type: http_check
      - name: backend-2
        endpoint: https://www.example.com/
        type: http_check
      - name: backend-3
        endpoint: https://www.example.com/
        type: http_check
      - name: backend-4
        endpoint: https://1234567890.example.com/
        type: http_check
      - name: backend-5
        endpoint: https://www.example.com/
        type: http_check
```

Run the application:
```bash
greenwall -staticDir $GOPATH/src/github.com/mtojek/greenwall/frontend
```

Go to the live dashboard:

[http://localhost:9001](http://localhost:9001)

### Running in cloud - Heroku, Cloud Foundry, etc.

As an alternative to command line arguments, GreenWall can also read primary configuration (host, port, config file, static path, etc.) from environment variables. Names of variables can be listed with a ```-help``` switch. 

Sample command:

```
PORT=9001 CONFIG=config.yaml STATIC_DIR=frontend greenwall
```

## Building

The project may be rebuilt using a single command - ```make```. This includes downloading dependencies, formatting, building code and testing.

The building process may require higher user permission:
```bash
--- PASS: TestLint (8.93s)
PASS
ok      github.com/greenwall    8.932s
?       github.com/greenwall/middleware/application     [no test files]
?       github.com/greenwall/middleware/healthcheck     [no test files]
?       github.com/greenwall/middleware/httpserver      [no test files]
?       github.com/greenwall/middleware/monitoring      [no test files]
go test -race  -i ./...
go install runtime/internal/sys: open /usr/lib/golang/pkg/linux_amd64_race/runtime/internal/sys.a: permission denied
make: *** [test] Error 1
[me@centos7t01 greenwall]$
```
To resolve this issue, please elevate user permissions with ```sudo``` or use local Go installation.

## Creating own, pluggable health check

The author is welcome to any contributions to this project, especially new health check types. To create a new plugin, please look at first into sample implementation of ```SampleCheck```. This check is responsible for comparing the current day with a "green day" provided in configuration.

See: [sample_check.go](https://github.com/mtojek/greenwall/blob/master/middleware/healthcheck/checks/sample_check.go)

High priority health check plugins:
* TCP
* DNS 
* REST
* SOAP

Please open a PR once you finish the implementation. Don't worry - I'll help you in pushing your change to the repository!

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
