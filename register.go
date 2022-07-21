package z80

func (r Register) getU16() uint16 {
	v := (uint16(r.Hi) << 8) | uint16(r.Lo)
	return v
}

func (r *Register) setU16(v uint16) {
	r.Hi = uint8(v >> 8)
	r.Lo = uint8(v & 0x00ff)
}
