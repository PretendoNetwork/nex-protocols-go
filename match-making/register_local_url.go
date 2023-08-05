// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterLocalURL sets the RegisterLocalURL handler function
func (protocol *Protocol) RegisterLocalURL(handler func(err error, client *nex.Client, callID uint32, gid uint32, url *nex.StationURL) uint32) {
	protocol.registerLocalURLHandler = handler
}

func (protocol *Protocol) handleRegisterLocalURL(packet nex.PacketInterface) {
	if protocol.registerLocalURLHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURL not implemented")
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
		go protocol.registerLocalURLHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.registerLocalURLHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), client, callID, 0, nil)
	}

	go protocol.registerLocalURLHandler(nil, client, callID, gid, url)
}
