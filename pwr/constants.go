package pwr

import (
	"encoding/binary"

	"github.com/itchio/wharf/sync"
)

var endianness = binary.LittleEndian

const (
	pwrMagic = int32(iota + 0xFEF5F00)
)

var BlockSize = 4 * 1024 // 64k

func mksync() *sync.SyncContext {
	return sync.NewContext(BlockSize)
}
