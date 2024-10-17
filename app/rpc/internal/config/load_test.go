package config

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/zhuud/go-library/svc/conf"
)

func Test_mustLoadWeCom(t *testing.T) {
	wcList := mustLoadWeCom()
	spew.Dump(wcList)
}

func Test_GetConf(t *testing.T) {
	var c any
	err := conf.GetUnmarshal(fmt.Sprintf("/qconf/web-config/%s", "kafka_cluster"), &c)
	spew.Dump(c, err)
}

func Test_GetMysql(t *testing.T) {
	c := Config{}
	m := mustLoadMysql(c, "db_passport")
	spew.Dump(m)
}
