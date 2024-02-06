// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

func (protocol *Protocol) handleAssociateNexUniqueIDWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.AssociateNexUniqueIDWithMyPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Utility::AssociateNexUniqueIDWithMyPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uniqueIDInfo := utility_types.NewUniqueIDInfo()

	err := uniqueIDInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AssociateNexUniqueIDWithMyPrincipalID(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AssociateNexUniqueIDWithMyPrincipalID(nil, packet, callID, uniqueIDInfo)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
