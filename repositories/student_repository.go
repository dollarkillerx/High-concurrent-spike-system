/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 上午10:35
* */
package repositories

import "High-concurrent-spike-system/datamodels"

type StudentRepository interface {
	GetStudentData() *datamodels.Student
}

type StudentManger struct {

}

func NewStudentManger() StudentRepository {
	return &StudentManger{}
}

func (s *StudentManger) GetStudentData() *datamodels.Student {
	student := &datamodels.Student{
		Name:"哆啦酱",
		Age:18,
		Sex:datamodels.Stu_Male,
		Id:1,
	}
	return student
}

