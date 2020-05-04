package models 

type TgLineNodeRelated struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    NodeId int `xorm:"int(10)" json:"node_id"` 
    LineId int `xorm:"int(10)" json:"line_id"` 
    Date string `xorm:"date" json:"date"` 
    TimeNode uint8 `xorm:"tinyint(2)" json:"time_node"` 
    Sort uint8 `xorm:"tinyint(2)" json:"sort"` 
    Week string `xorm:"char(10)" json:"week"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    PunchCount int `xorm:"int(10)" json:"punch_count"` 
    CommentCount int `xorm:"int(10)" json:"comment_count"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
}