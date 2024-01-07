// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetProfiles(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetProfiles == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetProfiles not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pidList := types.NewList[*types.PID]()
	pidList.Type = types.NewPID(0)
	err = pidList.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetProfiles(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetProfiles(nil, packet, callID, pidList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
