package unique

import (
	"fmt"
	"sync/atomic"
)

var globalID = uint64(0)

func GenUniqueKey(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, genUniqueID())
}

func genUniqueID() uint64 {
	return atomic.AddUint64(&globalID, 1)
}
