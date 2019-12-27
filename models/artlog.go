package model

import "time"

type Artlog struct {
	Id int64 `gorm:"primary_key"`
	Ip string
	Articleid int64
	Sessid string
	CreatedAt int64
	Useragent string
}

func SaveArtlog(artlog *Artlog) bool {
	artlog.CreatedAt = time.Now().Unix()
	db.Save(artlog)

	return true
}

func GetArtlog(sessid string, docId int64) (*Artlog,bool) {
	artlog := new(Artlog)
	db.Where("sessid=? and articleid=?", sessid, docId).Order("created_at desc").First(artlog)

	ok := false
	if artlog.Id>0 {
		ok = true
	}


	return artlog,ok
}