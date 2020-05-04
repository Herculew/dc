package models 

type TgAdminUserAuthCopy struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Uid int `xorm:"int(10)" json:"uid"` 
    Openid string `xorm:"varchar(50)" json:"openid"` 
    Cid int `xorm:"int(10)" json:"cid"` 
    Unionid int `xorm:"int(10)" json:"unionid"` 
    Nickname string `xorm:"varchar(255)" json:"nickname"` 
    Headimgurl string `xorm:"varchar(1024)" json:"headimgurl"` 
    Sex uint8 `xorm:"tinyint(1)" json:"sex"` 
    Createtime string `xorm:"datetime" json:"createtime"` 
}