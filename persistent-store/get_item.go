// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetItem sets the GetItem handler function
func (protocol *Protocol) GetItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup uint32, strTag string) uint32) {
	protocol.getItemHandler = handler
}

func (protocol *Protocol) handleGetItem(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getItemHandler == nil {
		globals.Logger.Warning("PersistentStore::GetItem not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.getItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getItemHandler(nil, packet, callID, uiGroup, strTag)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
