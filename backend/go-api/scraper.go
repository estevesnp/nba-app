package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PlayerResponse struct {
	ResultSets []struct {
		RowSet [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}

func GetPlayerData() ([]Player, error) {
	data, err := downloadPlayerData()
	if err != nil {
		return nil, err
	}

	players, err := parsePlayerData(data)
	if err != nil {
		return nil, err
	}

	return players, nil

}

func downloadPlayerData() ([]byte, error) {
	reqUrl := "https://stats.nba.com/stats/playerindex?College=&Country=&DraftPick=&DraftRound=&DraftYear=&Height=&Historical=1&LeagueID=00&Season=2023-24&SeasonType=Regular%20Season&TeamID=0&Weight="

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	headers := map[string]string{
		"Accept":             "*/*",
		"Accept-Encoding":    "gzip, deflate, br",
		"Accept-Language":    "en-US,en;q=0.8",
		"Connection":         "keep-alive",
		"Origin":             "https://www.nba.com",
		"Referer":            "https://www.nba.com",
		"Sec-Ch-Ua":          "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Brave\";v=\"120\"",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "\"Linux\"",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"Sec-Gpc":            "1",
		"User-Agent":         "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(gz)
	data := buf.Bytes()

	return data, nil

}

func parsePlayerData(data []byte) ([]Player, error) {
	players := []Player{}

	var r PlayerResponse

	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	for _, player := range r.ResultSets[0].RowSet {

		if len(player) < 26 {
			continue
		}

		if player[25] != "2023" {
			continue
		}

		players = append(players, Player{
			Id:       int(player[0].(float64)),
			Name:     fmt.Sprintf("%s %s", player[2], player[1]),
			Position: player[11].(string),
			Team:     player[9].(string),
		})
	}

	if len(players) == 0 {
		return nil, fmt.Errorf("no players found")
	}

	return players, nil
}
