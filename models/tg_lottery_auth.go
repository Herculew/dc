package models 

type TgLotteryAuth struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Uid int `xorm:"int(10)" json:"uid"` 
    IsDel uint8 `xorm:"tinyint(1)" json:"is_del"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    UpdateBy int `xorm:"int(10)" json:"update_by"` 
}