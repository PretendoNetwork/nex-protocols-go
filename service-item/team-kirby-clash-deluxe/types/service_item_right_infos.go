// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemRightInfos holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemRightInfos struct {
	nex.Structure
	SupportID                       string
	ConsumptionRightInfos           []*ServiceItemRightConsumptionInfo
	AdditionalTimeRightInfos        []*ServiceItemRightTimeInfo
	PermanentRightInfos             []*ServiceItemRightTimeInfo
	AlreadyPurchasedInitialOnlyItem bool
}

// ExtractFromStream extracts a ServiceItemRightInfos structure from a stream
func (serviceItemRightInfos *ServiceItemRightInfos) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemRightInfos.SupportID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.SupportID from stream. %s", err.Error())
	}

	consumptionRightInfos, err := stream.ReadListStructure(NewServiceItemRightConsumptionInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.ConsumptionRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.ConsumptionRightInfos = consumptionRightInfos.([]*ServiceItemRightConsumptionInfo)

	additionalTimeRightInfos, err := stream.ReadListStructure(NewServiceItemRightTimeInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AdditionalTimeRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.AdditionalTimeRightInfos = additionalTimeRightInfos.([]*ServiceItemRightTimeInfo)

	permanentRightInfos, err := stream.ReadListStructure(NewServiceItemRightTimeInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.PermanentRightInfos from stream. %s", err.Error())
	}

	serviceItemRightInfos.PermanentRightInfos = permanentRightInfos.([]*ServiceItemRightTimeInfo)

	serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRightInfos.AlreadyPurchasedInitialOnlyItem from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemRightInfos and returns a byte array
func (serviceItemRightInfos *ServiceItemRightInfos) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemRightInfos.SupportID)
	stream.WriteListStructure(serviceItemRightInfos.ConsumptionRightInfos)
	stream.WriteListStructure(serviceItemRightInfos.AdditionalTimeRightInfos)
	stream.WriteListStructure(serviceItemRightInfos.PermanentRightInfos)
	stream.WriteBool(serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemRightInfos
func (serviceItemRightInfos *ServiceItemRightInfos) Copy() nex.StructureInterface {
	copied := NewServiceItemRightInfos()

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
func (serviceItemRightInfos *ServiceItemRightInfos) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemRightInfos)

	if serviceItemRightInfos.SupportID != other.SupportID {
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

	return serviceItemRightInfos.AlreadyPurchasedInitialOnlyItem != other.AlreadyPurchasedInitialOnlyItem
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemRightInfos.StructureVersion()))
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
