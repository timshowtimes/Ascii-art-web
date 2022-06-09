package handlers

import (
	"crypto/sha256"
	"strings"
)

var hash_std = []byte{
	225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97,
	245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191,
}

var hash_thinkertoy = []byte{
	165, 123, 238, 196, 63, 222, 103, 81, 186, 29, 48, 73, 91,
	9, 38, 88, 160, 100, 69, 47, 50, 30, 34, 29, 8, 195, 172, 52, 169, 220, 18, 148,
}

var hash_shadow = []byte{
	38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129,
	116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115,
}

func Contains(s []string, str []string) bool {
	for _, v := range s {
		if v == str[0] {
			return true
		}
	}

	return false
}

func IsNotAscii(s string) bool {
	res := strings.ReplaceAll(s, "\n", "\\n")
	for _, val := range res {
		if val < 13 || val > 126 {
			return true
		}
	}
	return false
}

func GetStr(s string, x map[rune]string) string {
	if s == "" {
		return "\n"
	} else {
		res := ""
		temp := make([]string, 11)
		for _, val := range s {
			for n, r := range strings.Split(x[val], "\n") {
				temp[n] += r
			}
		}
		for _, val := range temp {
			res += val + "\n"
		}
		return res[1 : len(res)-2]
	}
}

func GetMap(s string) map[rune]string {
	symbol := make(map[rune]string)
	str := ""
	j := rune(32)
	count := 0
	for _, v := range s {
		str += string(v)
		if string(v) == "\n" {
			count++
		}
		if count == 9 {
			symbol[j] = str
			str = ""
			j++
			count = 0
		}
	}
	return symbol
}

func DHashSum(Fp []byte) bool {
	h := sha256.New()
	h.Write(Fp)

	if string(h.Sum(nil)) != string(hash_std) && string(h.Sum(nil)) != string(hash_thinkertoy) && string(h.Sum(nil)) != string(hash_shadow) {
		return false
	}
	return true
}
