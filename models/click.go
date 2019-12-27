package model

import "time"

type Click struct {
	Id int64
	Hash string
	Counter int64
	CreatedAt int64
	UpdatedAt int64
}

func GetClickCount(hash string) int64 {
	var click = new(Click)
	db.Where("hash=?", hash).First(click)
	if click.Id>0 {
		return click.Counter
	}
	return 0
}

func ClickIncr(hash string) int64 {
	var click = new(Click)
	db.Where("hash=?", hash).First(click)
	now := time.Now().Unix()
	if click.Id>0 {
		click.Counter = click.Counter+1
		click.UpdatedAt = now
	}else{
		click.Hash = hash
		click.Counter = 1
		click.CreatedAt = now
		click.UpdatedAt = 0
	}
	db.Save(click)

	return click.Counter
}
