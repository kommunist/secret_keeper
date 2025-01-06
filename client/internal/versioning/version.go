package versioning

import (
	"fmt"
	"time"
)

type Version struct{}

func (v Version) Get() string {
	return fmt.Sprintf("%v", time.Now().Unix())
}
