package models 

type TgLineNode struct { 
    Id int `xorm:"primary_key autoincr comment('') INT(11)" json:"id"` 
    Type uint8 `xorm:"tinyint(1)" json:"type"` 
    OrgId int `xorm:"int(10)" json:"org_id"` 
    MainWork uint8 `xorm:"tinyint(2)" json:"main_work"` 
    Name string `xorm:"varchar(55)" json:"name"` 
    Pinyin string `xorm:"varchar(55)" json:"pinyin"` 
    Letter string `xorm:"char(1)" json:"letter"` 
    ProvinceId int `xorm:"int(10)" json:"province_id"` 
    CityId int `xorm:"int(10)" json:"city_id"` 
    DistrictId int `xorm:"int(10)" json:"district_id"` 
    City string `xorm:"varchar(25)" json:"city"` 
    Address string `xorm:"varchar(150)" json:"address"` 
    Longitude float64 `xorm:"decimal(12,6)" json:"longitude"` 
    Latitude float64 `xorm:"decimal(12,6)" json:"latitude"` 
    Distance float64 `xorm:"decimal(9,1)" json:"distance"` 
    Detail string `xorm:"text" json:"detail"` 
    Score float64 `xorm:"decimal(2,1)" json:"score"` 
    TotalScore int `xorm:"int(10)" json:"total_score"` 
    PunchCount string `xorm:"mediumint(8)" json:"punch_count"` 
    PraiseCount string `xorm:"smallint(6)" json:"praise_count"` 
    NegativeCount string `xorm:"smallint(6)" json:"negative_count"` 
    OrdinaryCount string `xorm:"smallint(6)" json:"ordinary_count"` 
    CommentCount int `xorm:"int(10)" json:"comment_count"` 
    RoadbookCount int `xorm:"int(10)" json:"roadbook_count"` 
    Status uint8 `xorm:"tinyint(1)" json:"status"` 
    Prohibit uint8 `xorm:"tinyint(1)" json:"prohibit"` 
    IsDelete uint8 `xorm:"tinyint(1)" json:"is_delete"` 
    Annex string `xorm:"text" json:"annex"` 
    Desc string `xorm:"text" json:"desc"` 
    CreateTime string `xorm:"datetime" json:"create_time"` 
    UpdateTime string `xorm:"datetime" json:"update_time"` 
}