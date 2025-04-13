package dto

import "github.com/lucasbarbosaalves/go-gateway/internal/domain"

type CreateAccountInput struct {
	Name	string `json:"name" validate:"required"`
	Email	string `json:"email" validate:"required,email"`
}

type AccountResponse struct {
	ID		string `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
	Balance	float64 `json:"balance"`
	APIKey	string `json:"api_key,omitempty"`
	CreatedAt	string `json:"created_at"`
	UpdatedAt	string `json:"updated_at"`
}

func ToAccount(input CreateAccountInput) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) *AccountResponse {
	return &AccountResponse{
		ID: account.ID,
		Name: account.Name,
		Email: account.Email,
		Balance: account.Balance,
		APIKey: account.APIKey,
		CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: account.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
