package model

import (
	"time"
)

const (
	DOC_STATUS_NORMAL = 0
	DOC_STATUS_DISABLED = 1
)

var (
	defaultImgs = []string{"/resources/imgs/default1.png","/resources/imgs/default2.png","/resources/imgs/default.jpg"}
)

// 解析结果
// http://gorm.book.jasperxu.com/ gorm 文档
type Doc struct {
	Id int64 `gorm:"primary_key"`
	Hash string
	Path string
	Title string
	Desc string
	Intro string
	Status int32
	Parent string
	ParentHash string
	Img string
	CreatedAt int64
	UpdatedAt int64
	EditedAt int64
}

func DocSaveOrRepl(doc *Doc) bool {
	var result = new(Doc)
	db.Where("hash=?", doc.Hash).First(result)
	now := time.Now().Unix()
	if result.Id>0 {
		//if result.EditedAt<doc.EditedAt {
		//
		//
		//	log.Debug("", "edited %s", doc.Path)
		//}

		result.Path = doc.Path
		result.Title = doc.Title
		result.Desc = doc.Desc
		result.Intro = doc.Intro
		result.UpdatedAt = now
		result.EditedAt = doc.EditedAt
		db.Save(result)
	}else{
		result.CreatedAt = now
		result.Status = DOC_STATUS_NORMAL
		db.Save(doc)
	}

	return true
}

func GetDocsByPage(page int32, limit int32) ([]*Doc, int32) {
	results := make([]*Doc, 0)

	var counter int32 = 0
	db.Where("status=?", DOC_STATUS_NORMAL).Table("docs").Count(&counter)

	if page<=0 {
		page=1
	}
	start := (page-1)*limit
	db.Where("status=?", DOC_STATUS_NORMAL).Order("id desc").Offset(start).Limit(limit).Find(&results)

	if len(results)>0 {
		for _,v := range results {
			if v.Img=="" {
				v.Img = defaultImgs[v.Id%3]
			}
		}
	}

	return results, counter
}