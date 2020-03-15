package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const port = 8337

func main() {
	fmt.Println("Starting sungrabber...")

	// werks4me - bash script will restart every 24hrs
	time.AfterFunc(24*time.Hour, func() {
		os.Exit(0)
	})

	date := time.Now().Format("2006-01-02")
	err := downloadVid("latest.mp4", "ec-belmar", "0650", date)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting HTTP server...")
	http.HandleFunc("/latest", latestHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func latestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "latest.mp4")
}

func downloadVid(filepath string, camera string, time string, date string) (err error) {
	url := "https://camrewinds.cdn-surfline.com/%s/%s.%s.%s.mp4"
	url = fmt.Sprintf(url, camera, camera, time, date)
	fmt.Printf("Downloading %s...\n", url)

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Got non-OK status: %s", resp.Status)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
