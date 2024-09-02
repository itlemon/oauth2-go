package service

import (
	"context"
	"errors"
	"github.com/itlemon/oauth2-go/models"
)

var (
	ErrUserNotExist = errors.New("username is not exist")
	ErrPassword     = errors.New("password is invalid")
)

type UserDetailsService interface {
	// GetUserDetailByUsername 根据用户名加载并验证用户信息
	GetUserDetailByUsername(ctx context.Context, username, password string) (*models.UserDetails, error)
}

type InMemoryUserDetailsService struct {
	userDetailsDict map[string]*models.UserDetails
}

func (service *InMemoryUserDetailsService) GetUserDetailByUsername(ctx context.Context, username, password string) (*models.UserDetails, error) {
	userDetails, ok := service.userDetailsDict[username]
	if ok {
		if userDetails.Password == password {
			return userDetails, nil
		} else {
			return nil, ErrPassword
		}
	} else {
		return nil, ErrUserNotExist
	}
}

func NewInMemoryUserDetailsService(userDetailsList []*models.UserDetails) *InMemoryUserDetailsService {
	userDetailsDict := make(map[string]*models.UserDetails)
	if userDetailsList != nil {
		for _, userDetails := range userDetailsList {
			userDetailsDict[userDetails.Username] = userDetails
		}
	}

	return &InMemoryUserDetailsService{userDetailsDict: userDetailsDict}
}
