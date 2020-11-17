package main

type Task struct {
	Id int `json:"id"`
	Group_id int `json:"group_id"`
	Tsk string `json:"task"`
	Completed bool `json:"completed"`
	Created string `json:"created_at"`
	Complited_at string `json:"complited_at"`
}
type Group struct {
	Name string `json:"group_name"`
	Description string `json:"group_description"`
	Id int `json:"group_id"`
	Parent_id int `json:"parent_id"`
}
type Changes struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
type Server struct {
	Tasks []Task `json:"tasks"`
	Groups []Group `json:"groups"`
}
type Time struct {
	year int `json:"year"`
	month int `json:"month"`
	day int  `json:"day"`
}
func (t Time) new(year int,month int,day int) Time {
	t.day = day
	t.month = month
	t.year = year
	return t
}
type GiveGroups struct {
	Groups []Group `json:"groups"`
}
type GiveTasks struct{
	Tasks []Task `json:"tasks"`
}