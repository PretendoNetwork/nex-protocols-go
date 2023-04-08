package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AssociateNexUniqueIDsWithMyPrincipalID sets the AssociateNexUniqueIDsWithMyPrincipalID handler function
func (protocol *UtilityProtocol) AssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*UniqueIDInfo)) {
	protocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) HandleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.GetAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	structureCount := (int)(parametersStream.ReadUInt32LE())
	uniqueIDInfo := make([]*UniqueIDInfo, structureCount)

	for i := 0; i < structureCount; i++ {
		uniqueIDInfoStructureInterface, err := parametersStream.ReadStructure(NewUniqueIDInfo())
		if err != nil {
			go protocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, nil)
			return
		}
		uniqueIDInfo[i] = uniqueIDInfoStructureInterface.(*UniqueIDInfo)
	}

	go protocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo)
}
