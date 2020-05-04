package models 

type TgLottery struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Name string `xorm:"varchar(30)" json:"name"` 
    Cover string `xorm:"varchar(200)" json:"cover"` 
    Intro string `xorm:"varchar(200)" json:"intro"` 
    Content string `xorm:"text" json:"content"` 
    CreateType uint8 `xorm:"tinyint(1)" json:"create_type"` 
    LotteryType uint8 `xorm:"tinyint(1)" json:"lottery_type"` 
    ActTypes string `xorm:"varchar(30)" json:"act_types"` 
    SuperLuckyNums int `xorm:"int(10)" json:"super_lucky_nums"` 
    IsFollowMe uint8 `xorm:"tinyint(1)" json:"is_follow_me"` 
    FollowUids string `xorm:"varchar(100)" json:"follow_uids"` 
    MoneyCode int `xorm:"int(10)" json:"money_code"` 
    CodeMoney int `xorm:"int(10)" json:"code_money"` 
    LuckyType uint8 `xorm:"tinyint(1)" json:"lucky_type"` 
    Year int `xorm:"int(10)" json:"year"` 
    Month int `xorm:"int(10)" json:"month"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    IsOn uint8 `xorm:"tinyint(1)" json:"is_on"` 
    IsDel uint8 `xorm:"tinyint(1)" json:"is_del"` 
    IsOpen uint8 `xorm:"tinyint(1)" json:"is_open"` 
    CreateOrgId int `xorm:"int(10)" json:"create_org_id"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    UpdateBy int `xorm:"int(10)" json:"update_by"` 
    FbTime string `xorm:"datetime" json:"fb_time"` 
    FbBy int `xorm:"int(10)" json:"fb_by"` 
}