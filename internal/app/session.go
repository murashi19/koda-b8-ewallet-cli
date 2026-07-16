package app

import "github.com/murashi19/koda-b8-ewallet-cli/internal/models"

type Session struct {
	CurrentUser *models.User
}

func (s *Session) Login(user models.User) {
	s.CurrentUser = &user
}

func (s *Session) Logout() {
	s.CurrentUser = nil
}

func (s *Session) IsLoggedIn() bool {
	return s.CurrentUser != nil
}
