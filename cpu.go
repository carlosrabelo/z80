package z80

func (cpu *CPU) addU8(a, b uint8) uint8 {
	a16 := uint16(a)
	b16 := uint16(b)
	r08 := uint8(a16 + b16)
	return r08
}
