package main

import (
	"context"
	"erp/sql"
	"fmt"
)

func main() {
	con := "postgresql://postgres:DydsbRS3bN1228@db.bvwrjsxqzvnezeakqjrj.supabase.co:5432/postgres"
	ctx := context.Background()
	if err := sql.Connection(ctx, con); err != nil {
		panic(err)
	}
	fmt.Println("succes!")
}
