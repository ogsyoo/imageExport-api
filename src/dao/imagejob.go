package dao

type ImageJob struct {
	Id          int32  `xorm:"pk autoincr"`
	Describe    string `json:"describe"`
	From        string `json:"from"`
	To          string `json:"to"`
	IsPackage   int8   `json:"is_package"`
	UpdateTime  int64  `xorm:"update_time"`  //修改后自动更新时间
	CreatedTime int64  `xorm:"created_time"` //创建时间
	Status      int8   `json:"status"`
}
