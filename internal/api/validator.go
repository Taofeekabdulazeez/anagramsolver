package api

import "regexp"

func IsValidWordParam(param string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z]+$`, param)
	return match
}
