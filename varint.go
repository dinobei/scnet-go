package scnet

func encodeVarint(value int32) []byte {
	output := []byte{}

	for value > 127 {
		output = append(output, uint8(value&127)|128)
		value >>= 7
	}
	output = append(output, uint8(value&127))
	return output
}

func decodeVarint(input []uint8) int32 {
	var ret int32
	for i, value := range input {
		ret |= (int32(value) & 127) << uint32(7*i)
		if value&128 == 0 {
			break
		}
	}

	return ret
}
