package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Register sets the Register handler function
func (protocol *SecureConnectionProtocol) Register(handler func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL)) {
	protocol.RegisterHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleRegister(packet nex.PacketInterface) {
	if protocol.RegisterHandler == nil {
		globals.Logger.Warning("SecureConnection::Register not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[SecureConnection::Register] Data missing list length")
		go protocol.RegisterHandler(err, client, callID, make([]*nex.StationURL, 0))
		return
	}

	stationUrls := parametersStream.ReadListStationURL()

	go protocol.RegisterHandler(nil, client, callID, stationUrls)
}
