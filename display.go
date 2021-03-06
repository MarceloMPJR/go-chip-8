package chip8

type Display interface {
	/*
		Clear should clear the display
	*/
	Clear()

	/*
		Draw should draw the byte on position xDisplay and yDisplay
		If occourres collision returns true
	*/
	Draw(xDisplay, yDisplay, sprite byte) bool

	/*
		Flush should write on output of display
	*/
	Flush()
}
