package generator

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetDateTimeNow() string {
	now := time.Now()
	res := now.Format("2006-01-02 15:04:05")

	return res
}

func GetUUID() string {
	id := uuid.New()
	res := id.String()

	return res
}

func GetSlug(sample string) string {
	slug := strings.ToLower(strings.ReplaceAll(sample, " ", "-"))
	rege := regexp.MustCompile("[^a-z0-9-]+")
	res := rege.ReplaceAllString(slug, "")

	return res
}
