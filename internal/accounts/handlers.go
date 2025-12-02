package accounts

import (
	"log/slog"
	"net/http"

	"github.com/dimplesY/goose_test/internal/helper"
	"github.com/dimplesY/goose_test/internal/json"
)

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(service AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (handler *AccountHandler) LoginByNameAndPassword(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest

	if err := json.Read(r, &loginRequest); err != nil {
		slog.Info("登录失败", "error", err)
		json.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := handler.service.LoginByNameAndPassword(loginRequest.Name, loginRequest.Password)

	if err != nil {
		slog.Info("登录失败", "error", err)
		json.Write(w, http.StatusBadRequest, err.Error())
		return
	}

	token := helper.GenerateToken(account.Name)

	slog.Info("登录成功", "token", token)

	json.Write(w, http.StatusOK, &LoginResponse{Token: token})

}
