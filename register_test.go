package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSet(t *testing.T) {
	cpu := &CPU{}

	cpu.HL.setU16(0x1234)
	assert.Equal(t, cpu.HL.getU16(), uint16(0x1234), "The values should be the same")
}
