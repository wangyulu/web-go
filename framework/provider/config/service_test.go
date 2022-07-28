package config

import (
	"path/filepath"
	"testing"

	"github.com/wangyulu/web-go/framework/contract"
	tests "github.com/wangyulu/web-go/test"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_GetInt(t *testing.T) {
	container := tests.InitBaseContainer()

	Convey("test hade env normal case", t, func() {
		appService := container.MustMake(contract.AppKey).(contract.App)
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		folder := filepath.Join(appService.ConfigFolder(), envService.AppEnv())

		serv, err := NewHadeConfig(container, folder, map[string]string{})
		So(err, ShouldBeNil)
		conf := serv.(*HadeConfig)
		timeout := conf.GetString("database.default.timeout")
		So(timeout, ShouldEqual, "10s")
	})
}

func TestHadeConfig_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&HadeConfigProvider{})

	Convey("test hade config load", t, func() {
		db := &struct {
			ConnMaxLifetime string `yaml:"conn_max_lifetime"`
			ConnMaxIdle     int    `yaml:"conn_max_idle"`
		}{}

		configService := container.MustMake(contract.ConfigKey).(contract.Config)

		err := configService.Load("database", db)

		So(err, ShouldBeNil)
		So(db.ConnMaxIdle, ShouldEqual, 10)
	})
}
