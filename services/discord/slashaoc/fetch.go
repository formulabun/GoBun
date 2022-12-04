package slashaoc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type member struct {
	Score int    `json:"local_score"`
	Name  string `json:"name"`
	Stars int    `json:"stars"`
}

type members []member

type response struct {
	Members map[string]member `json:"members"`
}

func (mems members) Less(i, j int) bool {
	return mems[i].Score < mems[j].Score
}

func (mems members) Len() int {
  return len(mems)
}

func (mems members) Swap(i, j int) {
  mems[i], mems[j] = mems[j], mems[i]
}

func fetchScores() (members, error) {
	url := "https://adventofcode.com/2022/leaderboard/private/view/2605039.json"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("cookie", "session=53616c7465645f5ff71dccb0e6ac8dd210b40cfcac57b1a91f028c3ccf3125a982cff5b6294b1972ca6a9f0dde42a87c9ae3e76257f0eaaffe59e765f34aa503")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data response
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Could not unmarshal: %s\n", err)
	}

	result := make([]member, len(data.Members))
	i := 0
	for _, member := range data.Members {
		result[i] = member
		i++
	}

	sort.Sort(sort.Reverse(members(result)))

	return result, nil
}
