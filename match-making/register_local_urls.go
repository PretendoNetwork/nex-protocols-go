// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterLocalURLs sets the RegisterLocalURLs handler function
func (protocol *Protocol) RegisterLocalURLs(handler func(err error, client *nex.Client, callID uint32, gid uint32, lstURLs []*nex.StationURL) uint32) {
	protocol.registerLocalURLsHandler = handler
}

func (protocol *Protocol) handleRegisterLocalURLs(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerLocalURLsHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURLs not implemented")
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
		errorCode = protocol.registerLocalURLsHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	lstURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		errorCode = protocol.registerLocalURLsHandler(fmt.Errorf("Failed to read lstURLs from parameters. %s", err.Error()), client, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.registerLocalURLsHandler(nil, client, callID, gid, lstURLs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
