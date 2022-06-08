package model

import "time"

// Steps 地铁站点
type Steps struct {
	Id          int       `gorm:"column:id;type:mediumint(8);PRIMARY KEY;NOT NULL;" json:"id"`       //
	Name        string    `gorm:"column:name;type:varchar(255);NOT NULL;" json:"name"`               // 地铁站名称
	IsPractical int8      `gorm:"column:is_practical;type:tinyint(1);NOT NULL;" json:"is_practical"` //
	LineId      int8      `gorm:"column:line_id;type:int(10) unsigned;NOT NULL;" json:"line_id"`     // 所属线路
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`               // 创建时间
}

func (*Steps) TableName() string {
	return "steps"
}
