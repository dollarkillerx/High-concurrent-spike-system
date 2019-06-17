/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 上午10:31
* */
package datamodels

const (
	Stu_Male = 1
	Stu_Woman = 2
)

type Student struct {
	Id int
	Name string
	Age int
	Sex int
}

