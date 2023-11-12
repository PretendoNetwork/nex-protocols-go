// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindBySQLQuery sets the FindBySQLQuery handler function
func (protocol *Protocol) FindBySQLQuery(handler func(err error, packet nex.PacketInterface, callID uint32, strQuery string, resultRange *nex.ResultRange) uint32) {
	protocol.findBySQLQueryHandler = handler
}

func (protocol *Protocol) handleFindBySQLQuery(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.findBySQLQueryHandler == nil {
		globals.Logger.Warning("MatchMaking::FindBySQLQuery not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strQuery, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.findBySQLQueryHandler(fmt.Errorf("Failed to read strQuery from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		errorCode = protocol.findBySQLQueryHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, "", nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.findBySQLQueryHandler(nil, packet, callID, strQuery, resultRange.(*nex.ResultRange))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
