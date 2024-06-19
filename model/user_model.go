package model

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
	"unicode"
)

type UserModel struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	AvatarID string `json:"avatarID"`
	FileData []byte `json:"fileData"`
}

type UserModelList []*UserModel

func (a UserModelList) Len() int {
	return len(a)
}

func (a UserModelList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a UserModelList) Less(i, j int) bool {
	iStartsWithSymbol, jStartsWithSymbol := !isLetterOrChinese([]rune(a[i].NickName)[0]), !isLetterOrChinese([]rune(a[j].NickName)[0])

	if iStartsWithSymbol && jStartsWithSymbol {
		return a[i].NickName < a[j].NickName
	} else if iStartsWithSymbol {
		return false
	} else if jStartsWithSymbol {
		return true
	}

	pinyinI := getFirstPinyin(a[i].NickName)
	pinyinJ := getFirstPinyin(a[j].NickName)

	return strings.ToLower(pinyinI) < strings.ToLower(pinyinJ)
}

func isLetterOrChinese(r rune) bool {
	return unicode.IsLetter(r) || unicode.Is(unicode.Han, r)
}

func getFirstPinyin(name string) string {
	pinyinList := pinyin.LazyPinyin(name, pinyin.NewArgs())
	if len(pinyinList) > 0 {
		return pinyinList[0]
	}
	return string([]rune(name)[0])
}

func getPinyinFirstLetter(nickname string) (rune, bool) {
	pinyinList := pinyin.LazyPinyin(nickname, pinyin.NewArgs())
	if len(pinyinList) > 0 {
		firstLetter := rune(strings.ToLower(pinyinList[0])[0])
		if unicode.IsLetter(firstLetter) {
			return firstLetter, true
		}
	}
	return 0, false
}

func UserGroupByInitial(users []*UserModel) map[rune][]*UserModel {
	grouped := make(map[rune][]*UserModel)
	var specialGroup []*UserModel

	for _, user := range users {
		firstRune := rune(strings.ToLower(user.NickName)[0])
		var groupKey rune
		var found bool

		if unicode.IsLetter(firstRune) {
			groupKey = firstRune
		} else if unicode.Is(unicode.Han, firstRune) {
			groupKey, found = getPinyinFirstLetter(user.NickName)
			if !found {
				groupKey = '#'
			}
		} else {
			groupKey = '#'
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
