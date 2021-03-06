package appliances

import "errors"

// import errors to log errors

// Appliance the main interface used to describe appliances
type Appliance interface {
	Start()
	GetPurpose() string
}

// Our appliance types
const (
	STOVE =iota
	FRIDGE
	MICROWAVE
)

// CreateAppliance creates an appliance
func CreateAppliance(myType int) (Appliance, error) {
	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil
	default:
		return nil, errors.New("invalid Appliance Type")
	}
}