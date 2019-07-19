package controller

import (
	"github.com/bbcyyb/pcrs/article"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	AService article.IService
}

func NewArticleController(service article.IService) *ArticleController {
	return &ArticleController{
		AService: service,
	}
}

func (a *ArticleController) Get(c *gin.Context) error {
	panic("implement me")
}

func (a *ArticleController) Create(c *gin.Context) error {
	panic("implement me")
}

func (a *ArticleController) Post(c *gin.Context) error {
	panic("implement me")
}

func (a *ArticleController) Delete(c *gin.Context) error {
	panic("implement me")
}
