package main

import (
	"strings"
)

func normalizeDSN(dsn string) (normalized string, err error) {
	normalized = strings.Replace(dsn, "@", "@tcp(", -1)
	normalized = strings.Replace(normalized, "3306", "3306)", -1)
	normalized = strings.Replace(normalized, "/app", "/app?charset=utf8&parseTime=True", -1)
	normalized = strings.Replace(normalized, "mysql://", "", -1)

	return normalized, nil
}
