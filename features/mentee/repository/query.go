package repository

import (
	"be12/mentutor/features/mentee"
	"log"

	"gorm.io/gorm"
)

type menteeData struct {
	db *gorm.DB
}

func New(db *gorm.DB) mentee.RepoInterface {
	return &menteeData{
		db: db,
	}
}

func (md *menteeData) EditProfile(id uint, data mentee.MenteeCore) (mentee.MenteeCore, error) {
	var res Mentee
	
	res = FromEntity(data)
	if err := md.db.Where("id = ?", id).Updates(&res).Error; err != nil {
		log.Print(err.Error(), "ERROR INSERT TO DATBASE")
		return mentee.MenteeCore{}, err
	}
	
	cnv := ToEntity(id, res)
	
	return cnv, nil
}