package client

type IoPin struct {
	Alias    string // line alias such as pump, fan, etc
	Line     int    // gpio line struct
	GpioChip string // gpio chip
	Value    int    // current value of the pin.
	AsOutput bool   // 1 is output, 0 is input
}
