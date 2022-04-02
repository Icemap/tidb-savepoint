package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Player struct {
	ID    uint `gorm:"primarykey;column:id"`
	Coins int  `gorm:"column:coins"`
	Goods int  `gorm:"column:goods"`
}

func (Player) TableName() string {
	return "savepoint_player"
}

func main() {
	fmt.Printf("\n\nMySQL:\n")
	dbSavepoint("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	fmt.Printf("\n\nTiDB:\n")
	dbSavepoint("root:@tcp(127.0.0.1:4000)/test")
}

func dbSavepoint(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移 schema
	db.AutoMigrate(&Player{})

	db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&Player{ID: 1, Coins: 1, Goods: 1})

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&Player{ID: 2, Coins: 1, Goods: 1})
			return errors.New("rollback player2") // Rollback player2
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&Player{ID: 3, Coins: 1, Goods: 1})
			return nil
		})

		return nil
	})

	var players []Player
	db.Find(&players)

	for _, player := range players {
		fmt.Printf("id: %d, coins: %d, goods: %d\n", player.ID, player.Coins, player.Goods)
	}
}
