package main

type aa interface {
	m1()
}

type bb interface {
	m2()
}

//接口嵌套
type cc interface {
	aa
	bb
}

type myinter struct {

}

func (m myinter) m1() {

}

func (m myinter) m2() {

}


func main() {

}
