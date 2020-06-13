package access_token

import (
	"strings"

	"github.com/edilbertloquine/go-microservices/oauth-api/src/utils/errors"
)

// Repository -
type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

// Service -
type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

// NewService -
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
