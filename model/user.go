package model

type MyUser struct {
	name string
}

func (s *MyUser) GetUserName() string {
	return s.name
}
