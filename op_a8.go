package z80

// 8-Bit Arithmetic

// ADD A,r

func OP_ADD_A_B(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.BC.Hi
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_C(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.BC.Lo
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_D(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.DE.Hi
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_E(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.DE.Lo
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_H(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.HL.Hi
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_L(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.HL.Lo
	cpu.AF.Hi = cpu.addU8(a, r)
}

func OP_ADD_A_A(cpu *CPU) {
	a := cpu.AF.Hi
	r := cpu.AF.Hi
	cpu.AF.Hi = cpu.addU8(a, r)
}

// ADD A,n

// ADD A,(hl)

// ADD A,(IX+d)

// ADD A,(IY+d)

// ADC A,s

// SUB s

// SBC A,s

// AND s

// OR s

// XOR s

// CP s

// INC r

// INC (HL)

// INC (IX+d)

// INC (IY+d)

// DEC m
