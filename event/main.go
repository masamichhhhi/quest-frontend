package main

import (
	"context"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/masamichhhhi/quest-event/bindings"
)

type Room struct {
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

type RoomCreatedEvent struct {
	Creator        common.Address
	Room           common.Address
	DepositedValue *big.Int
}

func main() {
	db := connectDB()
	defer db.Close()

	client, err := ethclient.Dial(os.Getenv("GETH_URL"))
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	topics := map[string]common.Hash{
		"RoomCreated": crypto.Keccak256Hash([]byte("RoomCreated(address,address,uint256")),
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(os.Getenv("ROOM_FACTORY_ADDRESS"))},
		Topics: [][]common.Hash{{
			topics["RoomCreated"],
		}},
	}

	event := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, event)
	if err != nil {
		log.Fatal(err)
	}

	roomFactroyAbi, err := abi.JSON(strings.NewReader(string(bindings.RoomFactoryABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println(err)
			close(event)
		case vLog := <-event:
			if len(vLog.Topics) == 0 {
				log.Println("Topicsが存在しません")
			}

			switch vLog.Topics[0] {
			case topics["RoomCreated"]:
				var roomCreatedEvent RoomCreatedEvent
				if err = roomFactroyAbi.UnpackIntoInterface(&roomCreatedEvent, "RoomCreated", vLog.Data); err != nil {
					log.Println("unpackに失敗しました")
					continue
				}

				room := Room{
					Address:      roomCreatedEvent.Room.Hex(),
					CreateTxHash: vLog.TxHash.Hex(),
					OwnerAddress: common.BytesToAddress(vLog.Topics[1].Bytes()).Hex(),
					WeiBalance:   roomCreatedEvent.DepositedValue.Uint64(),
					EventCode:    getEventCode(4),
					OwnerID:      1,
				}

				if err := db.Create(&room).Error; err != nil {
					log.Println("Failed create room: %v\n", err)
					continue
				}
			}
		}
	}
}

func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "quest_db"
	option := "?charset=utf8&parseTime=True"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + option

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func getEventCode(n int) string {
	const allowedStrings = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	eventCode := make([]byte, n)
	for i := range eventCode {
		eventCode[i] = allowedStrings[rand.Intn(len(allowedStrings))]
	}

	return (string(eventCode))
}
