package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/masamichhhhi/quest-api/docs"
	resp "github.com/nicklaw5/go-respond"
	httpSwagger "github.com/swaggo/http-swagger"
)

type (
	room struct {
		ID           uint64    `json:"id"`
		OwnerID      uint64    `json:"owner_id"`
		OwnerAddress string    `json:"owner_address"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		EventCode    string    `json:"event_code"`
		Address      string    `json:"address"`
		CreateTxHash string    `json:"create_tx_hash"`
		IsPrivate    bool      `json:"is_private"`
		WeiBalance   uint64    `json:"wei_balance"`
		WeiPrize     uint64    `json:"wei_prize"`
		StartTime    time.Time `json:"start_time"`
		EndTime      time.Time `json:"end_time"`
		Active       bool      `json:"active"`
	}

	rooms []room
)

var db *gorm.DB

// API Documents
// @title Sample API
// @version 1.0.0
func main() {
	log.Println("Starting API Server!")
	db = connectDB()
	defer db.Close()

	router := chi.NewRouter()

	router.Get("/root", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("root"))
	})

	router.Get("/rooms", getRooms)

	router.Get("/swagger/*", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", router)
}

// @Tags Room
// @Accept json
// @Summary ルーム情報を取得します
// @Success 200 {array} room
// @Router /rooms [get]
func getRooms(w http.ResponseWriter, r *http.Request) {
	rooms := []room{}
	db.Find(&rooms)

	resp.NewResponse(w).Ok(rooms)
}

func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "test_db"
	option := "?charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + option

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
