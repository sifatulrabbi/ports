package models

type User struct {
	Username  string `bson:"username"`
	Password  string `bson:"password"`
	Email     string `bson:"email"`
	Fullname  string `bson:"fullname"`
	AvatarURL string `bson:"avatarURL"`
	Bio       string `bson:"bio"`
}

func (u *User) Save() {

}

func (u *User) GetById() {

}
