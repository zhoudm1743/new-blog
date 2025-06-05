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

type linkDep struct {
	fx.In
	Srv cms.LinkService
}

func LinkRoutes(t linkDep, r *types.AdminRouter) {
	api := r.Group("/cms/link")
	api.GET("/all", t.all)
	api.POST("/list", t.list)
	api.GET("/detail", t.detail)
	api.POST("/add", t.add)
	api.POST("/edit", t.edit)
	api.POST("/delete", t.delete)
}

func (t linkDep) all(c *gin.Context) {
	var listReq req.LinkQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &listReq)) {
		return
	}
	res, err := t.Srv.All(listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t linkDep) list(c *gin.Context) {
	var page req.PageReq
	var listReq req.LinkQueryReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &page, &listReq)) {
		return
	}
	res, err := t.Srv.List(page, listReq, req.GetAuth(c))
	response.CheckAndRespWithData(c, res, err)
}

func (t linkDep) detail(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	res, err := t.Srv.Detail(idReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

func (t linkDep) add(c *gin.Context) {
	var addReq req.LinkAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &addReq)) {
		return
	}
	err := t.Srv.Add(addReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t linkDep) edit(c *gin.Context) {
	var editReq req.LinkEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &editReq)) {
		return
	}
	err := t.Srv.Edit(editReq, req.GetAuth(c))
	response.CheckAndResp(c, err)
}

func (t linkDep) delete(c *gin.Context) {
	var idReq req.IdReq
	if response.IsFailWithResp(c, util.VerifyUtil.Verify(c, &idReq)) {
		return
	}
	err := t.Srv.Del(idReq.ID, req.GetAuth(c))
	response.CheckAndResp(c, err)
}
