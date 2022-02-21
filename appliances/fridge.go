package appliances

// Fridge is a Fridge struct type
type Fridge struct {
	typeName string
}

// Implement the Start fuction -- Appliance interface
func (fr *Fridge) Start() {
	fr.typeName = " Fridge "
}

// Implement the GetPurpose fuction -- Appliance interface
func (fr *Fridge) GetPurpose() string {
	return "I am a " + fr.typeName + " I cool stuff down!!"
}