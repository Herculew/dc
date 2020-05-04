package models 

type TgLineNodeCommentText struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"`
    Title string `xorm:"varchar(20)" json:"title"` 
    Pid uint8 `xorm:"tinyint(1)" json:"pid"` 
    MainWork string `xorm:"varchar(60)" json:"main_work"` 
    Text string `xorm:"varchar(250)" json:"text"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}