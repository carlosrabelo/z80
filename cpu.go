package z80

func (cpu *CPU) addU8(a, b uint8) uint8 {
	a16 := uint16(a)
	b16 := uint16(b)
	r08 := uint8(a16 + b16)
	return r08
}

func (cpu *CPU) addU16(a, b uint16) uint16 {
	a32 := uint32(a)
	b32 := uint32(b)
	r16 := uint16(a32 + b32)
	return r16
}
