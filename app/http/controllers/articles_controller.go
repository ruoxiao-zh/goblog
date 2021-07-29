package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ruoxiao-zh/goblog/app/models/article"
	"github.com/ruoxiao-zh/goblog/pkg/logger"
	"github.com/ruoxiao-zh/goblog/pkg/route"
	"github.com/ruoxiao-zh/goblog/pkg/types"
	"gorm.io/gorm"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

// Show 文章详情页面
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的文章数据
	articleInfo, err := article.Get(id)

	// 3. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 4. 读取成功，显示文章
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)
		tmpl.Execute(w, articleInfo)
	}
}

// Index 文章列表页
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 1. 获取结果集
	articles, err := article.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		// 2. 加载模板
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)

		// 3. 渲染模板，将所有文章的数据传输进去
		tmpl.Execute(w, articles)
	}
}
