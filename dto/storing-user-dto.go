package dto

// StoringUserDTO is a data transfer object for registering and updating data user
type StoringUserDTO struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Alamat   string `json:"alamat" form:"alamat" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	NomorHP  string `json:"nomor_hp" form:"nomor_hp" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
