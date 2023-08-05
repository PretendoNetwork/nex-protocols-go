// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSessionURLs sets the GetSessionURLs handler function
func (protocol *Protocol) GetSessionURLs(handler func(err error, client *nex.Client, callID uint32, gid uint32) uint32) {
	protocol.getSessionURLsHandler = handler
}

func (protocol *Protocol) handleGetSessionURLs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSessionURLsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetSessionURLs not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getSessionURLsHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSessionURLsHandler(nil, client, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
