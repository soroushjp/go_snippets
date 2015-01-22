func main() {
	myColor, _ := hex.DecodeString("ABFD55")     //Example our color
	myRedMask, _ := []byte(3 << 2) 
	myGreenMask, _ := hex.DecodeString("00FF00") //Our green mask
	myBlueMask, _ := hex.DecodeString("0000FF")  //Our blue mask
	myClearMask, _ := hex.DecodeString("000000") //Our clear mask. Not very useful in this particular case, but could be

	redOnly := AndMask(myColor, myRedMask)
	greenOnly := AndMask(myColor, myGreenMask)
	blueOnly := AndMask(myColor, myBlueMask)
	empty := AndMask(myColor, myClearMask)

	fmt.Println(redOnly, greenOnly, blueOnly, empty)

	//Flip all colours by Xor mask of FFFFFF
	myFlipMask, _ := hex.DecodeString("FFFFFF")
	fmt.Println(XorMask(myColor, myFlipMask))
}

func AndMask(data []byte, mask []byte) []byte {
	if len(data) != len(mask) {
		panic("Mask and data must have identical lengths")
	}

	maskedData := make([]byte, len(data))

	for i := range data {
		maskedData[i] = data[i] & mask[i]
	}

	return maskedData
}

func XorMask(data []byte, mask []byte) []byte {
	if len(data) != len(mask) {
		panic("Mask and data must have identical lengths")
	}

	maskedData := make([]byte, len(data))

	for i := range data {
		maskedData[i] = data[i] ^ mask[i]
	}

	return maskedData
}
