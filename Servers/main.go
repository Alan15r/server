package main

import (
	"github.com/labstack/echo"
)
var server = new(Server)
func main() {
	e := echo.New()
	e.POST("/group/new", addNewGroup)
	e.GET("/group/top_parents", showTopGroup)
	e.GET("/group/:id", showGroup)
	e.GET("/group/childs/:id", showChilds)
	e.DELETE("/DELETE/group/:id", deleteGroup)
	e.POST("/task/new", addTask)
	e.GET("/groups", showGroups)
	e.GET("/tasks/group/:id", tasksGroup)
	//e.POST("/tasks/:id", changeTask)
	//e.GET("/stat/{interval}",Statistics)
	e.Logger.Fatal(e.Start(":1323"))
}
