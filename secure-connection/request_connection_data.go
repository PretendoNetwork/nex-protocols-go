package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestConnectionData sets the RequestConnectionData handler function
func (protocol *SecureConnectionProtocol) RequestConnectionData(handler func(err error, client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)) {
	protocol.RequestConnectionDataHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleRequestConnectionData(packet nex.PacketInterface) {
	if protocol.RequestConnectionDataHandler == nil {
		globals.Logger.Warning("SecureConnection::RequestConnectionData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 8 {
		err := errors.New("[SecureConnection::RequestConnectionData] Data length too small")
		go protocol.RequestConnectionDataHandler(err, client, callID, 0, 0)
		return
	}

	stationCID := parametersStream.ReadUInt32LE()
	stationPID := parametersStream.ReadUInt32LE()

	go protocol.RequestConnectionDataHandler(nil, client, callID, stationCID, stationPID)
}
