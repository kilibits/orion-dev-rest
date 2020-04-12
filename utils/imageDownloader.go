package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	gin "github.com/gin-gonic/gin"
	"github.com/kilibits/orion-dev-rest/models"
)

//Temporary function to download images and store them
//Stored images will then be stored on a CDN instead of calling from parliament site

//TODO: Add http request retries to avoid timeouts
func downloadImages(c *gin.Context) {

	var links []models.Profile
	_, err := dbMap.Select(&links, "SELECT id, image FROM profile")

	if err != nil {
		log.Fatalf("Select statement failed -> %v", err.Error())
	}

	for _, profile := range links {
		var _, err = os.Stat("./images/" + strconv.Itoa(profile.Id) + ".jpeg")

		if os.IsNotExist(err) {
			file, err := os.Create("./images/" + strconv.Itoa(profile.Id) + ".jpeg")
			if err != nil {
				log.Fatalf("Failed to create image -> %v", err.Error())
			}
			defer file.Close()

			response, err := http.Get(profile.Image)

			if err != nil {
				log.Fatalf("Unable to get url -> %v", err.Error())
			}

			defer response.Body.Close()

			io.Copy(file, response.Body)
		}
	}
}
