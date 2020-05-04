package models 

type Tg2handcarPark struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    EventId int `xorm:"int(10)" json:"event_id"` 
    HallId int `xorm:"int(10)" json:"hall_id"` 
    BoothId int `xorm:"int(10)" json:"booth_id"` 
    EventBoothId int `xorm:"int(10)" json:"event_booth_id"` 
    Name string `xorm:"varchar(100)" json:"name"` 
    Cover string `xorm:"varchar(100)" json:"cover"` 
    ParkDetail string `xorm:"text" json:"park_detail"` 
    ParkSort int `xorm:"int(8)" json:"park_sort"` 
    ParkTypeId int `xorm:"int(10)" json:"park_type_id"` 
    CreateUserId int `xorm:"int(10)" json:"create_user_id"` 
    Addtime string `xorm:"datetime" json:"addtime"` 
    Public uint8 `xorm:"tinyint(4)" json:"public"` 
    Code string `xorm:"char(6)" json:"code"` 
    Delete uint8 `xorm:"tinyint(4)" json:"delete"` 
    IsReserve uint8 `xorm:"tinyint(4)" json:"is_reserve"` 
}