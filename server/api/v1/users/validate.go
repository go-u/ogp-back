package users

import (
	"server/config"
)

func isValidUsername(username string) bool {
	return config.UserNameRegexp.MatchString(username)
}

func isValidDisplayName(displayname string) bool {
	return config.DisplayNameRegexp.MatchString(displayname)
}
