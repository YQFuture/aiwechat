package model

type UserModel struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	AvatarID string `json:"avatarID"`
	FileData []byte `json:"fileData"`
}
