package jkconfig

import (
	. "goconfig/config"
	"jk/jklog"
	"strings"
)

type JKConfig struct {
	C *Config
}

var ownConfig JKConfig

func OpenConfig(path string) *Config {
	tmp, err := ReadDefault(path)
	if err != nil {
		jklog.L().Errorln("read error: ", err)
		return nil
	}
	ownConfig.C = tmp

	return tmp
}

func (c *JKConfig) getConfig() *Config {
	return c.C
}

func (c *JKConfig) getString(sect, subsect string) string {
	str, err := c.getConfig().String(sect, subsect)
	if err != nil {
		return ""
	} else {
		return str
	}
}

func GetCmdLineDir() string {
	return ownConfig.getString("Work", "cmdline")
}

func GetAppDir() string {
	return ownConfig.getString("Work", "appdir")
}

func GetErrorPage(err string) string {
	return ownConfig.getString("ErrorPage", err)
}

func GetLogLevel() string {
	return ownConfig.getString("Log", "level")
}

func GetLogFile() string {
	return ownConfig.getString("Log", "file")
}

func GetPicsBasicPath() string {
	return ownConfig.getString("Path", "imagebasic")
}

func GetFilesPath() string {
	return ownConfig.getString("Path", "FilesPath")
}

func GetDocsPath() string {
	return ownConfig.getString("Path", "docsPath")
}

func GetUseDB() bool {
	b, err := ownConfig.getConfig().Bool("Work", "usedb")
	if err != nil {
		// error, default return true
		return true
	}
	return b
}

func LogLevel() int {
	str := GetLogLevel()
	value := strings.Split(str, ",")
	level := jklog.LEVEL_CRITICAL | jklog.LEVEL_PANIC
	if len(value) == 0 {
		level |= jklog.LEVEL_MUCH
	}
	for _, v := range value {
		switch v {
		case "info":
			level |= jklog.LEVEL_INFO
		case "debug":
			level |= jklog.LEVEL_DEBUG
		case "notice":
			level |= jklog.LEVEL_NOTICE
		case "warn":
			level |= jklog.LEVEL_WARN
		case "error":
			level |= jklog.LEVEL_ERROR
		}
	}
	return level
}
