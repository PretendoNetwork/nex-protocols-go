// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByNameRegex sets the FindByNameRegex handler function
func (protocol *Protocol) FindByNameRegex(handler func(err error, client *nex.Client, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) uint32) {
	protocol.findByNameRegexHandler = handler
}

func (protocol *Protocol) handleFindByNameRegex(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findByNameRegexHandler == nil {
		globals.Logger.Warning("AccountManagement::FindByNameRegex not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.findByNameRegexHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strRegex, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findByNameRegexHandler(fmt.Errorf("Failed to read strRegex from parameters. %s", err.Error()), client, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findByNameRegexHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findByNameRegexHandler(nil, client, callID, uiGroups, strRegex, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
