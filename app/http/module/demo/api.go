package demo

import (
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/gin"

	demoService "github.com/wangyulu/web-go/app/provider/demo"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.GET("/demo/demo3", api.Demo3)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取所有用户
// @description.markdown demo.md
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	c.JSON(200, usersDTO)
}

// Demo2  for godoc
// @Summary 获取所有学生
// @Description 获取所有学生,不进行分页
// @Produce  json
// @Tags demo
// @Success 200 {array} UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

func (api *DemoApi) Demo3(c *gin.Context) {
	// 获取password
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	password := configService.GetString("database.mysql.password")
	username := configService.GetString("database.mysql.username")

	logger := c.MustMakeLog()
	logger.Info(c, "demo test error", map[string]interface{}{
		"api":  "demo/demo",
		"user": "jianfengye",
	})

	// 打印出来
	c.JSON(200, []string{password, username})
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}
