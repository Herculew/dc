package models 

type TgLiveMessage struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Jpushtext string `xorm:"varchar(255)" json:"jpushtext"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    IsRead uint8 `xorm:"tinyint(1)" json:"is_read"` 
    IsPush uint8 `xorm:"tinyint(1)" json:"is_push"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    PushTime string `xorm:"datetime" json:"push_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    SourceId int `xorm:"int(11)" json:"source_id"` 
}