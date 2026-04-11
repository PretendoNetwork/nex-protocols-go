package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

// ResultCode is a custom result value used by legacy Ranking representing the kind of operation performed by a method.
//
// Note: The names of this type and its values are guesses based on context
type ResultCode int16

// WriteTo writes the ResultCode to the given writable
func (rc ResultCode) WriteTo(writable types.Writable) {
	writable.WriteInt16LE(int16(rc))
}

// ExtractFrom extracts the ResultCode value from the given readable
func (rc *ResultCode) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadInt16LE()
	if err != nil {
		return err
	}

	*rc = ResultCode(value)
	return nil
}

const (
	// ResultCodeGetOne indicates an operation that gets individual rows from a database
	ResultCodeGetOne ResultCode = 20

	// ResultCodeGetMany indicates an operation that gets multiple rows from a database, including a row count
	ResultCodeGetMany = 30

	// ResultCodeUpdateOne indicates an operation that inserts, updates or deletes individual rows on a database
	ResultCodeUpdateOne = 50

	// ResultCodeUpdateMany indicates an operation that inserts, updates or deletes multiple rows on a database
	ResultCodeUpdateMany = 100
)
