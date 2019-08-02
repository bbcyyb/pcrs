package article

import (
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	ArticleService IService
}

func NewArticleController(service IService) *ArticleController {
	return &ArticleController{
		ArticleService: service,
	}
}

func (a *ArticleController) Get(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleController) Create(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleController) GetById(c *gin.Context) {
	panic("implement me")
}

func (a *ArticleController) Delete(c *gin.Context) {
	panic("implement me")
}
