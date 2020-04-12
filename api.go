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
		log.Fatalf("Select statement failed -> %v", err.Error())
	}
	c.JSON(200, profiles)
}

func getProfile(c *gin.Context) {

	var profile models.Profile
	id := c.Params.ByName("profile_id")
	_, err := dbMap.Select(&profile, "SELECT * FROM profile WHERE id = ?", id)

	if err != nil {
		log.Fatalf("Select statement failed -> %v", err.Error())
	}

	c.JSON(200, profile)
}

func getByParty(c *gin.Context) {
	var profiles []models.Profile
	id := c.Params.ByName("party")
	_, err := dbMap.Select(&profiles, "SELECT * FROM profile WHERE group = ?", id)

	if err != nil {
		log.Fatalf("Select statement failed -> %v", err.Error())
	}

	c.JSON(200, profiles)
}

func getByConstituency(c *gin.Context) {
	var profiles []models.Profile
	id := c.Params.ByName("area")
	_, err := dbMap.Select(&profiles, "SELECT * FROM profile WHERE area = ?", id)

	if err != nil {
		log.Fatalf("Select statement failed -> %v", err.Error())
	}
	c.JSON(200, profiles)
}

func getEducationHistory(c *gin.Context) {
	var edu []models.EducationHistory
	id := c.Params.ByName("id")
	_, err := dbMap.Select(&edu, "SELECT institution, level, award, [from], [to] FROM profile JOIN education_history ON profile.id = education_history.mp_id WHERE profile.name id = '%?%'", id)

	if err != nil {
		log.Fatalf("Select statement failed -> %v", err.Error())
	}

	c.JSON(200, edu)
}

func main() {
	dbMap = dbInit()
	app := gin.Default()
	app.GET("/profiles", getAllProfiles)
	app.GET("/profiles/:profile_id", getProfile)
	app.GET("/education/:id", getEducationHistory)
	//app.GET("/profileImages", downloadImages)
	app.Run(":9000")
}
