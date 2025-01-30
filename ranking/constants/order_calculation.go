package constants

// OrderCalculation is used in OrderParam.OrderCalculation to control how ties are handled.
type OrderCalculation uint8

const (
	// OrderCalculation113 requests standard "1224" competition ranking.
	OrderCalculation113 OrderCalculation = iota

	// OrderCalculation123 requests strictly ordinal "1234" ranking, with ties broken first by update time
	// (earlier is better) then user PID (lower is better).
	OrderCalculation123
)
