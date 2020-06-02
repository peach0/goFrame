package model

type BaseModel struct {
	Id          uint  `gorm:"primary_key"`
	CreatedTime int64 `gorm:"type:int;"`
	UpdatedTime int64 `gorm:"type:int;"`
	DeletedTime int64 `gorm:"type:int;"`
}
