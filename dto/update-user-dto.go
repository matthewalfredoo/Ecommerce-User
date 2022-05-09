package dto

// UpdateUserDTO is a data transfer object for updating data user
type UpdateUserDTO struct {
	Id       uint64 `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Alamat   string `json:"alamat" form:"alamat" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	NomorHP  string `json:"nomor_hp" form:"nomor_hp" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
