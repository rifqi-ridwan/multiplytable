package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	t := template.Must(template.New("main").Funcs(template.FuncMap{"mul": Mul}).ParseGlob("templates/*"))

	r.SetHTMLTemplate(t)

	r.GET("/", func(c *gin.Context) {
		data := []string{}
		var x, y = 1, 1
		for x <= 9 {
			for y <= 9 {
				data = append(data, fmt.Sprintf("%v x %v = %v", x, y, x*y))
				y++
			}
			y = 1
			x++
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{"data": data})
	})

	log.Fatal(r.Run(":" + port))
}

func Mul(param1 int, param2 int) int {
	return param1 * param2
}
