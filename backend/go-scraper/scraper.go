package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Team     string `json:"team"`
}

func main() {
	data, err := downloadPlayerData()
	if err != nil {
		log.Fatal("Error downloading player data:", err)
	}

	players, err := parsePlayerData(data)
	if err != nil {
		log.Fatal("Error parsing player data:", err)
	}

	err = savePlayersToJSON(players)
	if err != nil {
		log.Fatal("Error saving player data:", err)
	}

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

	currentPlayers := []Player{}

	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	resultSets, ok := m["resultSets"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("resultSets not found")
	}

	allPlayers, ok := resultSets[0].(map[string]interface{})["rowSet"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("rowSet not found")
	}

	for _, player := range allPlayers {
		p := player.([]interface{})

		if p[25] != "2023" {
			continue
		}

		currentPlayers = append(currentPlayers, Player{
			Id:       int(p[0].(float64)),
			Name:     fmt.Sprintf("%s %s", p[2], p[1]),
			Position: p[11].(string),
			Team:     p[9].(string),
		})
	}

	if len(currentPlayers) == 0 {
		return nil, fmt.Errorf("no players found")
	}

	return currentPlayers, nil
}

func savePlayersToJSON(players []Player) error {
	data, err := json.Marshal(players)
	if err != nil {
		return err
	}

	file, err := os.Create("players.json")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil

}
