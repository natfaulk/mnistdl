package mnistdl

import (
	"path"
	"strings"
)

func trimExtension(name string) string {
	return strings.TrimSuffix(name, path.Ext(name))
}

func getUint32(data []uint8, index int) uint32 {
	out := uint32(data[index+0]) << (3 * 8)
	out |= uint32(data[index+1]) << (2 * 8)
	out |= uint32(data[index+2]) << (1 * 8)
	out |= uint32(data[index+3]) << (0 * 8)
	return out
}
