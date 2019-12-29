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

type HotDoc struct {
	Id int64
	Hash string
	Title string
	Img string
	Counter int32
}

func (doc *Doc) Update() {
	db.Save(doc)
}

func DocSaveOrRepl(doc *Doc) bool {
	var result = new(Doc)
	db.Where("hash=?", doc.Hash).First(result)
	now := time.Now().Unix()
	if result.Id>0 {

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

func GetHotRanging() []*HotDoc {
	results := make([]*HotDoc, 0)
	sql := "select d.id,d.hash,d.title,d.img,c.counter from docs as d left join clicks as c on c.hash=d.hash where d.status=? order by c.counter desc,d.id desc limit 10"
	db.Raw(sql, DOC_STATUS_NORMAL).Scan(&results)

	if len(results)>0 {
		for _,v := range results {
			if v.Img=="" {
				v.Img = defaultImgs[v.Id%3]
			}
		}
	}

	return results
}

func GetDoc(hash string) (*Doc,bool) {
	doc := new(Doc)
	db.Where("hash=?", hash).First(doc)

	ok := false
	if doc.Id>0 {
		ok = true

		if doc.Img=="" {
			doc.Img = defaultImgs[doc.Id%3]
		}
	}


	return doc,ok
}

func SearchDocs(keywords string, page int32, limit int32) ([]*Doc, int32) {
	results := make([]*Doc, 0)

	var counter int32 = 0
	db.Where("title like ? and status=?", "%"+keywords+"%", DOC_STATUS_NORMAL).Table("docs").Count(&counter)

	if page<=0 {
		page=1
	}
	start := (page-1)*limit
	db.Where("title like ? and status=?", "%"+keywords+"%", DOC_STATUS_NORMAL).Order("id desc").Offset(start).Limit(limit).Find(&results)

	if len(results)>0 {
		for _,v := range results {
			if v.Img=="" {
				v.Img = defaultImgs[v.Id%3]
			}
		}
	}

	return results, counter
}

func GetDocsBySubject(subject string ,page int32, limit int32) ([]*Doc, int32) {
	results := make([]*Doc, 0)

	var counter int32 = 0
	db.Where("status=? and parent_hash=?", DOC_STATUS_NORMAL, subject).Table("docs").Count(&counter)

	if page<=0 {
		page=1
	}
	start := (page-1)*limit
	db.Where("status=? and parent_hash=?", DOC_STATUS_NORMAL, subject).Order("id desc").Offset(start).Limit(limit).Find(&results)

	if len(results)>0 {
		for _,v := range results {
			if v.Img=="" {
				v.Img = defaultImgs[v.Id%3]
			}
		}
	}

	return results, counter
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