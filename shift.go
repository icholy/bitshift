package bitshift

// Left performs a left bit shift operation on the provided bytes.
func Left(data []byte, bits int) {
	for i := 0; i < len(data)-1; i++ {
		data[i] = data[i]<<bits | data[i+1]>>(8-bits)
	}
	data[len(data)-1] <<= bits
}

// Right performs a right bit shift operation on the provided bytes.
func Right(data []byte, bits int) {
	for i := len(data) - 1; i > 0; i-- {
		data[i] = data[i]>>bits | data[i-1]<<(8-bits)
	}
	data[0] >>= bits
}
