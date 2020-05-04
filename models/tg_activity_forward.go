package models 

type TgActivityForward struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Name string `xorm:"varchar(100)" json:"name"` 
    Cover string `xorm:"varchar(200)" json:"cover"` 
    ForwardType uint8 `xorm:"tinyint(1)" json:"forward_type"` 
    ForwardId int `xorm:"int(10)" json:"forward_id"` 
    StartTime string `xorm:"datetime" json:"start_time"` 
    EndTime string `xorm:"datetime" json:"end_time"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    Scale string `xorm:"varchar(30)" json:"scale"` 
    CheckStatus uint8 `xorm:"tinyint(1)" json:"check_status"` 
    CheckFailReason string `xorm:"varchar(200)" json:"check_fail_reason"` 
    IsFb uint8 `xorm:"tinyint(1)" json:"is_fb"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    UpdateBy int `xorm:"int(10)" json:"update_by"` 
    CreateOid int `xorm:"int(10)" json:"create_oid"` 
    Limits int `xorm:"int(10)" json:"limits"` 
    LimitsType uint8 `xorm:"tinyint(1)" json:"limits_type"` 
    WheelNums int `xorm:"int(5)" json:"wheel_nums"` 
    IsOn uint8 `xorm:"tinyint(1)" json:"is_on"` 
    IsDel uint8 `xorm:"tinyint(1)" json:"is_del"` 
    CreateOname string `xorm:"varchar(100)" json:"create_oname"` 
    CreateUname string `xorm:"varchar(50)" json:"create_uname"` 
    VehicleType string `xorm:"varchar(20)" json:"vehicle_type"` 
    EventId int `xorm:"int(10)" json:"event_id"` 
    ForwardHtId int `xorm:"int(10)" json:"forward_ht_id"` 
}