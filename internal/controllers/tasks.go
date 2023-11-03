package controllers

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)
type TaskController struct {}

type task struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Contents string `json:"contents"`
	Deadline time.Time `json:"date"`
}

var tasks = []task {
	{Id: "1", Title: "Driving", Contents: "2 hours practice", Deadline: time.Date(2023, time.September, 5, 15, 0, 0, 0, time.UTC) },
	{Id: "2", Title: "Haircut", Contents: "in Chorzow", Deadline: time.Date(2023, time.September, 29, 13, 0, 0, 0, time.UTC) },
	{Id: "3", Title: "Doctor", Contents: "Galerea", Deadline: time.Date(2023, time.September, 19, 15, 50, 0, 0, time.UTC) },
}

func (u TaskController) GetMany (c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
	return
}

