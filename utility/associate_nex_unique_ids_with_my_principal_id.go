// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	utility_types "github.com/PretendoNetwork/nex-protocols-go/utility/types"
)

func (protocol *Protocol) handleAssociateNexUniqueIDsWithMyPrincipalID(packet nex.PacketInterface) {
	if protocol.AssociateNexUniqueIDsWithMyPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Utility::AssociateNexUniqueIDsWithMyPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uniqueIDInfo := types.NewList[*utility_types.UniqueIDInfo]()
	uniqueIDInfo.Type = utility_types.NewUniqueIDInfo()

	err := uniqueIDInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AssociateNexUniqueIDsWithMyPrincipalID(fmt.Errorf("Failed to read uniqueIDInfo from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AssociateNexUniqueIDsWithMyPrincipalID(nil, packet, callID, uniqueIDInfo)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
