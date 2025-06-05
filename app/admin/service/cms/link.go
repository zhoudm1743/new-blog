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

type LinkService interface {
	All(listReq req.LinkQueryReq, auth *req.AuthReq) ([]resp.LinkResp, error)
	List(page req.PageReq, listReq req.LinkQueryReq, auth *req.AuthReq) (response.PageResp, error)
	Detail(id uint) (resp.LinkResp, error)
	Add(addReq req.LinkAddReq, auth *req.AuthReq) error
	Edit(editReq req.LinkEditReq, auth *req.AuthReq) error
	Del(id uint, auth *req.AuthReq) error
}

type iLinkService struct {
	db *gorm.DB
}

func (iSrv iLinkService) All(listReq req.LinkQueryReq, auth *req.AuthReq) ([]resp.LinkResp, error) {
	chain := iSrv.db.Model(&models.CmsLink{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var links []models.CmsLink
	if err := chain.Find(&links).Error; err != nil {
		return nil, err
	}
	var res []resp.LinkResp
	response.Copy(&res, &links)
	return res, nil
}

func (iSrv iLinkService) List(page req.PageReq, listReq req.LinkQueryReq, auth *req.AuthReq) (response.PageResp, error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	chain := iSrv.db.Model(&models.CmsLink{}).Order("id desc")
	if listReq.Name != "" {
		chain = chain.Where("name LIKE ?", "%"+listReq.Name+"%")
	}
	if listReq.Status != -1 {
		chain = chain.Where("status = ?", listReq.Status)
	}
	var count int64
	if err := chain.Count(&count).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var links []models.CmsLink
	if err := chain.Limit(limit).Offset(offset).Find(&links).Error; err != nil {
		return response.PageResp{}, fmt.Errorf("查询失败: %v", err)
	}
	var res []resp.LinkResp
	response.Copy(&res, &links)
	return response.PageResp{
		Count:    count,
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Lists:    res,
	}, nil
}

func (iSrv iLinkService) Detail(id uint) (resp.LinkResp, error) {
	var link models.CmsLink
	if err := iSrv.db.First(&link, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp.LinkResp{}, fmt.Errorf("记录不存在: %v", err)
		}
		return resp.LinkResp{}, err
	}
	var res resp.LinkResp
	response.Copy(&res, &link)
	return res, nil
}

func (iSrv iLinkService) Add(addReq req.LinkAddReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.CmsLink{}).Where("name = ?", addReq.Name).Count(&count)
	if count > 0 {
		return fmt.Errorf("记录已存在")
	}
	var link models.CmsLink
	response.Copy(&link, &addReq)
	if err := iSrv.db.Create(&link).Error; err != nil {
		return fmt.Errorf("创建失败: %v", err)
	}
	return nil
}

func (iSrv iLinkService) Edit(editReq req.LinkEditReq, auth *req.AuthReq) error {
	var count int64
	iSrv.db.Model(&models.CmsLink{}).Where("name = ?", editReq.Name).Where("id != ?", editReq.ID).Count(&count)
	if count > 0 {
		return fmt.Errorf("记录已存在")
	}
	var link models.CmsLink
	if err := iSrv.db.First(&link, editReq.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("记录不存在: %v", err)
		}
		return err
	}
	response.Copy(&link, &editReq)
	if err := iSrv.db.Save(&link).Error; err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}
	return nil
}

func (iSrv iLinkService) Del(id uint, auth *req.AuthReq) error {
	var link models.CmsLink
	if err := iSrv.db.Delete(&link, id).Error; err != nil {
		return fmt.Errorf("删除失败: %v", err)
	}
	return nil
}

func NewLinkService(db *gorm.DB) LinkService {
	return iLinkService{
		db: db,
	}
}
