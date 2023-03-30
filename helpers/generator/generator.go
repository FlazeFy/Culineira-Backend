package generator

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

func GetToken(sample string) (string, error) {
	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = sample
	claims["expired"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
