package ft

type Color string

//Colors for terminal
const (
	Clr0        = "\x1b[30;1m"
	BLACK Color = "\x1b[30;1m"

	ClrR       = "\x1b[31;1m"
	RED  Color = "\x1b[31;1m"

	ClrG        = "\x1b[32;1m"
	GREEN Color = "\x1b[32;1m"

	ClrY         = "\x1b[33;1m"
	YELLOW Color = "\x1b[33;1m"

	ClrB = "\x1b[34;1m"
	BLUE = "\x1b[34;1m"
	
	ClrM          = "\x1b[35;1m"
	MARGENT Color = "\x1b[35;1m"

	ClrC       = "\x1b[36;1m"
	CYAN Color = "\x1b[36;1m"

	ClrW        = "\x1b[37;1m"
	WHITE Color = "\x1b[37;1m"

	ClrN = "\x1b[0m"
	NONE Color = "\x1b[0m"
)
