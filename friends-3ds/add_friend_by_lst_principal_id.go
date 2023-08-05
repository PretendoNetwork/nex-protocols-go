// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendBylstPrincipalID sets the AddFriendBylstPrincipalID handler function
func (protocol *Protocol) AddFriendBylstPrincipalID(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32) uint32) {
	protocol.addFriendBylstPrincipalIDHandler = handler
}

func (protocol *Protocol) handleAddFriendBylstPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendBylstPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::AddFriendBylstPrincipalID not implemented")
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
		errorCode = protocol.addFriendBylstPrincipalIDHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendBylstPrincipalIDHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendBylstPrincipalIDHandler(nil, client, callID, lfc, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
