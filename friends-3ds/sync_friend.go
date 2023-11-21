// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSyncFriend(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.SyncFriend == nil {
		globals.Logger.Warning("Friends3DS::SyncFriend not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pids, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		_, errorCode = protocol.SyncFriend(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
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
