package utils

import (
	"net/url"
	"strconv"
)

func ParseQuery(query url.Values) (page, limit int, err error) {
	page, limit = 1, 0

	if query.Get("page") != "" {
		page, err = strconv.Atoi(query.Get("page"))
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

	return page, limit, nil
}
