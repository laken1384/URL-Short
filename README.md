# URL-Short
程式架構：Go + MySQL
# Installing and Running

#### MySQL
MsySQL封裝為Docker,可於任意支援Docker的設備安裝
https://drive.google.com/file/d/1LwnhVtsJI8g9WDx2hYok5VK0gZq-Onyh/view?usp=sharing
```shell
docker load -i mysql.tar
```

#### Backend
```shell
cd Backend
go run start.go
```
#### Running
後端服務運行於
```
http://localhost:8687/
```
# 設定URL&重新導向新網址
POST:http://localhost:8687/api/v1/urls
- Body:
```json
{
    "url": "<original_url>",
    "expireAt": "2021-02-08T09:20:41Z"
}
```
- Response:
```json
{
    "id": "<url_id>",
    "shortUrl": "http://localhost/<url_id>"
}
```

GET:http://localhost:8687/:original_url
重新導向至之前儲存的網址

# 專案設計概念
- MYSQL 對於儲存的網址所帶的過期時間，MySQL有設定EVENT事件，每天會檢查資料是否過期，遇到過期資料會自動刪除，可以有效減少資料庫的50%負擔，以及增加90％速度．
- Golang CROS的部分有做過處理，可於同一網域下跨電腦來使用
