package util

import (
	"encoding/json"
	"os"
)

func ReadAngramFile(alphabet rune) (words []string, err error) {
	data, err := os.ReadFile("internal/data/words/" + string(alphabet) + ".json")
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &words); err != nil {
		return nil, err
	}

	return words, nil
}
