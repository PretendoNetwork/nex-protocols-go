// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateURLs sets the UpdateURLs handler function
func (protocol *Protocol) UpdateURLs(handler func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL)) {
	protocol.UpdateURLsHandler = handler
}

func (protocol *Protocol) handleUpdateURLs(packet nex.PacketInterface) {
	if protocol.UpdateURLsHandler == nil {
		globals.Logger.Warning("SecureConnection::UpdateURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.UpdateURLsHandler(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateURLsHandler(nil, client, callID, vecMyURLs)
}
