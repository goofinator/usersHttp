package utils

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// IDFromURL extracts id from endpoints like: /users/:id
func IDFromURL(url *url.URL) (int, error) {
	uri := url.String()
	re := regexp.MustCompile(`/:[0-9]+$`)
	uri = re.FindString(uri)
	uri = strings.TrimLeft(uri, "/:")

	return strconv.Atoi(uri)
}
