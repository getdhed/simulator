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

func init() {
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
	sqlStore := models.NewPgUserStore(con)
	wm := models.NewWorkManager(sqlStore)
	slc, err := sqlStore.ListUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	mp := models.UsersFromDB(slc)
	// mp := models.NewUsers()
	// mp.InsertMap(20)

	mp.CreateAll(ctx, sqlStore)

	wm.StartShift(ctx, mp)
	//mp.PrintMap()

}
