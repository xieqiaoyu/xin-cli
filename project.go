package main

type Project struct {
	BuildPath string
	tArgs     interface{} // 用于模板渲染的参数对象
}
