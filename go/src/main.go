package main

import (
	"github.com/aoimaru/0614-project-go-db-ver0.0.1.1/src/lib"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/city/:name", func(c *gin.Context) {
		name := c.Param("name")
		db, err := lib.Connect()
		if err != nil {
			c.JSON(200, gin.H{
				"message": name + "is not exist",
			})
		}
		defer db.Close()
		results := []*lib.City{}
		error := db.Where("Name = ?", name).Find(&results).Error
		if error != nil || len(results) == 0 {
			c.JSON(200, gin.H{
				"message": name + "is not exist",
			})
		}
		c.JSON(200, gin.H{
			"results": results,
		})
	})
	r.GET("/city", func(c *gin.Context) {
		db, err := lib.Connect()
		if err != nil {
			c.JSON(200, gin.H{
				"message": "can't connect",
			})
		}
		defer db.Close()
		results := []*lib.City{}
		error := db.Find(&results).Error
		if error != nil || len(results) == 0 {
			c.JSON(200, gin.H{
				"message": "not exist",
			})
		}
		c.JSON(200, gin.H{
			"results": results,
		})
	})
	r.GET("/countrycode/:countrycode", func(c *gin.Context) {
		countrycode := c.Param("countrycode")
		db, err := lib.Connect()
		if err != nil {
			c.JSON(200, gin.H{
				"message": countrycode + "is not exist",
			})
		}
		defer db.Close()
		results := []*lib.City{}
		error := db.Where("CountryCode = ?", countrycode).Find(&results).Error
		if error != nil || len(results) == 0 {
			c.JSON(200, gin.H{
				"message": countrycode + "is not exist",
			})
		}
		c.JSON(200, gin.H{
			"results": results,
		})
	})

	r.POST("/city", func(c *gin.Context) {
		db, err := lib.Connect()
		if err != nil {
			c.JSON(300, gin.H{
				"message": "can't connect",
			})
		}
		var city []lib.City
		if err := c.BindJSON(&city); err != nil {
			c.JSON(300, gin.H{
				"message": "failed",
			})
		}
		if err := db.Create(&city); err != nil {
			c.JSON(300, gin.H{
				"message": "failed",
			})
		} else {
			c.JSON(300, gin.H{
				"message": "successed",
			})
		}
		db.Close()

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
