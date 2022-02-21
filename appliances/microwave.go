package appliances

// Microwave is a Microwave struct type
type Microwave struct {
	typeName string
}

// Implement the Start fuction -- Appliance interface
func (fr *Microwave) Start() {
	fr.typeName = " Microwave "
}

// Implement the GetPurpose fuction -- Appliance interface
func (fr *Microwave) GetPurpose() string {
	return "I am a " + fr.typeName + " I heat stuff up!!"
}