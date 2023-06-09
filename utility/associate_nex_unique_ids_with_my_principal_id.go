package utility

import (
	"fmt"

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
	uniqueIDInfo, err := parametersStream.ReadListStructure(NewUniqueIDInfo())
	if err != nil {
		go protocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.AssociateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo.([]*UniqueIDInfo))
}
