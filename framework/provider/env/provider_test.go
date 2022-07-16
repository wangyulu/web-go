package env

import (
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/provider/app"
	"testing"

	tests "github.com/wangyulu/web-go/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeEnvProvider(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		sp := &app.HadeAppProvider{BaseFolder: basePath}

		err := c.Bind(sp)
		So(err, ShouldBeNil)

		sp2 := &HadeEnvProvider{}
		err = c.Bind(sp2)
		So(err, ShouldBeNil)

		envServ := c.MustMake(contract.EnvKey).(contract.Env)
		So(envServ.AppEnv(), ShouldEqual, "development")
		// So(envServ.Get("DB_HOST"), ShouldEqual, "127.0.0.1")
		// So(envServ.AppDebug(), ShouldBeTrue)
	})
}
