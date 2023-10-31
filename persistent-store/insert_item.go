// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InsertItem sets the InsertItem handler function
func (protocol *Protocol) InsertItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string, bufData []byte, bReplace bool) uint32) {
	protocol.insertItemHandler = handler
}

func (protocol *Protocol) handleInsertItem(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.insertItemHandler == nil {
		globals.Logger.Warning("PersistentStore::InsertItem not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.insertItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.insertItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bufData, err := parametersStream.ReadBuffer()
	if err != nil {
		errorCode = protocol.insertItemHandler(fmt.Errorf("Failed to read bufData from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReplace, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.insertItemHandler(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, 0, "", nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.insertItemHandler(nil, packet, callID, uiGroup, strTag, bufData, bReplace)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
