# supervisord-cfg-gen

用于创建 supervisord INI 格式配置文件的命令行应用。

---

## 英文文档

[ENGLISH README](README.md)

## 项目简介

**supervisord-cfg-gen** 是一个命令行应用，用于生成 supervisord 配置文件来管理演示服务。它创建标准 supervisord INI 格式的 `config.conf` 文件，包含程序和组配置。

## 功能特性

- **自动配置生成** - 一条命令创建完整的 supervisord 配置
- **组管理** - 将多个程序组织到 microservices 组中
- **可自定义设置** - 配置服务用户名、日志路径和重启行为
- **INI 格式输出** - 标准 supervisord 配置格式

## 使用方法

### 构建和运行

```bash
go build -o supervisord-cfg-gen .
./supervisord-cfg-gen
```

应用会在当前目录生成 `config.conf` 文件。

### 生成的配置

生成的配置包括：

- **microservices 组** - 包含 demo1kratos 和 demo2kratos 程序
- **demo1kratos 程序** - 启用自动重启，3 次启动重试
- **demo2kratos 程序** - 启用自动启动，3 次启动重试

### 配置内容

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

## 配置选项

编辑 `main.go` 中的常量来自定义：

```go
const (
    serviceUsername = "lele"              // 运行服务的 Unix 用户名
    logRootPath     = "/var/log/services" // 存储服务日志的根路径
    separatorLine   = 80                  // 分隔标记宽度
)
```

## 依赖项

- [github.com/orzkratos/supervisordkratos](https://github.com/orzkratos/supervisordkratos) - Supervisord 配置生成
- [github.com/yyle88/runpath](https://github.com/yyle88/runpath) - 运行时路径操作
- [github.com/yyle88/zaplog](https://github.com/yyle88/zaplog) - 日志支持

## 配合 supervisord 使用

生成 `config.conf` 后，在 supervisord 中使用：

```bash
# 在主 supervisord.conf 中引入
[include]
files = /path/to/supervisord-cfg-gen/config.conf

# 作为主配置使用
supervisord -c config.conf
```

## 项目说明

本项目是 [supervisordkratos-demos](https://github.com/orzkratos/supervisordkratos-demos) 的子项目，用于演示如何使用 supervisordkratos 包生成 supervisord 配置文件。

更多信息请参考主项目文档。
