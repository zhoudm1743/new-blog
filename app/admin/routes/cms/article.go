package cms

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/service/cms"
	"new-blog/app/admin/types"
	"new-blog/pkg/plugins/response"
	"new-blog/pkg/util"
)

type articleDep struct {
	fx.In
	Srv cms.ArticleService
}

func ArticleRoutes(t articleDep, r *types.AdminRouter) {
	api := r.Group("/cms/article")
	api.GET("/all", t.all)
	api.POST("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add)
	api.POST("/edit", t.edit)
	api.POST("/delete", t.delete)
}

func (t articleDep) all(c *gin.Context) {
	var listReq req.ArticleQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &listReq)) {
		return
	}
	res, err := t.Srv.All(listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t articleDep) list(c *gin.Context) {
	var page req.PageReq
	var listReq req.ArticleQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &page, &listReq)) {
		return
	}
	res, err := t.Srv.List(page, listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t articleDep) detail(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	res, err := t.Srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

func (t articleDep) add(c *gin.Context) {
	var addReq req.ArticleAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &addReq)) {
		return
	}
	err := t.Srv.Add(addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t articleDep) edit(c *gin.Context) {
	var editReq req.ArticleEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t articleDep) delete(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	err := t.Srv.Del(idReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
