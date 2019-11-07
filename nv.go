package nv

type Response struct {
	ErrorCode    []byte
	ErrorMessage string
	DataLen      int
	Data         []byte
}

func crc16(data []byte) []byte {
	seed := uint16(0xFFFF)
	poly := uint16(0x8005)
	crc := seed

	for _, d := range data {
		crc ^= uint16(d) << 8
		for i := 0; i < 8; i++ {
			bit := (crc & 0x8000) != 0
			crc <<= 1
			if bit {
				crc ^= poly
			}
		}
	}

	b := [2]byte{
		byte(crc & 0xFF),
		byte((crc >> 8) & 0xFF),
	}

	return b[:]
}
