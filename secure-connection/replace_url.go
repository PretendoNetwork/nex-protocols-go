package secure_connection

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceURL sets the ReplaceURL handler function
func (protocol *SecureConnectionProtocol) ReplaceURL(handler func(err error, client *nex.Client, callID uint32, oldStation *nex.StationURL, newStation *nex.StationURL)) {
	protocol.ReplaceURLHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleReplaceURL(packet nex.PacketInterface) {
	if protocol.ReplaceURLHandler == nil {
		globals.Logger.Warning("SecureConnection::ReplaceURL not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	oldStationString, err := parametersStream.ReadString()

	if err != nil {
		go protocol.ReplaceURLHandler(err, client, callID, nex.NewStationURL(""), nex.NewStationURL(""))
		return
	}

	newStationString, err := parametersStream.ReadString()

	if err != nil {
		go protocol.ReplaceURLHandler(err, client, callID, nex.NewStationURL(""), nex.NewStationURL(""))
		return
	}

	oldStation := nex.NewStationURL(oldStationString)
	newStation := nex.NewStationURL(newStationString)

	go protocol.ReplaceURLHandler(nil, client, callID, oldStation, newStation)
}
