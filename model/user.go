package model

type User struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Nama     string `gorm:"type:varchar(255)" json:"nama"`
	Alamat   string `gorm:"type:varchar(255)" json:"alamat"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	NomorHP  string `gorm:"type:varchar(255)" json:"nomor_hp"`
	Password string `gorm:"->;<-;not null" json:"-"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
