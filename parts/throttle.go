package parts

// Throttle interface is the generic interface type that should be passed around
type Throttle interface {
}

// throttle is the basic implementation of throttle object
type throttle struct {
}

// NewThrottle is the constructor function for throttle control interface
func NewThrottle() Throttle {
	return throttle{}
}
