package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	Port      int    `yaml:"port"`
	ServerUrl string `yaml:"server_url"`
}

var GlobalConfig *Config

func ParseConfig(filename string) *Config {

	configPath := getConfigPath(filename)

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("\n\nNot found config file, %+v\n\n", err)
	}

	var cfg *Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("\n\nUnable to parse config file, %+v\n\n", err)
	}
	GlobalConfig = cfg
	return cfg
}


func getConfigPath(filename string) string {
	// 通过环境变量判断是否是开发环境
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("\n\nUnable to get executable path: %+v\n\n", err)
	}
	tmpDir := os.TempDir()
	// 临时目录下，为开发调试环境
	if strings.HasPrefix(executable, tmpDir) {
		_, mainFilePath, _, ok := runtime.Caller(0)

		if !ok {
			log.Fatalf("Unable to determine caller information")
		}

		// 定位 main.go 文件所在的目录, 即项目根目录
		projectRoot := filepath.Join(filepath.Dir(mainFilePath), "../")

		log.Println("root: ", mainFilePath, projectRoot)

		return filepath.Join(projectRoot, filename)
	}


	dir := filepath.Dir(executable)
	return filepath.Join(dir, filename)
}