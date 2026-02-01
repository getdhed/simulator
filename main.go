package main

import (
	"context"
	"erp/models"
	"erp/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env не найден, используем системные переменные")
	}
}
func main() {
	constr := os.Getenv("DATABASE_URL")
	ctx := context.Background()

	con, err := sql.Connection(ctx, constr)
	if err != nil {
		log.Fatal("Ошибка подключения: ", err)
	}

	fmt.Println("succes!")
	if err := sql.CreateTable(ctx, con); err != nil {
		fmt.Println("что то пошло не так")
	}

	mp := models.NewUsers()
	mp.InsertMap(20)
	wm := models.NewWorkManager()
	sqlStore := sql.NewPgUserStore(con)
	mp.ForEach(func(u *models.User) { sqlStore.CreateUser(ctx, u) })
	wm.StartShift(mp)
	mp.PrintMap()

}
