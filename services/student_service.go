/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 上午10:41
* */
package services

import (
	"High-concurrent-spike-system/datamodels"
	"High-concurrent-spike-system/repositories"
)

type StudentService interface {
	ShowData() *datamodels.Student
}

type StudentServiceManger struct {
	dao repositories.StudentRepository
}

func NewStudentServiceMange(dao repositories.StudentRepository) StudentService {
	return &StudentServiceManger{
		dao:dao,
	}
}

func (s *StudentServiceManger) ShowData() *datamodels.Student {
	return s.dao.GetStudentData()
}

