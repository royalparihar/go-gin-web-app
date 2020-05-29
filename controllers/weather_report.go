package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /weather-report:city
func GetWeatherReport(c *gin.Context) {
	city := c.Param("city")
	apiEndpint := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=0cc717bbbfdf0fccb3b14df582c55aad"
	resp, err := http.Get(apiEndpint)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": string(body)})
}
