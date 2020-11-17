package main

import (
	"errors"
	"strconv"
	"time"
)
//adding a group
func(server *Server) AddGroup(group Group) (Group,error) {
	err := checkName(group)
	if err != nil {
		return group,err
	}
	err = checkParentId(group)
	if err != nil {
		return group, err
	}
	group.Id = len(server.Groups)+1
	server.Groups = append(server.Groups, group)
	return group, nil
}
//show the parent group
func(server *Server) ShowParentGroups() (GiveGroups, error) {
	var Give GiveGroups
	for _,v:= range server.Groups{
		if v.Parent_id == 0 {
			Give.Groups = append(Give.Groups, v)
		}
	}
	if len(Give.Groups) == 0 {
		return Give, errors.New("group not found")
	}
	return Give, nil
}
//show group by ID
func(server *Server) ShowGroupID(id int) (GiveGroups,error) {
	var Give GiveGroups
	for _,v:= range server.Groups{
		if v.Id == id {
			Give.Groups = append(Give.Groups, v)
		}
	}
	if len(Give.Groups) == 0 {
		return Give, errors.New("group not found")
	}
	return Give, nil
}
//show subgroups groups by ID
func (server *Server) ShowChilds(id int) (GiveGroups, error) {
	var Give GiveGroups
	err := checkID(id)
	if err!= nil{
		return Give, err
	}
	for _, v := range server.Groups {
		if v.Parent_id == id {
			Give.Groups = append(Give.Groups, v)
		}
	}
	if len(Give.Groups) == 0 {
		return Give, errors.New("childs group not found")
	}
	return Give, nil
}
//show subgroups groups by ID
func (server *Server) GroupDelete(id int) (GiveGroups,error) {
	var give GiveGroups
	err:= checkID(id)
	if err!=nil{
		return give, err
	}
	for i,v := range server.Groups {
		if v.Id == id {
			server.Groups = append(server.Groups[:i], server.Groups[i+1:]...)
		} else {
			give.Groups = append(give.Groups, v)
		}
	}
	return give, nil
}
//adding a task
func (server *Server) AddTask(task Task) (Task, error) {
	err := checkTaskName(task.Tsk)
	if err != nil {
		return task, err
	}
	err = checkID(task.Group_id)
	if err != nil {
		return task, err
	}
	task.Created = time.Now().String()
	server.Tasks = append(server.Tasks, task)
	return task, nil
}
func (server *Server) ShowGroups(sort string, limit string) (GiveGroups, error) {
	var Give GiveGroups
	var err error
	var lim int
	if limit != ""{
		lim, err = strconv.Atoi(limit)
		if err!=nil {
			return Give, err
		}
	} else {
		lim = 0
	}
	if lim == 0 || lim > len(server.Groups) {
		lim = len(server.Groups)
	}
	switch sort {
	case "name":
		{
			Give = GiveByName(lim)
		}
	case "parents_first":
		{
			Give = GiveByParentFirst(lim)
		}
	case "parent_with_childs":
		{
			Give = GiveByParentWithChilds(lim)
		}
	case "":
		{
			Give = GiveFul(lim)
		}
	default:
		err = errors.New("invalid command input")
	}
	return Give, err
}
func (server *Server) ShowTasks(id int) (GiveTasks,error) {
	var give GiveTasks
	err := checkID(id)
	if err!=nil {
		return give,err
	}
	for _, v := range server.Tasks {
		if v.Group_id == id {
			give.Tasks = append(give.Tasks, v)
		}
	}
	return give, nil
}
/*func (server *Server) ChangeTask(id int, body Changes) (GiveTasks, error) {
	var Give GiveTasks
	err:= checkID(id)
	if err!=nil {
		return Give, err
	}
} */
