package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	UserName string `gorm:"unique"`
	Basis    UserCostBasis
}

type UserCostBasis struct {
	gorm.Model
	// User **hasOne** UserCostBasis relationship
	UserID int
	// 钱嘛……还是不能直接用浮点数运算的……
	CurrentCostBasis decimal.Decimal `gorm:"type:numeric"`
	PendingCostBasis decimal.Decimal `gorm:"type:numeric"`
	AccountLimit     decimal.Decimal `gorm:"type:numeric"`
}
