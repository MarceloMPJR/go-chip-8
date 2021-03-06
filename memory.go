package chip8

type Memory interface {
	/*
		Save should saves the register on the memory starting from I on memory
	*/
	Save(register []byte, i uint16)

	/*
		SaveBCD should convert vx to decimal and saves each decimal on addresses I, I+1, I+2
	*/
	SaveBCD(vx byte, i uint16)

	/*
		Load should loads the register on the memory starting from I on memory
	*/
	Load(register []byte, i uint16)

	/*
		LoadInstruction should loads the instruction (2 bytes) of program
	*/
	LoadInstruction(pc uint16) Instruction

	/*
		LoadChar should return the address of char referring to vx
	*/
	LoadChar(vx byte) uint16

	/*
		LoadSprite should return the Sprite on I
	*/
	LoadSprite(i uint16) byte
}
