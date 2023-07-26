// Package utility implements the Utility NEX protocol
package utility

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

// AssociateNexUniqueIDsWithMyPrincipalID sets the AssociateNexUniqueIDsWithMyPrincipalID handler function
func (protocol *UtilityProtocol) AssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, client *nex.Client, callID uint32, uniqueIDInfo []*utility_types.UniqueIDInfo)) {
	protocol.associateNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *UtilityProtocol) handleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	uniqueIDInfo, err := parametersStream.ReadListStructure(utility_types.NewUniqueIDInfo())
	if err != nil {
		go protocol.associateNexUniqueIDsWithMyPrincipalIDHandler(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.associateNexUniqueIDsWithMyPrincipalIDHandler(nil, client, callID, uniqueIDInfo.([]*utility_types.UniqueIDInfo))
}
