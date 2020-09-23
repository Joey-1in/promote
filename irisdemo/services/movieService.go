package services

import (
	"promote/irisdemo/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManger struct {
	reposit repositories.MovieRepository
}

func (m *MovieServiceManger) ShowMovieName() string {
	return "我们获取的视频名称为：" + m.reposit.GetMoveName()
}

func NewMovieServiceManger(reposit repositories.MovieRepository) MovieService {
	return &MovieServiceManger{reposit: reposit}
}
