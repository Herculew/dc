package models 

type TgAdvertisementAdminLog struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    SourceId int `xorm:"int(11)" json:"source_id"` 
    ModuleName string `xorm:"varchar(255)" json:"module_name"` 
    CreateBy int `xorm:"int(11)" json:"create_by"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
    ActionName string `xorm:"varchar(255)" json:"action_name"` 
    UpdateBy int `xorm:"int(11)" json:"update_by"` 
}