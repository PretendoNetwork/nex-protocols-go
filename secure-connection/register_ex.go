package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterEx sets the RegisterEx handler function
func (protocol *SecureConnectionProtocol) RegisterEx(handler func(err error, client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData *nex.DataHolder)) {
	protocol.RegisterExHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleRegisterEx(packet nex.PacketInterface) {
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[SecureConnection::RegisterEx] Data missing list length")
		go protocol.RegisterExHandler(err, client, callID, make([]*nex.StationURL, 0), nex.NewDataHolder())
		return
	}

	stationURLCount := parametersStream.ReadUInt32LE()
	stationUrls := make([]*nex.StationURL, 0)

	for i := 0; i < int(stationURLCount); i++ {
		stationString, err := parametersStream.ReadString()

		if err != nil {
			go protocol.RegisterExHandler(err, client, callID, stationUrls, nex.NewDataHolder())
			return
		}

		station := nex.NewStationURL(stationString)
		stationUrls = append(stationUrls, station)
	}

	dataHolder := parametersStream.ReadDataHolder()

	if dataHolder.TypeName() != "NintendoLoginData" && dataHolder.TypeName() != "AccountExtraInfo" {
		err := errors.New("[SecureConnection::RegisterEx] Data holder name does not match")
		go protocol.RegisterExHandler(err, client, callID, stationUrls, nex.NewDataHolder())
		return
	}

	go protocol.RegisterExHandler(nil, client, callID, stationUrls, dataHolder)
}
