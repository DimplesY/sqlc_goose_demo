package accounts

import (
	"context"

	database "github.com/dimplesY/goose_test/internal/db"
)

type AccountService interface {
	// LoginByNameAndPassword 通过用户名和密码登录
	LoginByNameAndPassword(name, password string) (*database.Account, error)
}

type accountService struct {
	queries *database.Queries
}

func NewAccountService(queries *database.Queries) AccountService {
	return &accountService{
		queries: queries,
	}
}

func (srv accountService) LoginByNameAndPassword(name, password string) (*database.Account, error) {

	ctx := context.Background()

	account, err := srv.queries.GetAccountByName(ctx, name)

	if err != nil || account.Password != password {
		return nil, err
	}

	return &account, nil
}
