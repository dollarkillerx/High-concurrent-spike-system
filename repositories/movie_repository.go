/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-16
* Time: 下午11:03
* */
package repositories

import "High-concurrent-spike-system/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	// 模拟赋值个给模型
	movie := &datamodels.Movie{Name: "多拉A梦"}
	return movie.Name
}