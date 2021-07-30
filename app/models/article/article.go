package article

import (
	"strconv"

	"github.com/ruoxiao-zh/goblog/app/models"
	"github.com/ruoxiao-zh/goblog/pkg/route"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
