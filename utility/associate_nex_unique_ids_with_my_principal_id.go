// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

func (protocol *Protocol) handleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetAssociatedNexUniqueIDWithMyPrincipalID == nil {
		globals.Logger.Warning("Utility::AssociateNexUniqueIDsWithMyPrincipalID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)
	uniqueIDInfo, err := nex.StreamReadListStructure(parametersStream, utility_types.NewUniqueIDInfo())
	if err != nil {
		_, errorCode = protocol.AssociateNexUniqueIDsWithMyPrincipalID(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AssociateNexUniqueIDsWithMyPrincipalID(nil, packet, callID, uniqueIDInfo)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
