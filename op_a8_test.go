package z80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOP_ADD_A_R(t *testing.T) {
	cpu := &CPU{}

	cpu.AF.Hi = 0x00
	cpu.BC.Hi = 0x01
	cpu.BC.Lo = 0x02
	cpu.DE.Hi = 0x03
	cpu.DE.Lo = 0x04
	cpu.HL.Hi = 0x05
	cpu.HL.Lo = 0x06

	OP_ADD_A_B(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x01), "The values should be the same")
	OP_ADD_A_C(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x03), "The values should be the same")
	OP_ADD_A_D(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x06), "The values should be the same")
	OP_ADD_A_E(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x0A), "The values should be the same")
	OP_ADD_A_H(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x0F), "The values should be the same")
	OP_ADD_A_L(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x15), "The values should be the same")
	OP_ADD_A_A(cpu)
	assert.Equal(t, cpu.AF.Hi, uint8(0x2A), "The values should be the same")
}
