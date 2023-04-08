package utility

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AssociateNexUniqueIDWithMyPrincipalID sets the AssociateNexUniqueIDWithMyPrincipalID handler function
func (protocol *UtilityProtocol) AssociateNexUniqueIDWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo *UniqueIDInfo)) {
	protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) HandleAssociateNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueIDInfoStructureInterface, err := parametersStream.ReadStructure(NewUniqueIDInfo())
	if err != nil {
		go protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, nil)
		return
	}
	uniqueIDInfo := uniqueIDInfoStructureInterface.(*UniqueIDInfo)

	go protocol.AssociateNexUniqueIDWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo)
}
