// Copyright (C) 2016 Flueg Lau. All rights reserved.
// Use of this source code is governed by a BSD-styleGNU General Public
// License that can be found in the LICENSE file.

// package config is used to maintain the service configurations saved in
// xml format file.
// Usage:
// 		myConf := config.MCSingleton()
// 		err := myConf.Load(confFile)
package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Singleton class.
// TODO: should support dynamically reload configuration file!!!
type XmlConfig struct {
	Service struct {
		// host:port that server will listen on
		EndPoint string

		ServiceReadTimeout  int // Request read timeout.
		ServiceWriteTimeout int // Response write timeout.
	}
	/*
		Access struct {
			RedisEndPoints []string // redis配置
		}

		// 后端服务
		BackEnd struct {
			// push管理器域名
			PushDomain   string
			PushEndPoint string
		}

		ForwardConf struct {
			Path string
			Addr string
		}
		// 功能开关
		Option struct {
			EnableVerbose  bool // 开启详细调试日志
			ForceCloseConn bool // 是否强制关闭连接
		}
	*/
}

var _MCInstance *XmlConfig

// function to get the singleton configuration object instance.
func MCSingleton() *XmlConfig {
	if _MCInstance == nil {
		_MCInstance = new(XmlConfig)
	}
	return _MCInstance
}

// Loading configurations from file.
func (conf *XmlConfig) Load(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Failed to open configuration file %s. error:%s", fileName, err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Failed to read data. error:%s", err)
	}

	err = xml.Unmarshal(data, conf)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal XmlConfig. err:%s", err)
	}

	return nil
}

// Reload configurations from file.
func (conf *XmlConfig) ReLoad() error {
	// TODO reload configurations from file.
	return nil
}

// Save configurations into file.
func (conf *XmlConfig) Save(fileName string) error {
	data, err := xml.MarshalIndent(conf, "  ", "	")
	if err != nil {
		return fmt.Errorf("Failed to marshal XmlConfig. error:%s", err)
	}

	err = ioutil.WriteFile(fileName, data, 0777)
	if err != nil {
		return fmt.Errorf("Failed to write data into %s. error:%s", fileName, err)
	}

	return nil
}
