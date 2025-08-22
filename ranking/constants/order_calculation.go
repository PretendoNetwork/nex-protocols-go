package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// OrderCalculation is used in OrderParam.OrderCalculation to control how ties are handled.
type OrderCalculation uint8

const (
	// OrderCalculation113 requests standard "1224" competition ranking.
	OrderCalculation113 OrderCalculation = iota

	// OrderCalculation123 requests strictly ordinal "1234" ranking, with ties broken first by update time
	// (earlier is better) then user PID (lower is better).
	OrderCalculation123
)

// WriteTo writes the OrderCalculation to the given writable
func (oc OrderCalculation) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(oc))
}

// ExtractFrom extracts the OrderCalculation value from the given readable
func (oc *OrderCalculation) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*oc = OrderCalculation(value)
	return nil
}
