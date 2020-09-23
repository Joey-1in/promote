package repositories

import (
	"promote/irisdemo/datamodels"
)

type MovieRepository interface {
	GetMoveName() string
}

type MovieManager struct {
}

func (m *MovieManager) GetMoveName() string {
	movie := &datamodels.Movie{Name: "活动促销"}
	return movie.Name
}

// 工厂模式
func NewMovieManager() MovieRepository {
	return &MovieManager{}
}
