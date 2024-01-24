// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSyncFriend(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SyncFriend == nil {
		globals.Logger.Warning("Friends3DS::SyncFriend not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lfc := types.NewPrimitiveU64(0)
	err = lfc.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)
	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lfcList := types.NewList[*types.PrimitiveU64]()
	lfcList.Type = types.NewPrimitiveU64(0)
	err = lfcList.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SyncFriend(nil, packet, callID, lfc, pids, lfcList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
