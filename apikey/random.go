package apikey

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

func secureRandomString(n int) string {
	buf := make([]byte, 128)
	_, err := rand.Read(buf)
	if err != nil {
		logrus.Panic(err)
	}
	return base64.StdEncoding.EncodeToString(buf)
}
