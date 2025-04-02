package route

import "github.com/gin-gonic/gin"

type IAPI interface {
	SetupRoute(*gin.RouterGroup)
}

type IAdminAPI interface {
	SetupAdminRoute(group *gin.RouterGroup)
}

type IInternalAPI interface {
	SetupInternalRoute(group *gin.RouterGroup)
}

type IAPIGroup interface {
	RegisterAPIs(*gin.RouterGroup)
}

type IAdminAPIGroup interface {
	RegisterAdminAPIs(*gin.RouterGroup)
}

type IInternalAPIGroup interface {
	RegisterInternalAPIs(*gin.RouterGroup)
}

func RegisterAPI(router gin.IRouter, api IAPI, apiPath string) {
	rg := router.Group(apiPath)
	api.SetupRoute(rg)
}

func RegisterAdminAPI(router gin.IRouter, api IAdminAPI, apiPath string) {
	rg := router.Group(apiPath)
	api.SetupAdminRoute(rg)
}

func RegisterInternalAPI(router gin.IRouter, api IInternalAPI, apiPath string) {
	rg := router.Group(apiPath)
	api.SetupInternalRoute(rg)
}

func RegisterAPIGroup(router gin.IRouter, apiGroup IAPIGroup, apiGroupPath string) {
	rg := router.Group(apiGroupPath)
	apiGroup.RegisterAPIs(rg)
}

func RegisterAdminAPIGroup(router gin.IRouter, apiGroup IAdminAPIGroup, apiGroupPath string) {
	rg := router.Group(apiGroupPath)
	apiGroup.RegisterAdminAPIs(rg)
}

func RegisterInternalAPIGroup(router gin.IRouter, apiGroup IInternalAPIGroup, apiGroupPath string) {
	rg := router.Group(apiGroupPath)
	apiGroup.RegisterInternalAPIs(rg)
}
