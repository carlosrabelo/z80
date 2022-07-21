package z80

type CPU struct {
	GPRc
	GPRa
	SPR
}

type GPRc struct {
	AF Register
	BC Register
	DE Register
	HL Register
}

type GPRa struct {
	AF_ Register
	BC_ Register
	DE_ Register
	HL_ Register
}

type SPR struct {
	IR Register

	IX uint16
	IY uint16
	SP uint16
	PC uint16
}

type Register struct {
	Hi uint8
	Lo uint8
}
