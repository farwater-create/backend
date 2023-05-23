package apikey

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/sirupsen/logrus"
)

func secureRandomString(n int) string {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	if err != nil {
		logrus.Panic(err)
	}
	randomString := strings.TrimRight(base64.URLEncoding.EncodeToString(buf), "=")
	return randomString[:n]
}
