package main

import (
	"github.com/initiumfund/artis-go/models"
	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main()  {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.UserCostBasis{})
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		Name:     "Ye",
		UserName: "test",
		Basis:    models.UserCostBasis{
			CurrentCostBasis: decimal.NewFromFloat(952.21),
			PendingCostBasis: decimal.NewFromFloat(0.0596),
			AccountLimit:     decimal.NewFromFloat(2500.3054),
		},
	}

	result := db.Create(&user)
	//db.First(&user)
	//user.UserName = "yechs"
	//user.Basis.PendingCostBasis = decimal.NewFromFloat(100000000000.900123)
	//result = db.Save(&user)

	println(user.ID, result.Error, result.RowsAffected)
}

