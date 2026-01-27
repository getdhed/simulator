package main

import (
	"context"
	"erp/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init(){
	if err := godotenv.Load(); err!=nil{
		log.Println(".env не найден, используем системные переменные")
	}
}
func main() {
	con:=os.Getenv("SUPABASE_URL")
	ctx := context.Background()
	if err := sql.Connection(ctx, con); err != nil {
		panic(err)
	}


	fmt.Println("test code for connection!!!!!!!!!!!!!!!")
	fmt.Println("succes!")
}
