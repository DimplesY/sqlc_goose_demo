package main

import (
	"context"
	"fmt"

	"github.com/dimplesY/goose_test/db"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	conn, _ := pgx.Connect(ctx, "postgres://postgres:123456@localhost:5432/goose_test")

	defer conn.Close(ctx)
	queries := db.New(conn)

	user, _ := queries.CreateUser(ctx, db.CreateUserParams{
		Name:     "test",
		Email:    "test@qq.com",
		Password: "123456",
	})

	fmt.Println(user)

	u1, _ := queries.GetAccountById(ctx, user.ID)
	fmt.Println(u1)

	_ = queries.UpdateUser(ctx, db.UpdateUserParams{
		Name:     "test2",
		Email:    "test2@qq.com",
		Password: "1234567",
		ID:       user.ID,
	})

	u2, _ := queries.GetAccountById(ctx, user.ID)
	fmt.Println(u2)

	_ = queries.DeleteUser(ctx, user.ID)

	u3, _ := queries.GetAccountById(ctx, user.ID)
	fmt.Println(u3)

}
