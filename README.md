# NEX servers with protocol support in Go

[![GoDoc](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go?status.svg)](https://godoc.org/github.com/PretendoNetwork/nex-protocols-go)

### Install

`go get github.com/PretendoNetwork/nex-protocols-go`

## Example (Friends server)

### Authentication server

```Golang
package main

import (
	"math/rand"
	"strconv"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go"
)

func main() {
	nexServer := nex.NewServer()
	nexServer.SetPrudpVersion(0)
	nexServer.SetSignatureVersion(1)
	nexServer.SetKerberosKeySize(16)
	nexServer.SetAccessKey("ridfebb9")

	authenticationServer := nexproto.NewAuthenticationProtocol(nexServer)

	// Handle Login RMC method
	authenticationServer.Login(func(client *nex.Client, callID uint32, username string) {
		userPID, _ := strconv.Atoi(username)
		serverPID := 2 // Quazal Rendez-Vous

		encryptedTicket := generateKerberosTicket(uint32(userPID), uint32(serverPID), nexServer.GetKerberosKeySize())

		// Build the response body
		stationURL := "prudps:/address=192.168.0.28;port=60001;CID=1;PID=2;sid=1;stream=10;type=2"
		serverName := "Pretendo Friends Auth"

		rmcResponseStream := nex.NewStream()
		rmcResponseStream.Grow(int64(4 + 4 + 4 + len(encryptedTicket) + 2 + len(stationURL) + 1 + 4 + 2 + 1 + 2 + len(serverName) + 1))

		rmcResponseStream.WriteU32LENext([]uint32{0x10001}) // success
		rmcResponseStream.WriteU32LENext([]uint32{uint32(userPID)})
		rmcResponseStream.WriteNEXBufferNext(encryptedTicket)
		// RVConnectionData
		rmcResponseStream.WriteNEXStringNext(stationURL) // Station
		rmcResponseStream.WriteU32LENext([]uint32{0})    // Special protocols list
		rmcResponseStream.WriteU16LENext([]uint16{0})    // Special protocols station
		rmcResponseStream.WriteNEXStringNext(serverName)

		rmcResponseBody := rmcResponseStream.Bytes()

		// Build response packet
		rmcResponse := nex.NewRMCResponse(nexproto.AuthenticationProtocolID, callID)
		rmcResponse.SetSuccess(nexproto.AuthenticationMethodLogin, rmcResponseBody)

		rmcResponseBytes := rmcResponse.Bytes()

		responsePacket := nex.NewPacketV0(client, nil)

		responsePacket.SetVersion(0)
		responsePacket.SetSource(0xA1)
		responsePacket.SetDestination(0xAF)
		responsePacket.SetType(nex.DataPacket)
		responsePacket.SetPayload(rmcResponseBytes)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)

		nexServer.Send(responsePacket)
	})

	// Handle RequestTicket RMC method
	authenticationServer.RequestTicket(func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32) {
		encryptedTicket := generateKerberosTicket(userPID, serverPID, nexServer.GetKerberosKeySize())

		// Build the response body
		rmcResponseStream := nex.NewStream()
		rmcResponseStream.Grow(int64(4 + 4 + len(encryptedTicket)))

		rmcResponseStream.WriteU32LENext([]uint32{0x10001}) // success
		rmcResponseStream.WriteNEXBufferNext(encryptedTicket)

		rmcResponseBody := rmcResponseStream.Bytes()

		// Build response packet
		rmcResponse := nex.NewRMCResponse(nexproto.AuthenticationProtocolID, callID)
		rmcResponse.SetSuccess(nexproto.AuthenticationMethodRequestTicket, rmcResponseBody)

		rmcResponseBytes := rmcResponse.Bytes()

		responsePacket := nex.NewPacketV0(client, nil)

		responsePacket.SetVersion(0)
		responsePacket.SetSource(0xA1)
		responsePacket.SetDestination(0xAF)
		responsePacket.SetType(nex.DataPacket)
		responsePacket.SetPayload(rmcResponseBytes)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)

		nexServer.Send(responsePacket)
	})

	nexServer.Listen("192.168.0.28:60000")
}

func generateKerberosTicket(userPID uint32, serverPID uint32, keySize int) []byte {
	nexPassword := "PASSWORD" // TODO: Get this from database

	sessionKey := make([]byte, keySize)
	rand.Read(sessionKey)

	// Create ticket body info
	kerberosTicketInfoKey := make([]byte, 16)
	//rand.Read(kerberosTicketInfoKey) // TODO: enable random keys and make them shared with secure server

	ticketInfoEncryption := nex.NewKerberosEncryption(kerberosTicketInfoKey)
	ticketInfoStream := nex.NewStream()
	ticketInfoStream.Grow(28)

	expiration := nex.NewDateTime(0)

	ticketInfoStream.WriteU64LENext([]uint64{expiration.Now()})
	ticketInfoStream.WriteU32LENext([]uint32{userPID})
	ticketInfoStream.WriteBytesNext(sessionKey)

	encryptedTicketInfo := ticketInfoEncryption.Encrypt(ticketInfoStream.Bytes())

	// Create ticket
	kerberosTicketKey := []byte(nexPassword)
	for i := 0; i < 65000+int(userPID)%1024; i++ {
		kerberosTicketKey = nex.MD5Hash(kerberosTicketKey)
	}

	ticketEncryption := nex.NewKerberosEncryption(kerberosTicketKey)
	ticketStream := nex.NewStream()
	ticketStream.Grow(int64(keySize + 4 + 4 + len(encryptedTicketInfo)))

	ticketStream.WriteBytesNext(sessionKey)
	ticketStream.WriteU32LENext([]uint32{serverPID})
	ticketStream.WriteNEXBufferNext(encryptedTicketInfo)

	return ticketEncryption.Encrypt(ticketStream.Bytes())
}
```

### Secure server (game server)

```Golang
func main() {
	nexServer := nex.NewServer()
	nexServer.SetPrudpVersion(0)
	nexServer.SetSignatureVersion(1)
	nexServer.SetKerberosKeySize(16)
	nexServer.SetAccessKey("ridfebb9")

	secureServer := nexproto.NewSecureProtocol(nexServer)
	friendsServer := nexproto.NewFriendsProtocol(nexServer)

	// Handle PRUDP CONNECT packet (not an RMC method)
	nexServer.On("Connect", func(packet *nex.PacketV0) {
		packet.GetSender().SetClientConnectionSignature(packet.GetConnectionSignature())

		payload := packet.GetPayload()
		stream := nex.NewStream(payload)

		ticketData := stream.ReadNEXBufferNext()
		requestData := stream.ReadNEXBufferNext()

		// TODO: use random key from auth server
		ticketDataEncryption := nex.NewKerberosEncryption([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
		decryptedTicketData := ticketDataEncryption.Decrypt(ticketData)
		ticketDataStream := nex.NewStream(decryptedTicketData)

		_ = ticketDataStream.ReadU64LENext(1)[0] // expiration time
		_ = ticketDataStream.ReadU32LENext(1)[0] // User PID
		sessionKey := ticketDataStream.ReadBytesNext(16)

		requestDataEncryption := nex.NewKerberosEncryption(sessionKey)
		decryptedRequestData := requestDataEncryption.Decrypt(requestData)
		requestDataStream := nex.NewStream(decryptedRequestData)

		_ = requestDataStream.ReadU32LENext(1)[0] // User PID
		_ = requestDataStream.ReadU32LENext(1)[0] //CID of secure server station url
		responseCheck := requestDataStream.ReadU32LENext(1)[0]

		responseValueStream := nex.NewStream(make([]byte, 4))
		responseValueBufferStream := nex.NewStream()
		responseValueBufferStream.Grow(8)

		responseValueStream.WriteU32LENext([]uint32{responseCheck + 1})
		responseValueBufferStream.WriteNEXBuffer(responseValueStream.Bytes())

		packet.GetSender().UpdateRC4Key(sessionKey)

		nexServer.AcknowledgePacket(packet, responseValueBufferStream.Bytes())
	})

	// Secure protocol handles

	// Handle Register RMC method
	secureServer.Register(func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL) {
		localStation := stationUrls[0]

		address := client.GetAddress().IP.String()
		port := string(client.GetAddress().Port)

		localStation.SetAddress(&address)
		localStation.SetPort(&port)

		localStationURL := localStation.EncodeToString()

		rmcResponseStream := nex.NewStream()
		rmcResponseStream.Grow(int64(4 + 4 + 2 + len(localStationURL) + 1))

		rmcResponseStream.WriteU32LENext([]uint32{0x10001}) // Success
		rmcResponseStream.WriteU32LENext([]uint32{uint32(secureServer.ConnectionIDCounter.Increment())})
		rmcResponseStream.WriteNEXStringNext(localStationURL)

		rmcResponseBody := rmcResponseStream.Bytes()

		// Build response packet
		rmcResponse := nex.NewRMCResponse(nexproto.SecureProtocolID, callID)
		rmcResponse.SetSuccess(nexproto.SecureMethodRegisterEx, rmcResponseBody)

		rmcResponseBytes := rmcResponse.Bytes()

		responsePacket := nex.NewPacketV0(client, nil)

		responsePacket.SetVersion(0)
		responsePacket.SetSource(0xA1)
		responsePacket.SetDestination(0xAF)
		responsePacket.SetType(nex.DataPacket)
		responsePacket.SetPayload(rmcResponseBytes)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)

		nexServer.Send(responsePacket)
	})

	// Handle RegisterEx RMC method
	secureServer.RegisterEx(func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData nexproto.NintendoLoginData) {
		// TODO: Validate loginData
		secureServer.RegisterHandler(client, callID, stationUrls)
	})

	// Friends (WiiU) protocol handles

	friendsServer.UpdateAndGetAllInformation(func(client *nex.Client, callID uint32, nnaInfo *nexproto.NNAInfo, presence *nexproto.NintendoPresenceV2, birthday *nex.DateTime) {
		comment := "Pretendo Online"
		datetime := nex.NewDateTime(0)

		rmcResponseStream := nex.NewStream()
		rmcResponseStream.Grow(int64(
			3 + // PrincipalPreference
			1 + 2 + len(comment) + 1 + 8 + // Comment
			4 + // FriendInfo List length
			4 + // FriendRequest (Sent) List length
			4 + // FriendRequest (Received) List length
			4 + // BlacklistedPrincipal List length
			1 + // Unknown
			4 + // PersistentNotification List length
			1)) // Unknown

		// TODO: Make the following fields into structs and encode them

		//PrincipalPreference
		rmcResponseStream.WriteByteNext(0)
		rmcResponseStream.WriteByteNext(0)
		rmcResponseStream.WriteByteNext(0)
		//Comment
		rmcResponseStream.WriteByteNext(0)
		rmcResponseStream.WriteNEXStringNext(comment)
		rmcResponseStream.WriteU64LENext([]uint64{datetime.Now()})
		//List<FriendInfo>
		rmcResponseStream.WriteU32LENext([]uint32{0})
		//List<FriendRequest> (Sent)
		rmcResponseStream.WriteU32LENext([]uint32{0})
		//List<FriendRequest> (Received)
		rmcResponseStream.WriteU32LENext([]uint32{0})
		//List<BlacklistedPrincipal>
		rmcResponseStream.WriteU32LENext([]uint32{0})
		//Unknown
		rmcResponseStream.WriteByteNext(0)
		//List<PersistentNotification>
		rmcResponseStream.WriteU32LENext([]uint32{0})
		//Unknown
		rmcResponseStream.WriteByteNext(0)

		rmcResponseBody := rmcResponseStream.Bytes()

		// Build response packet
		rmcResponse := nex.NewRMCResponse(nexproto.FriendsProtocolID, callID)
		rmcResponse.SetSuccess(nexproto.FriendsMethodUpdateAndGetAllInformation, rmcResponseBody)

		rmcResponseBytes := rmcResponse.Bytes()

		responsePacket := nex.NewPacketV0(client, nil)

		responsePacket.SetVersion(0)
		responsePacket.SetSource(0xA1)
		responsePacket.SetDestination(0xAF)
		responsePacket.SetType(nex.DataPacket)
		responsePacket.SetPayload(rmcResponseBytes)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)

		nexServer.Send(responsePacket)
	})

	nexServer.Listen("192.168.0.28:60001")
}
```