package discount

import (
	"slices"
)

var db = map[int64][]int64{
	2100: {1, 2},
	2200: {3, 1, 2},
	2300: {2},
	2400: {192, 314, 436, 158},
	2500: {204, 326, 148, 370, 592},
	2600: {216},
	2700: {228, 350, 472, 194},
	// TODO: придумать что делать с этим крайним случаем
	2800: {},
	2900: {240, 362, 484, 206, 428},
	3000: {252, 374},
	3100: {264, 386, 508, 230},
	3200: {276, 398},
	3300: {288, 410, 532, 254},
	3400: {300, 422, 544, 166},
	3500: {312, 434},
	3600: {324, 446, 568, 190},
	3700: {336, 458},
	3800: {348, 470, 592, 214},
	3900: {360, 482, 604, 226},
	4000: {372, 494, 616, 238},
	4100: {384, 506, 628, 250},
	4200: {396, 518, 640, 262},
}

var UserIDs = GetUserIDsList(db)

func GetUserIDsList(m map[int64][]int64) []int64 {
	var result []int64
	for key := range m {
		result = append(result, key)
	}
	slices.Sort(result)
	return result
}

func GetSegmentsByID(userID int64) ([]int64, error) {
	left, right := 0, len(UserIDs)-1

	for left != right-1 {
		mid := left + (right-left)/2
		if UserIDs[mid] > userID {
			right = mid
		} else if UserIDs[mid] < userID {
			left = mid
		} else {
			return db[UserIDs[int64(mid)]], nil
		}
	}
	switch {
	case userID < UserIDs[right] || userID == UserIDs[left]:
		return db[UserIDs[int64(left)]], nil
	default:
		return db[UserIDs[int64(right)]], nil
	}
}

func GetSegmentsByUserIDs(userIDs []int64) map[int64][]int64 {
	result := make(map[int64][]int64, len(userIDs))

	for _, userID := range userIDs {
		result[userID] = db[userID]
	}
	return result
}
