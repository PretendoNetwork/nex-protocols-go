// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByName sets the AddFriendByName handler function
func (protocol *Protocol) AddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, strPlayerName string, uiDetails uint32, strMessage string) uint32) {
	protocol.addFriendByNameHandler = handler
}

func (protocol *Protocol) handleAddFriendByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendByNameHandler == nil {
		globals.Logger.Warning("Friends::AddFriendByName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strPlayerName, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendByNameHandler(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendByNameHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendByNameHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, "", 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendByNameHandler(nil, packet, callID, strPlayerName, uiDetails, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
