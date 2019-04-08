package instagram

import (
	"github.com/ahmdrz/goinsta"
)

type Instagram struct {
	insta *goinsta.Instagram
}

func New(username, password string) (*Instagram, error) {
	insta := goinsta.New(username, password)
	err := insta.Login()
	if err != nil {
		return nil, err
	}
	return &Instagram{
		insta: insta,
	}, nil
}

func (i *Instagram) Export(filePath string) error {
	return i.insta.Export(filePath)
}

func Import(filePath string) (*Instagram, error) {
	i, err := goinsta.Import(filePath)
	if err != nil {
		return nil, err
	}
	return &Instagram{
		insta: i,
	}, nil
}

func (i *Instagram) Followings() []User {
	output := make([]User, 0)
	users := i.insta.Account.Following()
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

func (i *Instagram) Followers() []User {
	output := make([]User, 0)
	users := i.insta.Account.Followers()
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
