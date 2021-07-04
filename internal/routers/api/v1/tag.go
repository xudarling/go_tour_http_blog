package v1

import (
	"github.com/gin-gonic/gin"
	"go_tour/blog_service/pkg/app"
	"go_tour/blog_service/pkg/errcode"
)

type Tag struct{}

func NewTag()Tag{
	return Tag{}
}

func (t Tag)Get(c *gin.Context){

}
func (t Tag) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}