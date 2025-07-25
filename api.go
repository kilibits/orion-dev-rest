package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kilibits/orion-dev-rest/models"
	_ "github.com/mattn/go-sqlite3"
	gorp "gopkg.in/gorp.v1"
)

var dbMap *gorp.DbMap

func dbInit() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "scraperwiki.sqlite")
	if err != nil {
		log.Fatalf("Error Opening Database -> %v", err.Error())
	}

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	return dbMap
}

func getAllProfiles(c *gin.Context) {
	var profiles []models.Profile
	_, err := dbMap.Select(&profiles, "SELECT * FROM profile")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch profiles"})
		return
	}
	c.JSON(200, profiles)
}

func getProfile(c *gin.Context) {
	var profile models.Profile
	id := c.Param("profile_id")
	err := dbMap.SelectOne(&profile, "SELECT * FROM profile WHERE id = ?", id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(200, profile)
}

func getByParty(c *gin.Context) {
	var profiles []models.Profile
	party := c.Param("party")
	_, err := dbMap.Select(&profiles, "SELECT * FROM profile WHERE [group] = ?", party)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch profiles by party"})
		return
	}
	c.JSON(200, profiles)
}

func getByConstituency(c *gin.Context) {
	var profiles []models.Profile
	area := c.Param("area")
	_, err := dbMap.Select(&profiles, "SELECT * FROM profile WHERE area = ?", area)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch profiles by area"})
		return
	}
	c.JSON(200, profiles)
}

func getEducationHistory(c *gin.Context) {
	var edu []models.EducationHistory
	id := c.Param("id")
	_, err := dbMap.Select(&edu, `
        SELECT institution, level, award, [from], [to]
        FROM education_history
        WHERE mp_id = ?
    `, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch education history"})
		return
	}
	c.JSON(200, edu)
}

func main() {
	dbMap = dbInit()
	defer dbMap.Db.Close()

	app := gin.Default()
	app.GET("/profiles", getAllProfiles)
	app.GET("/profiles/:profile_id", getProfile)
	app.GET("/profiles/party/:party", getByParty)
	app.GET("/profiles/area/:area", getByConstituency)
	app.GET("/education/:id", getEducationHistory)
	//app.GET("/profileImages", downloadImages)
	app.Run(":9000")
}
