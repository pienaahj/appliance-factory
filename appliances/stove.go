package appliances

// Stove is a Stove struct type
type Stove struct {
	typeName string
}

// Implement the Start fuction -- Appliance interface
func (fr *Stove) Start() {
	fr.typeName = " Stove "
}

// Implement the GetPurpose fuction -- Appliance interface
func (fr *Stove) GetPurpose() string {
	return "I am a " + fr.typeName + " I cook food!!"
}