// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

// AssociateNexUniqueIDsWithMyPrincipalID sets the AssociateNexUniqueIDsWithMyPrincipalID handler function
func (protocol *Protocol) AssociateNexUniqueIDsWithMyPrincipalID(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueIDInfo []*utility_types.UniqueIDInfo) uint32) {
	protocol.associateNexUniqueIDsWithMyPrincipalIDHandler = handler
}

func (protocol *Protocol) handleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAssociatedNexUniqueIDWithMyPrincipalIDHandler == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDsWithMyPrincipalID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()

	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	uniqueIDInfo, err := parametersStream.ReadListStructure(utility_types.NewUniqueIDInfo())
	if err != nil {
		errorCode = protocol.associateNexUniqueIDsWithMyPrincipalIDHandler(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.associateNexUniqueIDsWithMyPrincipalIDHandler(nil, packet, callID, uniqueIDInfo.([]*utility_types.UniqueIDInfo))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
