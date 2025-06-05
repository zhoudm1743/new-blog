package cms

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"new-blog/app/admin/schemas/req"
	"new-blog/app/admin/schemas/resp"
	"new-blog/app/models"
	"new-blog/pkg/plugins/response"
)

type ArticleService interface {
	All(listReq req.ArticleQueryReq, auth *req.AuthReq) ([]resp.ArticleResp, error)
	List(page req.PageReq, listReq req.ArticleQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Detail(id uint) (resp.ArticleResp, error)
	Add(addReq req.ArticleAddReq, auth *req.AuthReq) error
	Edit(editReq req.ArticleEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type iArticleService struct {
	db *gorm.DB
}

func (iSrv iArticleService) All(listReq req.ArticleQueryReq, auth *req.AuthReq) ([]resp.ArticleResp, error) {
	chain := iSrv.db.Model(&models.CmsArticle{}).Order("id desc")
	if listReq.Title != "" {
		chain = chain.Where("title LIKE ?", "%"+listReq.Title+"%")
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	if listReq.IsTop != -1 {
		chain = chain.Where("is_top = ?", listReq.IsTop)
	}
	if listReq.IsComment != -1 {
		chain = chain.Where("is_comment = ?", listReq.IsComment)
	}
	if listReq.CategoryID != 0 {
		chain = chain.Where("category_id = ?", listReq.CategoryID)
	}
	var articles []models.CmsArticle
	if err := chain.Find(&articles).Error; err != nil {
		return nil, err
	}
	var res []resp.ArticleResp
	response.Copy(&res, &articles)
	return res, nil
}

func (iSrv iArticleService) List(page req.PageReq, listReq req.ArticleQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := iSrv.db.Model(&models.CmsArticle{}).Order("id desc")
	if listReq.Title != "" {
		chain = chain.Where("title LIKE ?", "%"+listReq.Title+"%")
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	if listReq.IsTop != -1 {
		chain = chain.Where("is_top = ?", listReq.IsTop)
	}
	if listReq.IsComment != -1 {
		chain = chain.Where("is_comment = ?", listReq.IsComment)
	}
	if listReq.CategoryID != 0 {
		chain = chain.Where("category_id = ?", listReq.CategoryID)
	}
	var count int64
	if err := chain.Count(&count).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var articles []models.CmsArticle
	if err := chain.Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var res []resp.ArticleResp
	response.Copy(&res, &articles)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (iSrv iArticleService) Detail(id uint) (resp.ArticleResp, error) {
	var article models.CmsArticle
	if err := iSrv.db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.ArticleResp{}, fmt.Errorf("记录不存在: %v", err)
		}
		return resp.ArticleResp{}, err
	}
	var res resp.ArticleResp
	response.Copy(&res, &article)
	return res, nil
}

func (iSrv iArticleService) Add(addReq req.ArticleAddReq, auth *req.AuthReq) error {
	var article models.CmsArticle
	response.Copy(&article, &addReq)
	if err := iSrv.db.Create(&article).Error; err != nil {
		return fmt.Errorf("创建失败: %v", err)
	}
	return nil
}

func (iSrv iArticleService) Edit(editReq req.ArticleEditReq, auth *req.AuthReq) error {
	var article models.CmsArticle
	if err := iSrv.db.First(&article, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %v", err)
		}
		return err
	}
	response.Copy(&article, &editReq)
	if err := iSrv.db.Save(&article).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

func (iSrv iArticleService) Del(id uint, auth *req.AuthReq) error {
	var article models.CmsArticle
	if err := iSrv.db.Delete(&article, id).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}
	return nil
}

func NewArticleService(db *gorm.DB) ArticleService {
	return iArticleService{
		db: db,
	}
}
