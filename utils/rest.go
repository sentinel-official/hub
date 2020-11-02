package utils

import (
	"net/url"
	"strconv"
)

func ParseQuery(query url.Values) (skip, limit int, err error) {
	skip, limit = 1, 25

	if query.Get("skip") != "" {
		skip, err = strconv.Atoi(query.Get("skip"))
		if err != nil {
			return 0, 0, err
		}
	}

	if query.Get("limit") != "" {
		limit, err = strconv.Atoi(query.Get("limit"))
		if err != nil {
			return 0, 0, err
		}
	}

	return skip, limit, nil
}
