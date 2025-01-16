package versioning

import (
	"fmt"
	"time"
)

type Version struct{}

// Пакет, генерирующий номер версии из времени. Версия всегда растет
func (v Version) Get() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}
