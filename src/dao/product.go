package dao

import "time"

type Project struct {
	Id          string    `xorm:"pk" json:"id"`
	Name        string    `json:"name"`
	Describe    string    `xorm:"text 'describe'" json:"describe"`
	CreatedTime time.Time `xorm:"created_time created" json:"created_time"` //创建时间
	Registry    string    `xorm:"'registry'" json:"registry"`
}
