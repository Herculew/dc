package models 

type TgLine struct { 
    Id int `xorm:"primary_key autoincr comment('路线id') INT(11)" json:"id"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    Longitude float64 `xorm:"decimal(12,6)" json:"longitude"` 
    Latitude float64 `xorm:"decimal(12,6)" json:"latitude"` 
    Address string `xorm:"varchar(50)" json:"address"` 
    Title string `xorm:"varchar(20)" json:"title"` 
    Cover string `xorm:"varchar(100)" json:"cover"` 
    Tips string `xorm:"text" json:"tips"` 
    Distance float64 `xorm:"decimal(9,1)" json:"distance"` 
    Annex string `xorm:"text" json:"annex"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    RideSurplus string `xorm:"varchar(50)" json:"ride_surplus"` 
    DrawSpots string `xorm:"mediumtext" json:"draw_spots"` 
    CreateBy int `xorm:"int(10)" json:"create_by"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}