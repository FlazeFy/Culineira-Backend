package generator

import (
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetDateTimeNow() string {
	now := time.Now()
	res := now.Format(os.Getenv("DATETIME_FORMAT"))

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
