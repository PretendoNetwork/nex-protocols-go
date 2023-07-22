// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterLocalURL sets the RegisterLocalURL handler function
func (protocol *MatchMakingProtocol) RegisterLocalURL(handler func(err error, client *nex.Client, callID uint32, gid uint32, url *nex.StationURL)) {
	protocol.registerLocalURLHandler = handler
}

func (protocol *MatchMakingProtocol) handleRegisterLocalURL(packet nex.PacketInterface) {
	if protocol.registerLocalURLHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterLocalURL not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
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