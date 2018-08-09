package uid

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
)

func NewId() string {
	UUID, _ := uuid.NewV4()
	b64 := base64.URLEncoding.EncodeToString(UUID.Bytes()[:12])
	return b64
}
