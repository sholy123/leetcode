package main

import "fmt"

type Animail interface {
	Running()
}
type dog struct{
	name string
}
type cat struct{
	name string
}
func (d dog)Running(){
	fmt.Println("this is a dog run")
}
func (c cat)Running(){
	fmt.Println("this is a cat run")
}
func (c *dog)Set(name string){
	c.name = name
}
func main(){
	a := Animail(dog{})
	a.Running()
	a = cat{}
	a.Running()
	d := dog{"zhangsan"}
	d.Set("lisi")
	fmt.Println(d.name)
}