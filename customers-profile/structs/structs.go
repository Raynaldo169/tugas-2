package structs

type User struct {
	UserID int    `json:"id" gorm:"primary_key;auto_increment"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

type Risk_Profile struct {
	UserID       int     `json:"id" gorm:"index"`
	MMPercent    float32 `json:"mm"`
	BondPercent  float32 `json:"bond"`
	StockPercent float32 `json:"stock"`
	User         *User   `json:"user" gorm:"foreignKey:UserID"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
