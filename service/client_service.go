package service

import (
	"context"
	"errors"
	"github.com/itlemon/oauth2-go/models"
)

var (
	ErrClientNotExist = errors.New("clientId is not exist")
	ErrClientSecret   = errors.New("clientSecret is invalid")
)

type ClientDetailsService interface {
	// GetClientDetailByClientId 根据clientId加载并验证客户端信息
	GetClientDetailByClientId(ctx context.Context, clientId, clientSecret string) (*models.ClientDetails, error)
}

type InMemoryClientDetailsService struct {
	clientDetailsDict map[string]*models.ClientDetails
}

func (service *InMemoryClientDetailsService) GetClientDetailByClientId(ctx context.Context, clientId, clientSecret string) (*models.ClientDetails, error) {
	clientDetails, ok := service.clientDetailsDict[clientId]
	if ok {
		if clientDetails.ClientSecret == clientSecret {
			return clientDetails, nil
		} else {
			return nil, ErrClientSecret
		}
	} else {
		return nil, ErrClientNotExist
	}
}

func NewInMemoryClientDetailsService(clientDetailsList []*models.ClientDetails) *InMemoryClientDetailsService {
	clientDetailsDict := make(map[string]*models.ClientDetails)
	if clientDetailsList != nil {
		for _, clientDetails := range clientDetailsList {
			clientDetailsDict[clientDetails.ClientId] = clientDetails
		}
	}

	return &InMemoryClientDetailsService{clientDetailsDict: clientDetailsDict}
}
