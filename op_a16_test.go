package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOP_ADD_HL_SS(t *testing.T) {
	cpu := &CPU{}

	cpu.HL.setU16(0x0000)
	cpu.BC.setU16(0x0001)
	cpu.DE.setU16(0x0002)

	OP_ADD_HL_BC(cpu)
	assert.Equal(t, cpu.HL.getU16(), uint16(0x0001), "The values should be the same")
	OP_ADD_HL_DE(cpu)
	assert.Equal(t, cpu.HL.getU16(), uint16(0x0003), "The values should be the same")
	OP_ADD_HL_HL(cpu)
	assert.Equal(t, cpu.HL.getU16(), uint16(0x0006), "The values should be the same")

	cpu.SP = 0x1000

	OP_ADD_HL_SP(cpu)
	assert.Equal(t, cpu.HL.getU16(), uint16(0x1006), "The values should be the same")
}
