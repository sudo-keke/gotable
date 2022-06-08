package model

import "time"

// City 中国已开通地铁城市
type City struct {
	Id        int8      `gorm:"column:id;type:int(10);PRIMARY KEY;NOT NULL;" json:"id"`    // id
	CnName    string    `gorm:"column:cn_name;type:varchar(255);NOT NULL;" json:"cn_name"` // 中文名
	EnName    string    `gorm:"column:en_name;type:varchar(255);NOT NULL;" json:"en_name"` // 英文名
	Code      int8      `gorm:"column:code;type:int(11);NOT NULL;" json:"code"`            // code
	Pre       string    `gorm:"column:pre;type:varchar(255);NOT NULL;" json:"pre"`         // 简称
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`       // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`       // 更新时间
}

func (*City) TableName() string {
	return "citys"
}
