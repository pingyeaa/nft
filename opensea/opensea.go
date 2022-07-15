package opensea

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nft/types"
)

func GetStats(slug string) (*types.OpenseaStats, error) {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/collection/%s/stats", slug)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var stats *types.OpenseaStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
