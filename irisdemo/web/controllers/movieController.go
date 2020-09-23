package controllers

import (
	"promote/irisdemo/repositories"
	"promote/irisdemo/services"

	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (m *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManger(movieRepository)
	movieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieResult,
	}
}
