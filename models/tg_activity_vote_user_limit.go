package models 

type TgActivityVoteUserLimit struct { 
    Id int `xorm:"primary_key autoincr comment('主键 ID  A_T') INT(11)" json:"id"` 
    UserId int `xorm:"int(11)" json:"user_id"` 
    EventId int `xorm:"int(10)" json:"event_id"` 
    Sequence uint8 `xorm:"tinyint(1)" json:"sequence"` 
    VoteTotalLimit uint8 `xorm:"tinyint(1)" json:"vote_total_limit"` 
    CreateTime int `xorm:"int(11)" json:"create_time"` 
    UpdateTime int `xorm:"int(11)" json:"update_time"` 
}