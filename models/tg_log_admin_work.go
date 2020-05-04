package models 

type TgLogAdminWork struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    UserId int `xorm:"int(10)" json:"user_id"` 
    Module string `xorm:"varchar(255)" json:"module"` 
    Controller string `xorm:"varchar(255)" json:"controller"` 
    Action string `xorm:"varchar(255)" json:"action"` 
    Condition string `xorm:"text" json:"condition"` 
    Header string `xorm:"text" json:"header"` 
    Remark string `xorm:"varchar(255)" json:"remark"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
}