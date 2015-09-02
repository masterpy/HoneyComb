package utility

import (
//	"math/rand"
//	"strconv"
//	"time"
	"io"
	"encoding/base64"
	crand "crypto/rand"
)

func RandomCode32() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
