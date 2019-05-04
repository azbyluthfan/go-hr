package main

import (
	"fmt"
	employeeQuery "github.com/azbyluthfan/go-hr/modules/employees/query"
	employeeUseCase "github.com/azbyluthfan/go-hr/modules/employees/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type Service struct {
	EmployeeUseCase employeeUseCase.EmployeeUseCase
}

func MakeHandler() *Service {

	// create db connection
	maxAttempt := 10
	waitTime := 10

	db, err := ConnectDb()
	connected := false

	if err != nil {
		for i := 1; i <= maxAttempt; i++ {
			fmt.Println(err)
			db, err = ConnectDb()

			if err == nil {
				connected = true
				break
			}

			// Set sleep for a moment to make interval connection
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}

	// exiting after can not make db connection within retry attempt
	if !connected {
		os.Exit(1)
	}

	employeeQuery := employeeQuery.NewEmployeeQueryMysql(db)
	employeeUC := employeeUseCase.NewEmployeeUseCase(employeeQuery)

	return &Service{
		EmployeeUseCase: employeeUC,
	}
}

func ConnectDb() (*sqlx.DB, error) {

	db, err := sqlx.Connect("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME"))
	return db, err
}