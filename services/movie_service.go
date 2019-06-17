/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午11:09
* */
package services

import (
	"High-concurrent-spike-system/repositories"
	"fmt"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManger struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManger(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManger{
		repo:repo,
	}
}

func (m *MovieServiceManger) ShowMovieName() string {
	return fmt.Sprintln("获取到的数据:" + m.repo.GetMovieName())
}