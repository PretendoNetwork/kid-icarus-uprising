package nex

import (
	"context"

	"github.com/PretendoNetwork/kid-icarus-uprising/database"
	"github.com/PretendoNetwork/kid-icarus-uprising/globals"
	local_globals "github.com/PretendoNetwork/kid-icarus-uprising/globals"
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	mm_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	secure := common_secure.NewCommonProtocol(secureProtocol)
	secure.CreateReportDBRecord = func(pid types.PID, reportID types.UInt32, reportData types.QBuffer) error {
		return nil
	}

	matchmakingManager := common_globals.NewMatchmakingManager(local_globals.SecureEndpoint, database.Postgres)

	matchmakingManager.GetUserFriendPIDs = globals.GetUserFriendPIDs
	natTraversalProtocol := nat_traversal.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	commonMatchMakingProtocol := common_match_making.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(matchmakingManager)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(matchmakingManager)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	local_globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(matchmakingManager)

	commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession = func(matchmakeSession *mm_types.MatchmakeSession) {
		matchmakeSession.Attributes[0] = types.NewUInt32(0)
		matchmakeSession.Attributes[1] = types.NewUInt32(0)
		matchmakeSession.Attributes[5] = types.NewUInt32(0)
		matchmakeSession.Attributes[2] = types.NewUInt32(0)
	}
	commonMatchmakeExtensionProtocol.OnAfterAutoMatchmakeWithSearchCriteriaPostpone = func(packet nex.PacketInterface, lstSearchCriteria types.List[mm_types.MatchmakeSessionSearchCriteria], anyGathering types.AnyObjectHolder[mm_types.GatheringInterface], strMessage types.String) {
		matchmakingManager.Database.ExecContext(context.Background(), "UPDATE matchmaking.matchmake_sessions SET open_participation=true")

	}
	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = func(searchCriterias types.List[mm_types.MatchmakeSessionSearchCriteria]) {
		for i := 0; i < len(searchCriterias); i++ {
			searchCriterias[i].Attribs[4] = types.NewString("")
		}
	}
	commonMatchmakeExtensionProtocol.OnAfterUpdateNotificationData = func(packet nex.PacketInterface, uiType, uiParam1, uiParam2 types.UInt32, strParam types.String) {}
}
