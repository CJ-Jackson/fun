package commonUtil

import "strings"

func CleanAndSplitString(str string) []string {
	str = strings.ReplaceAll(str, " ", "")
	if str == "" {
		return nil
	}
	return strings.Split(str, ",")
}
