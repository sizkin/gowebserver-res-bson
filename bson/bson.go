package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/gzip"
    "gopkg.in/mgo.v2/bson"
)

type Image struct {
    Name string "name"
    Path string "path,omitempty"
}

func main() {
    r := gin.Default()
    r.Use(gzip.Gzip(gzip.DefaultCompression))
    r.GET("/bson", func(c *gin.Context) {
        //data, err := bson.Marshal(&Image{Name: "people_top_hair_000"}) 
        data, err := bson.Marshal(gin.H{"name": "people_top_hair_000", "path": "assect/images/hair"}) 
        if err != nil {
            panic(err)
        }
        c.Data(http.StatusOK, "application/bson", data)
    })

    r.GET("/json", func(c *gin.Context) {
        c.JSON(http.StatusOK, &Image{Name: "people_top_hair_000"})
    })

    r.Static("/html", "./html")
    r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

