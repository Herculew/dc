package models 

type TgLotteryConfig struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    LotteryId int `xorm:"int(10)" json:"lottery_id"` 
    ConfType uint8 `xorm:"tinyint(1)" json:"conf_type"` 
    ConfTypeSub uint8 `xorm:"tinyint(1)" json:"conf_type_sub"` 
    ConfValue string `xorm:"varchar(30)" json:"conf_value"` 
    LimitType uint8 `xorm:"tinyint(1)" json:"limit_type"` 
    Weight int `xorm:"int(10)" json:"weight"` 
    IsDel uint8 `xorm:"tinyint(1)" json:"is_del"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    UpdateBy int `xorm:"int(10)" json:"update_by"` 
}