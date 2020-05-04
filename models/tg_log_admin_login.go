package models 

type TgLogAdminLogin struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    Ip string `xorm:"varchar(255)" json:"ip"` 
    Remark string `xorm:"varchar(255)" json:"remark"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
}