package models 

type TgAdminUserAvatar struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    Path string `xorm:"varchar(100)" json:"path"` 
    Sort uint8 `xorm:"tinyint(1)" json:"sort"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
}