// supervisord-cfg-gen generates supervisord config file to manage demo services
// Generates config.conf in supervisord INI format
//
// supervisord-cfg-gen 生成 supervisord 配置文件用于管理演示服务
// 生成 INI 格式的 config.conf 配置文件
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/orzkratos/supervisordkratos"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

const (
	serviceUsername = "lele"              // Unix username to run services // 运行服务的 Unix 用户名
	logRootPath     = "/var/log/services" // Root path to store service logs // 存储服务日志的根路径
	separatorLine   = 80                  // Dividing mark width // 分隔标记宽度
)

// main generates supervisord config file with demo service configs
// main 生成包含演示服务配置的 supervisord 配置文件
func main() {
	demo1Config := supervisordkratos.NewProgramConfig(
		"demo1kratos",
		osmustexist.ROOT(runpath.PARENT.UpTo(1, "demo1kratos")),
		serviceUsername,
		logRootPath,
	).WithStartRetries(3).WithAutoRestart(true)

	demo2Config := supervisordkratos.NewProgramConfig(
		"demo2kratos",
		osmustexist.ROOT(runpath.PARENT.UpTo(1, "demo2kratos")),
		serviceUsername,
		logRootPath,
	).WithStartRetries(3).WithAutoStart(true)

	microservicesGroup := supervisordkratos.NewGroupConfig("microservices").
		AddProgram(demo1Config).
		AddProgram(demo2Config)

	configContent := supervisordkratos.GenerateGroupConfig(microservicesGroup)

	// Print generated config to console with dividing marks
	// 在控制台打印生成的配置，使用分隔标记
	zaplog.SUG.Debug("Generated supervisord config content")
	fmt.Println(strings.Repeat("#", separatorLine))
	fmt.Println(configContent)
	fmt.Println(strings.Repeat("#", separatorLine))

	// Write config to file in current app DIR
	// 将配置写入当前应用目录的文件
	configPath := runpath.PARENT.Join("config.conf")
	zaplog.LOG.Info("Writing supervisord config", zap.String("path", configPath))
	must.Done(os.WriteFile(configPath, []byte(configContent), 0644))
}
