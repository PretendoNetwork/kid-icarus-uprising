package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/kid-icarus-uprising/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(0)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(3)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(2,7,2))
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("58a7e494")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Println("==Kid Icarus: Uprising - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_KIU_AUTHENTICATION_SERVER_PORT")))
}
