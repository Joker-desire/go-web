/**
 * @Author: yy
 * @Description:
 * @File:  provider_test
 * @Version: 1.0.0
 * @Date: 2023/07/03 13:52
 */

package config

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/provider/app"
	"github.com/Joker-desire/simple/framework/provider/env"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSimpleConfig_GetInt(t *testing.T) {
	Convey("test simple config get int", t, func() {
		container := framework.NewSimpleContainer()

		sp := &app.SimpleAppProvider{BaseFolder: "/Users/yinyang/Projects/golang/go-web"}
		err := container.Bind(sp)
		sp1 := &env.SimpleEnvProvider{Folder: "/Users/yinyang/Projects/golang/go-web"}
		err = container.Bind(sp1)
		sp2 := &SimpleConfigProvider{}
		err = container.Bind(sp2)

		confService := container.MustMake(contract.ConfigKey).(contract.Config)
		timeout := confService.GetInt("database.mysql.port")
		So(timeout, ShouldEqual, 3306)
		So(err, ShouldBeNil)
	})
}
