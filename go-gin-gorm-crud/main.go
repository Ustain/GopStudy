package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	Name    string `unique;gorm:"type:varchar(20); not null" json:"name" binding:"required"`
	Phone   int    `gorm:"type:int; not null" json:"phone" binding:"required"`
	Email   string `gorm:"type:varchar(20); not null" json:"email" binding:"required"`
	Pwd     string `gorm:"type:varchar(20); not null" json:"pwd" binding:"required"`
	Address string `gorm:"type:varchar(20); not null" json:"address" binding:"required"`
}

func main() {
	fmt.Println("ass")

	dsn := "root:123456@tcp(127.0.0.1:3306)/go-gin-gorm-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败")
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	db.AutoMigrate(&User{})

	r := gin.Default()

	//增加接口
	r.POST("user/add", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(200, gin.H{
				"msg":   "add失败",
				"error": err.Error(),
				"user":  gin.H{},
				"code":  400,
			})
		} else {
			//操作数据库
			db.Create(&user)

			c.JSON(200, gin.H{
				"msg":  "add成功",
				"user": user,
				"code": 200,
			})
		}
	})

	//删除接口
	r.DELETE("user/delete/:id", func(c *gin.Context) {
		var user User
		id := c.Param("id")

		if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
			c.JSON(200, gin.H{
				"error": "找不到要删除的id",
				"code":  400,
			})
		} else {
			db.Delete(&user)
			c.JSON(200, gin.H{
				"msg":  "删除成功",
				"user": user,
				"code": 200,
			})
		}

	})

	r.Run(":3000")

}
