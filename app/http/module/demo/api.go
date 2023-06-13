/**
 * @Author: yy
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2023/06/08 13:37
 */

package demo

import (
	demoService "github.com/Joker-desire/simple/app/provider/demo"
	"github.com/Joker-desire/simple/framework/contract"
	"github.com/Joker-desire/simple/framework/gin"
)

type ApiDemo struct {
	service *UserService
}

func NewApiDemo() *ApiDemo {
	service := NewUserService()
	return &ApiDemo{service: service}
}

func Register(r *gin.Engine) error {
	api := NewApiDemo()
	_ = r.Bind(&demoService.ServiceProviderDemo{})
	d := r.Group("/demo")
	{
		d.GET("/demo", api.Demo)
		d.GET("/demo2", api.Demo2)
		d.GET("/demo3", api.Demo3)
		d.POST("/demo", api.DemoPost)
	}
	return nil
}

func (a *ApiDemo) Demo(c *gin.Context) {
	users := a.service.GetUsers()
	dtOs := UserModelsToUserDTOs(users)
	c.JSON(200, dtOs)
}

func (a *ApiDemo) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.Key).(demoService.IService)
	students := demoProvider.GetAllStudent()
	dtOs := StudentsToUserDTOs(students)
	c.JSON(200, dtOs)
}

func (a *ApiDemo) Demo3(c *gin.Context) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	baseFolder := appService.BaseFolder()
	c.JSON(200, baseFolder)

}

func (a *ApiDemo) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(foo)
	if err != nil {
		_ = c.AbortWithError(500, err)
	}
	c.JSON(200, foo)
}
