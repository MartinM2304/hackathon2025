import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Pair struct {
	Turn    int `json:"turn"`
	Entropy int `json:"entropy"`
}

func calculateDirectionEntropy(userVotes map[int][4]int) []Pair {
	var result []Pair

	for turn, votes := range userVotes {
		totalVotes := [4]float64{0, 0, 0, 0}
		total := 0.0

		for dir := 0; dir < 4; dir++ {
			totalVotes[dir] = float64(votes[dir])
			total += totalVotes[dir]
		}

		if total == 0 {
			result = append(result, Pair{Turn: turn, Entropy: 0})
			continue
		}

		entropy := 0.0
		for _, count := range totalVotes {
			if count > 0 {
				p := count / total
				entropy -= p * math.Log2(p)
			}
		}

		maxEntropy := math.Log2(4)
		normalizedEntropy := entropy / maxEntropy
		entropyPercent := int(math.Round(normalizedEntropy * 100))

		result = append(result, Pair{Turn: turn, Entropy: entropyPercent})
	}

	return result
}

func postEntropy(data []Pair) error {
	url := "http://your-api-url.com/api/entropy"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post entropy, status: %d", resp.StatusCode)
	}

	return nil
}