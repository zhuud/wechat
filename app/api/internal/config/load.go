package config

import (
    "github.com/zeromicro/go-zero/core/conf"
)

func MustLoad(path string) Config {
    c := Config{}
    conf.MustLoad(path, &c)

    return c
}
