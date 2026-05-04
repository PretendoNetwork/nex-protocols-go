package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

type ReportCategory uint32

// WriteTo writes the ReportCategory to the given writable
func (rc ReportCategory) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(rc))
}

// ExtractFrom extracts the ReportCategory value from the given readable
func (rc *ReportCategory) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*rc = ReportCategory(value)
	return nil
}

// String returns a human-readable representation of the ReportCategory.
func (rc ReportCategory) String() string {
	switch rc {
	case ReportCategoryInvalid:
		return "Invalid"
	case ReportCategoryPersonal:
		return "Personal"
	case ReportCategoryCriminal:
		return "Criminal"
	case ReportCategoryImmoral:
		return "Immoral"
	case ReportCategoryHarassment:
		return "Harassment"
	case ReportCategoryCommercial:
		return "Commercial"
	case ReportCategorySexuallyExplicit:
		return "SexuallyExplicit"
	case ReportCategoryOther:
		return "Other"
	default:
		return fmt.Sprintf("ReportCategory(%d)", int(rc))
	}
}

const (
	// ReportCategoryInvalid represents an invalid category
	ReportCategoryInvalid ReportCategory = iota

	// ReportCategoryPersonal means the reported content contained personal information
	ReportCategoryPersonal

	// ReportCategoryCriminal means the reported content contained criminal material
	ReportCategoryCriminal

	// ReportCategoryImmoral means the reported content contained immoral material
	ReportCategoryImmoral

	// ReportCategoryHarassment means the reported content contained harassment material
	ReportCategoryHarassment

	// ReportCategoryCommercial means the reported content contained commercial material
	ReportCategoryCommercial

	// ReportCategorySexuallyExplicit means the reported content contained sexually explicit material
	ReportCategorySexuallyExplicit

	// ReportCategoryOther means the reported content didn't fit into the categories
	ReportCategoryOther
)
