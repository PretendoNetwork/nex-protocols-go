// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemRightInfos holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightInfos struct {
	types.Structure
	SupportID                       string
	ConsumptionRightInfos           []*ServiceItemRightConsumptionInfo
	AdditionalTimeRightInfos        []*ServiceItemRightTimeInfo
	PermanentRightInfos             []*ServiceItemRightTimeInfo
	AlreadyPurchasedInitialOnlyItem *types.PrimitiveBool
}

// ExtractFrom extracts the ServiceItemRightInfos from the given readable
func (serviceItemRightInfos *ServiceItemRightInfos) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemRightInfos.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemRightInfos header. %s", err.Error())
	}

	err = serviceItemRightInfos.SupportID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.SupportID from stream. %s", err.Error())
	}

	consumptionRightInfos, err := nex.StreamReadListStructure(stream, NewServiceItemRightConsumptionInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.ConsumptionRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.ConsumptionRightInfos = consumptionRightInfos

	additionalTimeRightInfos, err := nex.StreamReadListStructure(stream, NewServiceItemRightTimeInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AdditionalTimeRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.AdditionalTimeRightInfos = additionalTimeRightInfos

	permanentRightInfos, err := nex.StreamReadListStructure(stream, NewServiceItemRightTimeInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.PermanentRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.PermanentRightInfos = permanentRightInfos

	err = serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AlreadyPurchasedInitialOnlyItem from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemRightInfos to the given writable
func (serviceItemRightInfos *ServiceItemRightInfos) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemRightInfos.SupportID.WriteTo(contentWritable)
	serviceItemRightInfos.ConsumptionRightInfos.WriteTo(contentWritable)
	serviceItemRightInfos.AdditionalTimeRightInfos.WriteTo(contentWritable)
	serviceItemRightInfos.PermanentRightInfos.WriteTo(contentWritable)
	serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemRightInfos.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (serviceItemRightInfos *ServiceItemRightInfos) Copy() types.RVType {
	copied := NewServiceItemRightInfos()

	copied.StructureVersion = serviceItemRightInfos.StructureVersion

	copied.SupportID = serviceItemRightInfos.SupportID
	copied.ConsumptionRightInfos = make([]*ServiceItemRightConsumptionInfo, len(serviceItemRightInfos.ConsumptionRightInfos))

	for i := 0; i < len(serviceItemRightInfos.ConsumptionRightInfos); i++ {
		copied.ConsumptionRightInfos[i] = serviceItemRightInfos.ConsumptionRightInfos[i].Copy().(*ServiceItemRightConsumptionInfo)
	}

	copied.AdditionalTimeRightInfos = make([]*ServiceItemRightTimeInfo, len(serviceItemRightInfos.AdditionalTimeRightInfos))

	for i := 0; i < len(serviceItemRightInfos.AdditionalTimeRightInfos); i++ {
		copied.AdditionalTimeRightInfos[i] = serviceItemRightInfos.AdditionalTimeRightInfos[i].Copy().(*ServiceItemRightTimeInfo)
	}

	copied.PermanentRightInfos = make([]*ServiceItemRightTimeInfo, len(serviceItemRightInfos.PermanentRightInfos))

	for i := 0; i < len(serviceItemRightInfos.PermanentRightInfos); i++ {
		copied.PermanentRightInfos[i] = serviceItemRightInfos.PermanentRightInfos[i].Copy().(*ServiceItemRightTimeInfo)
	}

	copied.AlreadyPurchasedInitialOnlyItem = serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemRightInfos *ServiceItemRightInfos) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemRightInfos); !ok {
		return false
	}

	other := o.(*ServiceItemRightInfos)

	if serviceItemRightInfos.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemRightInfos.SupportID.Equals(other.SupportID) {
		return false
	}

	if len(serviceItemRightInfos.ConsumptionRightInfos) != len(other.ConsumptionRightInfos) {
		return false
	}

	for i := 0; i < len(serviceItemRightInfos.ConsumptionRightInfos); i++ {
		if !serviceItemRightInfos.ConsumptionRightInfos[i].Equals(other.ConsumptionRightInfos[i]) {
			return false
		}
	}

	if len(serviceItemRightInfos.AdditionalTimeRightInfos) != len(other.AdditionalTimeRightInfos) {
		return false
	}

	for i := 0; i < len(serviceItemRightInfos.AdditionalTimeRightInfos); i++ {
		if !serviceItemRightInfos.AdditionalTimeRightInfos[i].Equals(other.AdditionalTimeRightInfos[i]) {
			return false
		}
	}

	if len(serviceItemRightInfos.PermanentRightInfos) != len(other.PermanentRightInfos) {
		return false
	}

	for i := 0; i < len(serviceItemRightInfos.PermanentRightInfos); i++ {
		if !serviceItemRightInfos.PermanentRightInfos[i].Equals(other.PermanentRightInfos[i]) {
			return false
		}
	}

	return serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem == other.AlreadyPurchasedInitialOnlyItem
}

// String returns a string representation of the struct
func (serviceItemRightInfos *ServiceItemRightInfos) String() string {
	return serviceItemRightInfos.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemRightInfos *ServiceItemRightInfos) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRightInfos{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemRightInfos.StructureVersion))
	b.WriteString(fmt.Sprintf("%sSupportID: %q,\n", indentationValues, serviceItemRightInfos.SupportID))

	if len(serviceItemRightInfos.ConsumptionRightInfos) == 0 {
		b.WriteString(fmt.Sprintf("%sConsumptionRightInfos: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sConsumptionRightInfos: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightInfos.ConsumptionRightInfos); i++ {
			str := serviceItemRightInfos.ConsumptionRightInfos[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightInfos.ConsumptionRightInfos)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if len(serviceItemRightInfos.AdditionalTimeRightInfos) == 0 {
		b.WriteString(fmt.Sprintf("%sAdditionalTimeRightInfos: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sAdditionalTimeRightInfos: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightInfos.AdditionalTimeRightInfos); i++ {
			str := serviceItemRightInfos.AdditionalTimeRightInfos[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightInfos.AdditionalTimeRightInfos)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	if len(serviceItemRightInfos.PermanentRightInfos) == 0 {
		b.WriteString(fmt.Sprintf("%sPermanentRightInfos: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sPermanentRightInfos: [\n", indentationValues))

		for i := 0; i < len(serviceItemRightInfos.PermanentRightInfos); i++ {
			str := serviceItemRightInfos.PermanentRightInfos[i].FormatToString(indentationLevel + 2)
			if i == len(serviceItemRightInfos.PermanentRightInfos)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sAlreadyPurchasedInitialOnlyItem: %t,\n", indentationValues, serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRightInfos returns a new ServiceItemRightInfos
func NewServiceItemRightInfos() *ServiceItemRightInfos {
	return &ServiceItemRightInfos{}
}
