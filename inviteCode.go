package icd

var (
	Alphabet_name = map[int32]rune{
		0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E', 5: 'F', 6: 'G', 7: 'H', 8: 'I',
		9: 'J', 10: 'K', 11: 'M', 12: 'N', 13: 'P', 14: 'Q', 15: 'R', 16: 'S',
		17: 'T', 18: 'U', 19: 'V', 20: 'W', 21: 'X', 22: 'Y', 23: '3', 24: '4',
		25: '5', 26: '6', 27: '7', 28: '8', 29: '9',
	}
	Alphabet_value = map[rune]int32{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7, 'I': 8,
		'J': 9, 'K': 10, 'M': 11, 'N': 12, 'P': 13, 'Q': 14, 'R': 15, 'S': 16,
		'T': 17, 'U': 18, 'V': 19, 'W': 20, 'X': 21, 'Y': 22, '3': 23, '4': 24,
		'5': 25, '6': 26, '7': 27, '8': 28, '9': 29,
	}
	baseNumber int32 = 100000000
)

func MagicUID(uid int32) int32 {
	return (uid % 10) * baseNumber
}

func Create(uid int32) string {
	var mod int32 = 0
	uid += MagicUID(uid)
	result := []rune{}
	for uid > 0 {
		mod = uid % 30
		uid = uid / 30
		result = append(result, Alphabet_name[mod])
	}
	return string(result)
}
func Decode(inviteCode string) int32 {
	var uid int32 = 0
	r := []rune(inviteCode)
	for l := len(r) - 1; l >= 0; l-- {
		uid = uid*30 + Alphabet_value[r[l]]
	}
	return uid - MagicUID(uid)
}
