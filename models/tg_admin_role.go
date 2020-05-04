package models 

type TgAdminRole struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Name string `xorm:"varchar(64)" json:"name"` 
    Rule string `xorm:"text" json:"rule"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UpdateBy int `xorm:"int(10)" json:"update_by"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
}