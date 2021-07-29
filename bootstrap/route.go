package bootstrap

import (

	"github.com/gorilla/mux"
	"github.com/ruoxiao-zh/goblog/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	return router
}
