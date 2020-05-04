package models 

type TgLotteryCron struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    ActType uint8 `xorm:"tinyint(1)" json:"act_type"` 
    ActId int `xorm:"int(10)" json:"act_id"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}