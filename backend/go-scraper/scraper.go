package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	downloadPlayerJson()
}

func downloadPlayerJson() {
	client := &http.Client{}
	reqUrl := "https://stats.nba.com/stats/playerindex?College=&Country=&DraftPick=&DraftRound=&DraftYear=&Height=&Historical=1&LeagueID=00&Season=2023-24&SeasonType=Regular%20Season&TeamID=0&Weight="

	req, err := http.NewRequest("GET", reqUrl, nil)

	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.8")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Origin", "https://www.nba.com")
	req.Header.Add("Referer", "https://www.nba.com")
	req.Header.Add("Sec-Ch-Ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Brave\";v=\"120\"")
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Gpc", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)

	if err != nil {
		log.Fatal("Error creating gzip reader:", err)
	}
	defer gz.Close()

	file, err := os.Create("players.json")

	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	_, err = io.Copy(file, gz)
	if err != nil {
		log.Fatal("Error copying data:", err)
	}
}
