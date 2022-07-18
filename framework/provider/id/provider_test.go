package id

import (
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/provider/app"
	"github.com/wangyulu/web-go/framework/provider/config"
	"github.com/wangyulu/web-go/framework/provider/env"
	"github.com/wangyulu/web-go/framework/util"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test hade console log normal case", t, func() {
		basePath := util.GetExecDirectory()
		c := framework.NewHadeContainer()
		c.Singleton(&app.HadeAppProvider{BasePath: basePath})
		c.Singleton(&env.HadeEnvProvider{})
		c.Singleton(&config.HadeConfigProvider{})

		err := c.Singleton(&HadeIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		t.Log(xid)
		So(xid, ShouldNotBeEmpty)
	})
}
