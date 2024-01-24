// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetFriendPersistentInfo == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPersistentInfo not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
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
		_, errorCode = protocol.GetFriendPersistentInfo(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetFriendPersistentInfo(nil, packet, callID, pidList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
