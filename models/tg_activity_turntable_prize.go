package models 

type TgActivityTurntablePrize struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    Name string `xorm:"varchar(100)" json:"name"` 
    PrizeType uint8 `xorm:"tinyint(1)" json:"prize_type"` 
    PrizeId int `xorm:"int(10)" json:"prize_id"` 
    Nums int `xorm:"int(10)" json:"nums"` 
    Scale string `xorm:"varchar(10)" json:"scale"` 
    Remarks string `xorm:"varchar(200)" json:"remarks"` 
    Img string `xorm:"varchar(200)" json:"img"` 
    TurntableId int `xorm:"int(10)" json:"turntable_id"` 
    NumsSend float64 `xorm:"decimal(10,2)" json:"nums_send"` 
    EventId int `xorm:"int(10)" json:"event_id"` 
    HtId int `xorm:"int(10)" json:"ht_id"` 
    Cash float64 `xorm:"decimal(8,2)" json:"cash"` 
}