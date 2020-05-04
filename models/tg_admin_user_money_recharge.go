package models 

type TgAdminUserMoneyRecharge struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Uid int `xorm:"int(10)" json:"uid"` 
    Cid int `xorm:"int(10)" json:"cid"` 
    Way string `xorm:"varchar(200)" json:"way"` 
    Money float64 `xorm:"decimal(10,2)" json:"money"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    RechargeFrom uint8 `xorm:"tinyint(1)" json:"recharge_from"` 
    ToAccountNum string `xorm:"varchar(255)" json:"to_account_num"` 
    OrderSn string `xorm:"varchar(255)" json:"order_sn"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
}