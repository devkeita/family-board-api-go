package main

import (
	"family-board-api/config"
	"family-board-api/registry"
	"family-board-api/usecase"
	"family-board-api/usecase/input"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// envファイルを読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// db初期化
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// DI
	repository := registry.NewRepository(db)
	uu := usecase.NewUserUsecase(repository.User)
	i := &input.LoginWithLine{
		LiffIdToken: "123456789",
	}

	token, err := uu.LoginWithLine(i)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(token)
}
