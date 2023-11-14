// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleInsertCustomItem(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.InsertCustomItem == nil {
		globals.Logger.Warning("PersistentStore::InsertCustomItem not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	hData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read hData from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReplace, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.InsertCustomItem(nil, packet, callID, uiGroup, strTag, hData, bReplace)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
