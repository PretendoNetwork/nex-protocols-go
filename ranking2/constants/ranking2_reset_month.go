package constants

import (
	"time"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2ResetMonth determines on what months are enabled for for category reset
// modes `Ranking2ResetModeMultiMonth` and `Ranking2ResetModeMultiMonthWeekday`.
// The value is a 12-bit flag set, where each bit represents a month of the year.
// If a bit is set, rankings reset in that month. Bits start at January as the LSB
// and go up to December in the standard calendar order.
//
// Note: This is not a real type. This is a bespoke type made for our convenience.
type Ranking2ResetMonth uint16

// WriteTo writes the Ranking2ResetMonth to the given writable
func (r2rm Ranking2ResetMonth) WriteTo(writable types.Writable) {
	writable.WriteUInt16LE(uint16(r2rm))
}

// ExtractFrom extracts the Ranking2ResetMonth value from the given readable
func (r2rm *Ranking2ResetMonth) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt16LE()
	if err != nil {
		return err
	}

	*r2rm = Ranking2ResetMonth(value)
	return nil
}

// IsEnabled checks if a specific month is enabled
func (r2rm Ranking2ResetMonth) IsEnabled(month Ranking2ResetMonth) bool {
	return r2rm&month != 0
}

// Enable adds a month to the enabled set
func (r2rm Ranking2ResetMonth) Enable(month Ranking2ResetMonth) Ranking2ResetMonth {
	return r2rm | month
}

// Disable removes a month from the enabled set
func (r2rm Ranking2ResetMonth) Disable(month Ranking2ResetMonth) Ranking2ResetMonth {
	return r2rm &^ month
}

// Toggle switches the state of a month
func (r2rm Ranking2ResetMonth) Toggle(month Ranking2ResetMonth) Ranking2ResetMonth {
	return r2rm ^ month
}

// IsEnabledForTime checks if the given time.Time's month is enabled
func (r2rm Ranking2ResetMonth) IsEnabledForTime(t time.Time) bool {
	monthFlag := Ranking2ResetMonth(1 << (t.Month() - 1))
	return r2rm.IsEnabled(monthFlag)
}

// EnabledMonths returns a slice of time.Month values that are enabled
func (r2rm Ranking2ResetMonth) EnabledMonths() []time.Month {
	var months []time.Month
	for i := 0; i < 12; i++ {
		if r2rm&(1<<i) != 0 {
			months = append(months, time.Month(i+1))
		}
	}
	return months
}

// Count returns the number of enabled months
func (r2rm Ranking2ResetMonth) Count() int {
	return len(r2rm.EnabledMonths())
}

const (
	// Ranking2ResetMonthJanuary means that January is enabled for ranking resets.
	Ranking2ResetMonthJanuary Ranking2ResetMonth = 0x001

	// Ranking2ResetMonthFebruary means that February is enabled for ranking resets.
	Ranking2ResetMonthFebruary Ranking2ResetMonth = 0x002

	// Ranking2ResetMonthMarch means that March is enabled for ranking resets.
	Ranking2ResetMonthMarch Ranking2ResetMonth = 0x004

	// Ranking2ResetMonthApril means that April is enabled for ranking resets.
	Ranking2ResetMonthApril Ranking2ResetMonth = 0x008

	// Ranking2ResetMonthMay means that May is enabled for ranking resets.
	Ranking2ResetMonthMay Ranking2ResetMonth = 0x010

	// Ranking2ResetMonthJune means that June is enabled for ranking resets.
	Ranking2ResetMonthJune Ranking2ResetMonth = 0x020

	// Ranking2ResetMonthJuly means that July is enabled for ranking resets.
	Ranking2ResetMonthJuly Ranking2ResetMonth = 0x040

	// Ranking2ResetMonthAugust means that August is enabled for ranking resets.
	Ranking2ResetMonthAugust Ranking2ResetMonth = 0x080

	// Ranking2ResetMonthSeptember means that September is enabled for ranking resets.
	Ranking2ResetMonthSeptember Ranking2ResetMonth = 0x100

	// Ranking2ResetMonthOctober means that October is enabled for ranking resets.
	Ranking2ResetMonthOctober Ranking2ResetMonth = 0x200

	// Ranking2ResetMonthNovember means that November is enabled for ranking resets.
	Ranking2ResetMonthNovember Ranking2ResetMonth = 0x400

	// Ranking2ResetMonthDecember means that December is enabled for ranking resets.
	Ranking2ResetMonthDecember Ranking2ResetMonth = 0x800
)

// * Some common combinations we might find useful
const (
	// Ranking2ResetMonthNone disables all months.
	Ranking2ResetMonthNone Ranking2ResetMonth = 0

	// Ranking2ResetMonthAll enables all months.
	Ranking2ResetMonthAll Ranking2ResetMonth = 0x0FFF

	// Ranking2ResetMonthQuarterly resets rankings 4 times a year.
	Ranking2ResetMonthQuarterly Ranking2ResetMonth = Ranking2ResetMonthJanuary | Ranking2ResetMonthApril | Ranking2ResetMonthJuly | Ranking2ResetMonthOctober

	// Ranking2ResetMonthSemiannual resets rankings twice a year.
	Ranking2ResetMonthSemiannual Ranking2ResetMonth = Ranking2ResetMonthJanuary | Ranking2ResetMonthJuly
)
