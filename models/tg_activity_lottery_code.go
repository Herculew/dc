package models 

type TgActivityLotteryCode struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    EventId int `xorm:"int(10)" json:"event_id"` 
    TicketId string `xorm:"varchar(100)" json:"ticket_id"` 
    OrderSn string `xorm:"varchar(30)" json:"order_sn"` 
    Code string `xorm:"char(8)" json:"code"` 
    CodeSn string `xorm:"char(12)" json:"code_sn"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    NickName string `xorm:"varchar(128)" json:"nick_name"` 
    NickNameEmoji string `xorm:"varchar(100)" json:"nick_name_emoji"` 
    Avatar string `xorm:"varchar(255)" json:"avatar"` 
    Tel string `xorm:"varchar(64)" json:"tel"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
}