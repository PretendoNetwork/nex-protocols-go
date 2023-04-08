package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestURLs sets the RequestURLs handler function
func (protocol *SecureConnectionProtocol) RequestURLs(handler func(err error, client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)) {
	protocol.RequestURLsHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleRequestURLs(packet nex.PacketInterface) {
	if protocol.RequestURLsHandler == nil {
		globals.Logger.Warning("SecureConnection::RequestURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[SecureConnection::RequestURLs] Data length too small")
		go protocol.RequestURLsHandler(err, client, callID, 0, 0)
		return
	}

	stationCID := parametersStream.ReadUInt32LE()
	stationPID := parametersStream.ReadUInt32LE()

	go protocol.RequestURLsHandler(nil, client, callID, stationCID, stationPID)
}
