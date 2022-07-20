package function

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-numb/go-utilities/notify"
)

var (
	WEBHOOK_ID    = os.Getenv("WEBHOOK_ID")
	WEBHOOK_TOKEN = os.Getenv("WEBHOOK_TOKEN")

	ChannelID = os.Getenv("ChannelID")
	PostName  = os.Getenv("PostName")
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func init() {
	WEBHOOK_ID = os.Getenv("WEBHOOK_ID")
	WEBHOOK_TOKEN = os.Getenv("WEBHOOK_TOKEN")

	ChannelID = os.Getenv("ChannelID")
	PostName = os.Getenv("PostName")
}

// 実行関数
func Pingpong(w http.ResponseWriter, r *http.Request) {
	results, err := dosomething()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		log.Println(err)
		return
	}

	if err := toDiscord(results); err != nil {
		http.Error(w, results, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	fmt.Fprintf(w, "ok")
	log.Println("ok to log")
}

// 適宜変更
func dosomething() (string, error) {

	return "results", nil
}

func toDiscord(results string) error {
	discord := &notify.Discord{
		ID:        WEBHOOK_ID,
		Token:     WEBHOOK_TOKEN,
		ChannelID: ChannelID,
		PostName:  PostName,
		Message:   results,
	}
	if err := discord.Send(); err != nil {
		return err
	}

	return nil
}
