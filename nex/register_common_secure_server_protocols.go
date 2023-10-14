package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
	matchmaking "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking-ext"
	matchmakeextension "github.com/PretendoNetwork/nex-protocols-common-go/matchmake-extension"
	nattraversal "github.com/PretendoNetwork/nex-protocols-common-go/nat-traversal"
	"github.com/PretendoNetwork/kid-icarus-uprising/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"

	"fmt"
)

func cleanupSearchMatchmakeSessionHandler(matchmakeSession *match_making_types.MatchmakeSession){
	matchmakeSession.Attributes[2] = 0
	fmt.Println(matchmakeSession.String())
}

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
	matchmaking.NewCommonMatchMakingProtocol(globals.SecureServer)
	matchmakingext.NewCommonMatchMakingExtProtocol(globals.SecureServer)
	commonMatchmakeExtensionProtocol := matchmakeextension.NewCommonMatchmakeExtensionProtocol(globals.SecureServer)
	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession(cleanupSearchMatchmakeSessionHandler)
	nattraversal.NewCommonNATTraversalProtocol(globals.SecureServer)
}
