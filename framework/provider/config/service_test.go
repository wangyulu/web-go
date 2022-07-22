package config

import (
	"github.com/wangyulu/web-go/framework"
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/provider/app"
	"github.com/wangyulu/web-go/framework/provider/env"
	"path/filepath"
	"testing"

	tests "github.com/wangyulu/web-go/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_GetInt(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		container := framework.NewHadeContainer()
		basePath := tests.BasePath
		container.Bind(&app.HadeAppProvider{basePath})
		container.Bind(&env.HadeEnvProvider{})
		container.Bind(&HadeConfigProvider{})

		envService := container.MustMake(contract.EnvKey).(contract.Env)

		folder := filepath.Join(basePath, "config", envService.AppEnv())
		serv, err := NewHadeConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*HadeConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
