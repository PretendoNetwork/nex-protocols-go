// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetApplicationInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SetApplicationInfo == nil {
		globals.Logger.Warning("AAUser::SetApplicationInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationInfo, err := nex.StreamReadListStructure(parametersStream, aauser_types.NewApplicationInfo())
	if err != nil {
		_, errorCode = protocol.SetApplicationInfo(fmt.Errorf("Failed to read applicationInfo from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SetApplicationInfo(nil, packet, callID, applicationInfo)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
