package shamazing

import (
	"regexp"
	"strconv"
)

// FindLongestString will retrieve a string like a SHA1, MD5 or whatever
// and return the longest string (first one)
func FindLongestString(str string) string {
	var re = regexp.MustCompile("[a-zA-Z]+")
	var values = re.FindAll([]byte(str), -1)

	return string(findLongest(values))
}

// FindLongestInteger will retrieve a string like a SHA1, MD5 or whatever
// and return longest integer (first one)
func FindLongestInteger(str string) (int64, error) {
	var re = regexp.MustCompile("[0-9]+")
	var values = re.FindAll([]byte(str), -1)
	var max = findLongest(values)

	var a, err = strconv.Atoi(string(max))
	if err != nil {
		return 0, err
	}
	return int64(a), nil
}

func findLongest(values [][]byte) []byte {
	var max []byte
	for _, value := range values {
		if len(value) > len(max) {
			max = value
		}
	}
	return max
}
