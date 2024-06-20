package model

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

type UserModel struct {
	UserName   string `json:"userName"`
	NickName   string `json:"nickName"`
	RemarkName string `json:"remarkName"`
	AvatarID   string `json:"avatarID"`
	FileData   []byte `json:"fileData"`
}

type UserModelList []*UserModel

func isEnglishLetter(r rune) bool {
	return r >= 'A' && r <= 'Z' || r >= 'a' && r <= 'z'
}

func getPinyinFirstLetter(name string) (rune, bool) {
	pinyinList := pinyin.LazyPinyin(name, pinyin.NewArgs())
	if len(pinyinList) > 0 {
		firstLetter := rune(strings.ToLower(pinyinList[0])[0])
		if isEnglishLetter(firstLetter) {
			return firstLetter, true
		}
	}
	return 0, false
}

func UserGroupByInitial(users []*UserModel) map[rune][]*UserModel {
	grouped := make(map[rune][]*UserModel)
	var specialGroup []*UserModel

	for _, user := range users {
		firstRune := rune(strings.ToLower(user.RemarkName)[0])
		var groupKey rune
		var found bool

		if isEnglishLetter(firstRune) {
			groupKey = firstRune
		} else {
			groupKey, found = getPinyinFirstLetter(user.RemarkName)
			if !found {
				groupKey = '#'
			}
		}

		if groupKey == '#' {
			specialGroup = append(specialGroup, user)
		} else {
			grouped[groupKey] = append(grouped[groupKey], user)
		}
	}

	for char := 'a'; char <= 'z'; char++ {
		if _, exists := grouped[char]; !exists {
			grouped[char] = []*UserModel{}
		}
	}

	grouped['#'] = specialGroup

	return grouped
}
