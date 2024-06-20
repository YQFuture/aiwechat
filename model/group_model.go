package model

type GroupModel struct {
	GroupName string `json:"groupName"`
	AvatarID  string `json:"avatarID"`
	FileData  []byte `json:"fileData"`
}

type GroupModelList []*GroupModel
