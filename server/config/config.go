package config

import "regexp"

const (
	AvatarWidth    = uint(180)
	AvatarHeight   = uint(180)
	MaxRequestSize = 3 * 1000 * 1000 // 3Mb
)

// regex
const (
	userNameRegex    = `^[a-z0-9_]{2,10}$`
	displayNameRegex = `^.{1,50}$`
)

// 初期化コストはなくすため、グローバル保持
var UserNameRegexp = regexp.MustCompile(userNameRegex)
var DisplayNameRegexp = regexp.MustCompile(displayNameRegex)
