// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InsertCustomItem sets the InsertCustomItem handler function
func (protocol *Protocol) InsertCustomItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool) uint32) {
	protocol.insertCustomItemHandler = handler
}

func (protocol *Protocol) handleInsertCustomItem(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.insertCustomItemHandler == nil {
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
		errorCode = protocol.insertCustomItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.insertCustomItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	hData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.insertCustomItemHandler(fmt.Errorf("Failed to read hData from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReplace, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.insertCustomItemHandler(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.insertCustomItemHandler(nil, packet, callID, uiGroup, strTag, hData, bReplace)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
