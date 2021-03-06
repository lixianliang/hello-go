package main

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
    <script src="/assets/app.js"></script>
	</head>
	<body>
	  <h1 style="color:red;">Welcome, Ginner!</h1>
	  </body>
	  </html>
`))

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// use pusher.Push() to do server push
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	//router.Run(":8080")
	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
