package generator

import (
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
