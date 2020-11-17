package main

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func addNewGroup(c echo.Context) error {
	g := new(Group)
	err := c.Bind(g)
	if err != nil {
		return err
	}
	newGroup,er := server.AddGroup(*g)
	if er != nil {
		return echo.NewHTTPError(http.StatusBadRequest, er.Error())
	}
	return c.JSON(http.StatusCreated, newGroup)
}
func showTopGroup(c echo.Context) error {
     give, err := server.ShowParentGroups()
     if err!=nil{
     	echo.NewHTTPError(http.StatusNotFound, err)
	 }
	 return c.JSON(http.StatusOK,give)
}
func showGroup(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Bad Request").Error())
	}
	give,er := server.ShowGroupID(id)
	if er!=nil{
		return echo.NewHTTPError(http.StatusNotFound, er.Error())
	}
	return c.JSON(http.StatusOK,give)
}
func showChilds(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Bad Request").Error())
	}
    give, er:= server.ShowChilds(id)
	if er!=nil{
		return echo.NewHTTPError(http.StatusNotFound, er.Error())
	}
	return c.JSON(http.StatusOK,give)
}
func deleteGroup(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Bad Request").Error())
	}
	Give, er := server.GroupDelete(id)
	if er!= nil{
		return echo.NewHTTPError(http.StatusNotFound, er.Error())
	}
	return c.JSON(http.StatusOK, Give)
}
func addTask(c echo.Context) error {
	t := new(Task)
	err := c.Bind(t)
	if err != nil {
		return err
	}
	newTask, er := server.AddTask(*t)
	if er!=nil {
		return echo.NewHTTPError(http.StatusNotFound,er.Error())
	}
	return c.JSON(http.StatusOK, newTask)
}
func showGroups(c echo.Context) error {
	sort := c.FormValue("sort")
	limit := c.FormValue("limit")
    give, err := server.ShowGroups(sort, limit)
    if err!=nil {
    	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, give)
}
func tasksGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	give, er := server.ShowTasks(id)
	if er != nil {
		return echo.NewHTTPError(http.StatusNotFound, er.Error())
	}
	return c.JSON(http.StatusOK, give)
}
/*func changeTask(c echo.Context) error {
	ch := new(Changes)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.Bind(ch)
	if err!=nil {
		return err
	}
}*/