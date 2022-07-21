package z80

// 16-Bit Arithmetic

// ADD HL,ss

func OP_ADD_HL_BC(cpu *CPU) {
	a := cpu.HL.getU16()
	x := cpu.BC.getU16()
	cpu.HL.setU16(cpu.addU16(a, x))
}

func OP_ADD_HL_DE(cpu *CPU) {
	a := cpu.HL.getU16()
	x := cpu.DE.getU16()
	cpu.HL.setU16(cpu.addU16(a, x))
}

func OP_ADD_HL_HL(cpu *CPU) {
	a := cpu.HL.getU16()
	x := cpu.HL.getU16()
	cpu.HL.setU16(cpu.addU16(a, x))
}

func OP_ADD_HL_SP(cpu *CPU) {
	a := cpu.HL.getU16()
	x := cpu.SP
	cpu.HL.setU16(cpu.addU16(a, x))
}

// ADC HL,ss

// SBC HL,ss

// ADD IX,pp

// ADD IY,rr

// INC ss

// INC IX

// INC IY

// DEC ss

// DEC IX

// DEC IY
