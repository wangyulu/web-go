package ssh

import (
	"testing"

	"github.com/wangyulu/web-go/framework/provider/config"
	"github.com/wangyulu/web-go/framework/provider/log"

	. "github.com/smartystreets/goconvey/convey"
	tests "github.com/wangyulu/web-go/test"
)

func TestHadeSSHService_Load(t *testing.T) {
	container := tests.InitBaseContainer()
	container.Bind(&config.HadeConfigProvider{})
	container.Bind(&log.HadeLogServiceProvider{})

	Convey("test get client", t, func() {
		hadeRedis, err := NewHadeSSH(container)
		So(err, ShouldBeNil)
		service, ok := hadeRedis.(*HadeSSH)
		So(ok, ShouldBeTrue)
		client, err := service.GetClient(WithConfigPath("ssh.web-01"))
		So(err, ShouldBeNil)
		So(client, ShouldNotBeNil)
		session, err := client.NewSession()
		So(err, ShouldBeNil)
		out, err := session.Output("pwd")
		So(err, ShouldBeNil)
		So(out, ShouldNotBeNil)
		session.Close()
	})
}
