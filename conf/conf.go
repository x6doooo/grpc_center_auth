package conf

import (
    "flag"
    "github.com/BurntSushi/toml"
)

const (
    ENV_MODE_DEV = "debug"
    ENV_MODE_PROD = "release"
)

type EnvConf struct {
    Mode string
}

type LogConf struct {
    File    string
    MaxLine int `toml:"max_line"`
    Backups int
}

type MysqlConf struct {
    Username string
    Password string
    Addr     string
    DbName   string `toml:"db_name"`
    MaxConns int `toml:"max_conns"`
    MaxIdle  int `toml:"max_idle"`
}

type PswdConf struct {
    Key  string
    Iv   string
    Salt string
}

type ServerConf struct {
    Addr string
}

type MainConf struct {
    Env    EnvConf
    Log    LogConf
    Mysql  MysqlConf
    Pswd   PswdConf
    Server ServerConf
}

var (
    confFile string
    Conf = &MainConf{}
    IsDebugMode = false
)

func init() {
    flag.StringVar(&confFile, "conf", "/path/file.toml", "config file")
    flag.Parse()

    if confFile == "" {
        panic("need args --conf=/path/file")
    }
    toml.DecodeFile(confFile, Conf)

    IsDebugMode = Conf.Env.Mode == ENV_MODE_DEV
}
