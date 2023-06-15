/**
 * @Author: yy
 * @Description:
 * @File:  provide_test
 * @Version: 1.0.0
 * @Date: 2023/06/13 13:06
 */

package env

import (
	"github.com/Joker-desire/simple/framework"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/provider/app"
	"testing"
)

func TestSimpleEnvProvider(t *testing.T) {
	container := framework.NewSimpleContainer()
	sp := &app.SimpleAppProvider{BaseFolder: "/Users/yinyang/Projects/golang/go-web"}

	err := container.Bind(sp)
	sp2 := &SimpleEnvProvider{}
	err = container.Bind(sp2)

	envService := container.MustMake(contract.EnvKey).(contract.Env)
	t.Log("-----APP_ENV:", envService.Get("APP_ENV"))
	t.Log("-----name:", envService.Get("name"))
	if err != nil {
		t.Fatal(err)
	}

}
