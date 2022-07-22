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
		c.Bind(&app.HadeAppProvider{BaseFolder: basePath})
		c.Bind(&env.HadeEnvProvider{})
		c.Bind(&config.HadeConfigProvider{})

		err := c.Bind(&HadeIDProvider{})
		So(err, ShouldBeNil)

		idService := c.MustMake(contract.IDKey).(contract.IDService)
		xid := idService.NewID()
		t.Log(xid)
		So(xid, ShouldNotBeEmpty)
	})
}
