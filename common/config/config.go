package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../../..")
)

//const ConfigName = "landlordConf"

var Config config

type config struct {
	Version string `yaml:"version"`

	Server struct {
		Ip   string `yaml:"ip"`
		Port string `yaml:"port"`
	}

	Landlords struct {
		MaxSecondsForEveryRound int64 `yaml:"max_seconds_for_every_round"`
	}

	Session struct {
		UserSessionKey string `yaml:"user_session_key"`
		Secret         string `yaml:"secret"`
		Name           string `yaml:"name"`
	}

	TLS struct {
		Addr string `yaml:"addr"`
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	}

	MySQL struct {
		Address       []string `yaml:"address"`
		UserName      string   `yaml:"username"`
		Password      string   `yaml:"password"`
		DatabaseName  string   `yaml:"database_name"`
		MaxOpenConns  int      `yaml:"max_open_conns"`
		MaxIdleConns  int      `yaml:"max_idle_conns"`
		MaxLifeTime   int      `yaml:"max_life_time"`
		LogLevel      int      `yaml:"log_level"`
		SlowThreshold int      `yaml:"slow_threshold"`
	}

	TokenPolicy struct {
		JwtSecret string `yaml:"jwt_secret"`
		JwtExpire int64  `yaml:"jwt_expire"`
	}

	Websocket struct {
		Port             []string `yaml:"port"`
		MaxConnNum       int      `yaml:"max_conn_num"`
		MaxMsgLen        int      `yaml:"max_msg_len"`
		HandshakeTimeOut int      `yaml:"handshake_time_out"`
		OnlineTimeOut    int      `yaml:"online_time_out"`
	}
}

func init() {
	unmarshalConfig(&Config, "config.yaml")
}

func unmarshalConfig(config interface{}, configName string) {

	env := "CONFIG_NAME"

	cfgName := os.Getenv(env)

	if len(cfgName) != 0 {
		bytes, err := os.ReadFile(filepath.Join(cfgName, "config", configName))
		if err != nil {
			bytes, err = os.ReadFile(filepath.Join(Root, "config", configName))
			if err != nil {
				panic(err.Error() + " config: " + filepath.Join(cfgName, "config", configName))
			}
		} else {
			Root = cfgName
		}
		if err = yaml.Unmarshal(bytes, config); err != nil {
			panic(err.Error())
		}
	} else {
		//bytes, err := os.ReadFile(fmt.Sprintf("../config/%s", configName))
		bytes, err := os.ReadFile(fmt.Sprintf("config/%s", configName))
		if err != nil {
			panic(err.Error() + configName)
		}
		if err = yaml.Unmarshal(bytes, config); err != nil {
			panic(err.Error())
		}
	}
}
