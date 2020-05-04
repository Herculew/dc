package models 

type TgLiveUser struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    VId int `xorm:"int(10)" json:"v_id"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    UserName string `xorm:"varchar(225)" json:"user_name"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    IpAddress string `xorm:"varchar(255)" json:"ip_address"` 
    LeaveTime string `xorm:"datetime" json:"leave_time"` 
}