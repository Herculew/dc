package models 

type TgLineNodeComment struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Content string `xorm:"text" json:"content"` 
    NodeId int `xorm:"int(10)" json:"node_id"` 
    BookId int `xorm:"int(10)" json:"book_id"` 
    Star uint8 `xorm:"tinyint(1)" json:"star"` 
    LikeCount int `xorm:"int(10)" json:"like_count"` 
    ReplyCount int `xorm:"int(10)" json:"reply_count"` 
    ForwardCount int `xorm:"int(10)" json:"forward_count"` 
    ViewCount int `xorm:"int(10)" json:"view_count"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    IsRecommend uint8 `xorm:"tinyint(1)" json:"is_recommend"` 
    Detail string `xorm:"mediumtext" json:"detail"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    VehicleType string `xorm:"varchar(20)" json:"vehicle_type"` 
}