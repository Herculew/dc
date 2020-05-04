package models 

type TgLineNodeCommentMinutia struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    NodeId string `xorm:"smallint(6)" json:"node_id"` 
    NodeCommentId string `xorm:"text" json:"node_comment_id"` 
    TextId int `xorm:"int(10)" json:"text_id"` 
    Count int `xorm:"int(10)" json:"count"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}