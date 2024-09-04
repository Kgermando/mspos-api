package models

import "gorm.io/gorm"

type Asm struct {
	gorm.Model

	Name       string `gorm:"not null" json:"name"`
	ProvinceID uint   `gorm:"foreignKey:province_id" json:"province_id"`
	Signature  string `json:"signature"`
	Sups       []Sup
}

func (p *Asm) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Asm{}).Count(&total)
	return total
}

func (p *Asm) Paginate(db *gorm.DB, limit int, offset int) interface{} {
	sp := []Asm{}
	db.Offset(offset).Limit(limit).Find(&sp)
	return sp
}