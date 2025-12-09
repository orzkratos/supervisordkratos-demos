# supervisord-cfg-gen

Command-line app to create supervisord configuration files in INI format.

---

## CHINESE README

[中文说明](README.zh.md)

## Overview

**supervisord-cfg-gen** is a command-line app that generates supervisord configuration files to manage demo services. It creates `config.conf` in standard supervisord INI format with program and group configurations.

## Features

- **Auto Config Generation** - Creates complete supervisord config with one command
- **Group Management** - Organizes multiple programs into microservices group
- **Customizable Settings** - Configure service username, log paths, and restart behaviors
- **INI Format Output** - Standard supervisord configuration format

## Usage

### Build and Run

```bash
go build -o supervisord-cfg-gen .
./supervisord-cfg-gen
```

The app generates `config.conf` in the current DIR.

### Generated Config

The generated config includes:

- **microservices group** - Contains demo1kratos and demo2kratos programs
- **demo1kratos program** - Auto restart enabled, 3 start retries
- **demo2kratos program** - Auto start enabled, 3 start retries

### Config Content

```ini
[group:microservices]
programs=demo1kratos,demo2kratos

[program:demo1kratos]
user            = lele
directory       = /path/to/demo1kratos
command         = /path/to/demo1kratos/bin/demo1kratos
autorestart     = true
startretries    = 3
stdout_logfile  = /var/log/services/demo1kratos.log
stderr_logfile  = /var/log/services/demo1kratos.err

[program:demo2kratos]
user            = lele
directory       = /path/to/demo2kratos
command         = /path/to/demo2kratos/bin/demo2kratos
autostart       = true
startretries    = 3
stdout_logfile  = /var/log/services/demo2kratos.log
stderr_logfile  = /var/log/services/demo2kratos.err
```

## Configuration

Edit `main.go` constants to customize:

```go
const (
    serviceUsername = "lele"              // Unix username to run services
    logRootPath     = "/var/log/services" // Root path to store service logs
    separatorLine   = 80                  // Dividing mark width
)
```

## Dependencies

- [github.com/orzkratos/supervisordkratos](https://github.com/orzkratos/supervisordkratos) - Supervisord config generation
- [github.com/yyle88/runpath](https://github.com/yyle88/runpath) - Runtime path operations
- [github.com/yyle88/zaplog](https://github.com/yyle88/zaplog) - Logging support

## Use with supervisord

Once `config.conf` is generated, use it with supervisord:

```bash
# Include in main supervisord.conf
[include]
files = /path/to/supervisord-cfg-gen/config.conf

# Use it as main config
supervisord -c config.conf
```

## About

This is a subproject of [supervisordkratos-demos](https://github.com/orzkratos/supervisordkratos-demos), demonstrating how to use the supervisordkratos package to generate supervisord config files.

See the main project docs to get more info.
