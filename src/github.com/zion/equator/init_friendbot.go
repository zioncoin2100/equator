package equator

import (
	"github.com/zion/go/strkey"
	"github.com/zion/equator/friendbot"
)

func initFriendbot(app *App) {
	if app.config.FriendbotSecret == "" {
		return
	}

	// ensure its a seed if its not blank
	strkey.MustDecode(strkey.VersionByteSeed, app.config.FriendbotSecret)

	app.friendbot = &friendbot.Bot{
		Secret:    app.config.FriendbotSecret,
		Submitter: app.submitter,
		Network:   app.networkPassphrase,
	}

}

func init() {
	appInit.Add("friendbot", initFriendbot, "txsub", "zionCoreInfo")
}
