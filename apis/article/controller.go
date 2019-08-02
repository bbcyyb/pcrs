package article

import (
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	AService IService
}

func NewArticleController(service IService) *ArticleController {
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
