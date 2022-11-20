package entity

// omitemptyについて [https://ema-hiro.hatenablog.com/entry/2018/01/17/024321]
// indexについて [https://gorm.io/ja_JP/docs/indexes.html]
// タグの記述について [https://gorm.io/ja_JP/docs/models.html]

type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"-"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    *[]Book `json:"books,omitempty"`
}
