package main

import (
	"errors"
	"sort"
)

func checkParentId (group Group) error {
	switch group.Parent_id {
	case 0:
		{
			return nil
		}
	default:
		{
			for _,v := range server.Groups{
				if v.Id == group.Parent_id {
					return nil
				}
			}
			return errors.New("there is no group with this ID")
		}
	}

}
func checkName(group Group) error {
	for _,v := range server.Groups{
		if v.Name == group.Name {
			return errors.New("name used")
		}
	}
	return nil
}
func checkID(id int) error {
	for _,v:=range server.Groups{
		if v.Id == id{
			return nil
		}
	}
	return errors.New("group not found")
}
func checkTaskName(name string) error {
	for _,v := range server.Tasks {
		if v.Tsk == name{
			return errors.New("name used")
		}
	}
	return nil
}
func GiveByName(lim int) GiveGroups {
	var limit int
	var Give GiveGroups
	sort.SliceStable(server.Groups, func(i, j int) bool {
		return server.Groups[i].Name < server.Groups[j].Name
	})
	for _, v := range server.Groups {
		Give.Groups = append(Give.Groups, v)
		limit++
		if limit == lim {
			break
		}
	}
	return Give
}
func GiveByParentFirst(lim int) GiveGroups {
	var Give GiveGroups
	var limit int
	for _, v := range server.Groups {
		if v.Parent_id == 0 {
			Give.Groups = append(Give.Groups, v)
			limit++
		}
		if limit == lim {
			break
		}
	}
	for _, v := range server.Groups {
		if v.Parent_id != 0 {
			Give.Groups = append(Give.Groups, v)
			limit++
		}
		if limit == lim {
			break
		}
	}
	return Give
}
func GiveByParentWithChilds(lim int) GiveGroups {
	var limit int
	var Give GiveGroups
	for _, v := range server.Groups {
		if v.Parent_id == 0 {
			Give.Groups = append(Give.Groups, v)
			limit++
			if limit == lim {
				break
			}
		}
		for _, u := range server.Groups {
			if u.Parent_id == v.Id {
				Give.Groups = append(Give.Groups, v)
				limit++
				if limit == lim {
					break
				}
			}
		}
		if limit == lim {
			break
		}
	}
	return Give
}
func GiveFul(lim int) GiveGroups {
	var limit int
	var Give GiveGroups
	for _, v := range server.Groups {
		Give.Groups = append(Give.Groups, v)
		limit++
		if limit == lim {
			break
		}
	}
	return Give
}