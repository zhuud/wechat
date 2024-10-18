package externalcontactway

import (
	"context"
	"fmt"
	"testing"
	"time"

	"rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/fx"
)

func Test(t *testing.T) {
	newSyncExternalContactWayCmd(context.Background(), svc.NewServiceContext()).Do([]string{})
}

func TestFx(t *testing.T) {
	listFunc := func(source chan<- any) {
		for k := 0; k < 10; k++ {
			source <- k
			time.Sleep(1 * time.Second)
		}
	}

	showFunc := func(item any) any {
		fmt.Println(item)
		return item
	}

	lastFunc := func(item any) {
		fmt.Println("last", item)
	}

	fx.From(listFunc).Map(showFunc).Parallel(lastFunc)
}
