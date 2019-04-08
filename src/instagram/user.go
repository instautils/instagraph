package instagram

type User struct {
	ID         int64  `json:"pk"`
	Username   string `json:"username"`
	ProfilePic string `json:"profile_pic"`
}

func (u User) Followings(i *Instagram) []User {
	output := make([]User, 0)
	user, err := i.insta.Profiles.ByID(u.ID)
	if err != nil {
		return nil
	}
	users := user.Following()
	for users.Next() {
		for _, user := range users.Users {
			output = append(output, User{
				ID:         user.ID,
				Username:   user.Username,
				ProfilePic: user.ProfilePicURL,
			})
		}
	}
	return output
}

func (u User) Followers(i *Instagram) []User {
	output := make([]User, 0)
	user, err := i.insta.Profiles.ByID(u.ID)
	if err != nil {
		return nil
	}
	users := user.Followers()
	for users.Next() {
		for _, user := range users.Users {
			output = append(output, User{
				ID:         user.ID,
				Username:   user.Username,
				ProfilePic: user.ProfilePicURL,
			})
		}
	}
	return output
}
