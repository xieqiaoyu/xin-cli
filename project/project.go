package project

type Project struct {
	BuildPath string
	TArgs     interface{} // 用于模板渲染的参数对象
}
