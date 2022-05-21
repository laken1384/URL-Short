package controllers

import (
    "Backend/database"
    "net/http"
    // "log"
    "strings"
    "math/rand"
    // "github.com/google/uuid"
    "github.com/gin-gonic/gin"
)

var db = database.CreateDatabase()

func GetRedirect(c *gin.Context) {
    url_id := c.Param("url_id")
    var URL string
    db.Select("URL").Table("URL_Shortener").Where("Id = ?", url_id).Scan(&URL)
    c.Redirect(http.StatusFound, URL)
}

func SetURL(c *gin.Context) {
    var response InsertBody
    c.Bind(&response)
    uuid := randomString(5)
    // if err != nil {
    //     log.Fatal(err)
    // }
    var save_url []URL_Sortend
    urlcheck := response.URL
    if !strings.HasPrefix(response.URL, "http") {
        urlcheck = strings.Join([]string{"http", response.URL}, "")
    }
    save_url = append(save_url, URL_Sortend{uuid, response.ExpireAt, urlcheck,})
    db.Table("URL_Shortener").Save(&save_url)
    ip := c.ClientIP()
    SortURL := "http://" + ip + ":8687/" + uuid
    ReturnSortURL := Return_url{uuid, SortURL,}
    c.JSON(http.StatusOK, ReturnSortURL)
}
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
    }
func randomString(l int) string {
    bytes := make([]byte, l)
    for i := 0; i < l; i++ {
    bytes[i] = byte(randomInt(65, 90))
    }
    var Id string
    db.Select("Id").Table("URL_Shortener").Where("Id = ?", string(bytes)).Scan(&Id)
    if Id == string(bytes) {
        return randomString(l)
    }
    return string(bytes)
    }

type Return_url struct {
    Id string `json:"id"`
    ShortUrl string `json:"shortUrl"`
}

type URL_Sortend struct {
    Id  string  `json:"Id" gorm:"column:Id"`
    Expireat string  `json:"ExpireAt" gorm:"ExpireAt"`
    URL string `json:"URL" gorm:"URL"`
}

type InsertBody struct {
    URL string `json:"url"`
    ExpireAt string `json:"expireAt"`
}
