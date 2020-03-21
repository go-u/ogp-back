package config

import (
	"regexp"
)

const (
	MaxRequestSize = 3 * 1000 * 1000 // 3Mb
)

//regex
const (
	//userNameRegex    = `^[a-z0-9_]{2,10}$`
	displayNameRegex = `^.{1,50}$`
)

//// 都度コンパイルのコストを無くすためグローバル保持
var (
	//UserNameRegexp = regexp.MustCompile(userNameRegex)
	DisplayNameRegexp = regexp.MustCompile(displayNameRegex)
)
