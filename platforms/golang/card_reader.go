package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

func fetchAllCards() ([]interface{}, error) {
	var result []interface{}
	path := "../../katas/MTGCardReader/allcards.json"

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return result, errors.Wrapf(err, "unable to open file: %s", path)
	}

	// remove byte-order-mark (BOM)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, errors.Wrapf(err, "unable to parse card JSON")
	}

	return result, nil
}
