// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SyncFriend sets the SyncFriend handler function
func (protocol *Protocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64) uint32) {
	protocol.syncFriendHandler = handler
}

func (protocol *Protocol) handleSyncFriend(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.syncFriendHandler == nil {
		globals.Logger.Warning("Friends3DS::SyncFriend not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.syncFriendHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.syncFriendHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lfcList, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.syncFriendHandler(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), client, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.syncFriendHandler(nil, client, callID, lfc, pids, lfcList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
