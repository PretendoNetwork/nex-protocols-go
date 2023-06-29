package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterEx sets the RegisterEx handler function
func (protocol *SecureConnectionProtocol) RegisterEx(handler func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL, hCustomData *nex.DataHolder)) {
	protocol.RegisterExHandler = handler
}

func (protocol *SecureConnectionProtocol) handleRegisterEx(packet nex.PacketInterface) {
	if protocol.RegisterExHandler == nil {
		globals.Logger.Warning("SecureConnection::RegisterEx not implemented")
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
		go protocol.RegisterExHandler(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	hCustomData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.RegisterExHandler(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.RegisterExHandler(nil, client, callID, vecMyURLs, hCustomData)
}
