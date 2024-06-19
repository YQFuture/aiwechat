package model

import (
	"strings"
	"unicode"
)

type GroupModel struct {
	GroupName string `json:"groupName"`
	AvatarID  string `json:"avatarID"`
}

type GroupModelList []*GroupModel

func GroupGroupByInitial(groups []*GroupModel) map[rune][]*GroupModel {
	grouped := make(map[rune][]*GroupModel)
	var specialGroup []*GroupModel

	for _, group := range groups {
		firstRune := rune(strings.ToLower(group.GroupName)[0])
		var groupKey rune
		var found bool

		if unicode.IsLetter(firstRune) {
			groupKey = firstRune
		} else if unicode.Is(unicode.Han, firstRune) {
			groupKey, found = getPinyinFirstLetter(group.GroupName)
			if !found {
				groupKey = '#'
			}
		} else {
			groupKey = '#'
		}

		if groupKey == '#' {
			specialGroup = append(specialGroup, group)
		} else {
			grouped[groupKey] = append(grouped[groupKey], group)
		}
	}

	for char := 'a'; char <= 'z'; char++ {
		if _, exists := grouped[char]; !exists {
			grouped[char] = []*GroupModel{}
		}
	}

	grouped['#'] = specialGroup

	return grouped
}
