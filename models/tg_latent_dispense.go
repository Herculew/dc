package models 

type TgLatentDispense struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    DispenseOrgId int `xorm:"int(10)" json:"dispense_org_id"` 
    OrgId int `xorm:"int(10)" json:"org_id"` 
    ResourceLatentId int `xorm:"int(10)" json:"resource_latent_id"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    LatentType string `xorm:"varchar(255)" json:"latent_type"` 
}