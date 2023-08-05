// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByGroup sets the FindByGroup handler function
func (protocol *Protocol) FindByGroup(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32) uint32) {
	protocol.findByGroupHandler = handler
}

func (protocol *Protocol) handleFindByGroup(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByGroupHandler == nil {
		globals.Logger.Warning("PersistentStore::FindByGroup not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findByGroupHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByGroupHandler(nil, client, callID, uiGroup)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
