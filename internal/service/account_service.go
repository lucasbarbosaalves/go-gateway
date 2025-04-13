package service

import (
	"github.com/lucasbarbosaalves/go-gateway/internal/domain"
	"github.com/lucasbarbosaalves/go-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(accountRepository domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: accountRepository,
	}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountInput) (*dto.AccountResponse, error) {

	account := dto.ToAccount(input)
	actualAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err == nil && err != domain.ErrAccountNotFound {
		return nil, err
	}
	if actualAccount != nil {
		return nil, domain.ErrDuplicatedAPIKey
	}

	err = s.repository.Save(account)
	
	if err != nil {
		return nil, err
	}
	return dto.FromAccount(account), nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	if amount == 0 {
		return nil, domain.ErrInvalidAmount
}
	account.AddBalance(amount)

	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}
	return dto.FromAccount(account), nil

}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	reponse := dto.FromAccount(account)
	return reponse, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	reponse := dto.FromAccount(account)
	return reponse, nil
}