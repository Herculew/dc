package models 

type TgLineNodePunchCensus struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    NodeId int `xorm:"int(10)" json:"node_id"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    PunchCount string `xorm:"mediumint(6)" json:"punch_count"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}