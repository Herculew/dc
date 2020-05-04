package models 

type TgLineNodeAlbum struct { 
    Id int `xorm:"primary_key autoincr comment('数据表主键') INT(11)" json:"id"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    CommentId int `xorm:"int(10)" json:"comment_id"` 
    Photo string `xorm:"varchar(250)" json:"photo"` 
    Platform string `xorm:"varchar(10)" json:"platform"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    BookId int `xorm:"int(10)" json:"book_id"` 
    NodeId int `xorm:"int(10)" json:"node_id"` 
    Width string `xorm:"mediumint(6)" json:"width"` 
    Height string `xorm:"mediumint(6)" json:"height"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
}