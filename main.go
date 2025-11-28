package main

import (
	"context"
	"fmt"

	"github.com/dimplesY/goose_test/db"
	"github.com/jackc/pgx/v5"
)

func main() {

	conn, _ := pgx.Connect(context.Background(), "postgres://postgres:123456@localhost:5432/goose_test")
	defer conn.Close(context.Background())
	queries := db.New(conn)

	user, _ := queries.CreateUser(context.Background(), db.CreateUserParams{
		Name:     "test",
		Email:    "test@qq.com",
		Password: "123456",
	})

	fmt.Println(user)

	u1, _ := queries.GetAccountById(context.Background(), user.ID)
	fmt.Println(u1)

	_ = queries.UpdateUser(context.Background(), db.UpdateUserParams{
		Name:     "test2",
		Email:    "test2@qq.com",
		Password: "1234567",
		ID:       user.ID,
	})

	u2, _ := queries.GetAccountById(context.Background(), user.ID)
	fmt.Println(u2)

	_ = queries.DeleteUser(context.Background(), user.ID)

	u3, _ := queries.GetAccountById(context.Background(), user.ID)
	fmt.Println(u3)

}
