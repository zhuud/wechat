package config

import (
    "testing"

    "github.com/davecgh/go-spew/spew"
)

func Test_mustLoadWeCom(t *testing.T) {

    wcList := mustLoadWeCom()

    spew.Dump(wcList)
}
