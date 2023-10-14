package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/kid-icarus-uprising/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(0)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(3)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(2,7,2))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("58a7e494")

	globals.Timeline = make(map[uint32][]uint8)

	globals.SecureServer.On("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Println("==Kid Icarus: Uprising - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_KIU_SECURE_SERVER_PORT")))
}
