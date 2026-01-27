package main

import (
	"context"
	"erp/sql"
	"fmt"
)

func main() {
	con := "postgresql://postgres.bvwrjsxqzvnezeakqjrj:DydsbRS3bN1228@aws-1-eu-west-2.pooler.supabase.com:5432/postgres"
	ctx := context.Background()
	if err := sql.Connection(ctx, con); err != nil {
		panic(err)
	}


	fmt.Println("test code for connection!!!!!!!!!!!!!!!")
	fmt.Println("succes!")
}
