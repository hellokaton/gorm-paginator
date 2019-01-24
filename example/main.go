package main

import (
	"fmt"
	"strconv"

	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// User 用户
type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

func main() {
	db, err := gorm.Open("sqlite3", "example.db")
	if err == nil {
		db.AutoMigrate(&User{})
		count := 0
		db.Model(User{}).Count(&count)
		if count == 0 {
			db.Create(User{ID: 1, UserName: "biezhi"})
			db.Create(User{ID: 2, UserName: "rose"})
			db.Create(User{ID: 3, UserName: "jack"})
			db.Create(User{ID: 4, UserName: "lili"})
			db.Create(User{ID: 5, UserName: "bob"})
			db.Create(User{ID: 6, UserName: "tom"})
			db.Create(User{ID: 7, UserName: "anny"})
			db.Create(User{ID: 8, UserName: "wat"})
			fmt.Println("Insert OK!")
		}
	} else {
		fmt.Println(err)
		return
	}

	var users []User

	pagination.Paging(&pagination.Param{
		DB:      db.Where("id > ?", 0),
		Page:    1,
		Limit:   3,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &users)

	fmt.Println("users:", users)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))
		var users []User

		paginator := pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    page,
			Limit:   limit,
			OrderBy: []string{"id desc"},
			ShowSQL: true,
		}, &users)
		c.JSON(200, paginator)
	})

	r.Run()

}
