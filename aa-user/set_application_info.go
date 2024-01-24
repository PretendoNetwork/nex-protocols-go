// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetApplicationInfo(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SetApplicationInfo == nil {
		globals.Logger.Warning("AAUser::SetApplicationInfo not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	applicationInfo := types.NewList[*aauser_types.ApplicationInfo]()
	applicationInfo.Type = aauser_types.NewApplicationInfo()
	err = applicationInfo.ExtractFrom(parametersStream)
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
