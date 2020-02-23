package tools

import "server/db/models"

func GetUserMap(users []*models.User) (map[uint]models.User, error) {
	user_map := map[uint]models.User{}
	for _, user := range users {
		user_map[user.ID] = *user
	}
	return user_map, nil
}
