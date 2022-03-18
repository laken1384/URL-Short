package controllers

import (
    "URL-Sort/Backend/database"
    "net/http"
    "log"
    "strings"
    "github.com/google/uuid"
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
    uuid, err := uuid.NewUUID()
    if err != nil {
        log.Fatal(err)
    }
    var save_url []URL_Sortend
    urlcheck := response.URL
    if !strings.HasPrefix(response.URL, "http://") {
        urlcheck = strings.Join([]string{"http://", response.URL}, "")
    }
    save_url = append(save_url, URL_Sortend{uuid.String(), response.ExpireAt, urlcheck,})
    db.Table("URL_Shortener").Save(&save_url)
    ip := c.ClientIP()
    SortURL := "http://" + ip + ":8687/" + uuid.String()
    ReturnSortURL := Return_url{uuid.String(), SortURL,}
    c.JSON(http.StatusOK, ReturnSortURL)
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
