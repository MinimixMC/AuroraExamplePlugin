package events

import (
	"github.com/MinimixMC/AuroraAPI/API/aurora"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/chat"
	"github.com/MinimixMC/AuroraAPI/API/aurora/data/status"
	"github.com/MinimixMC/AuroraAPI/API/aurora/event"
	"github.com/rs/zerolog/log"
)

func HandlePingEvent(e aurora.Event) aurora.Event {
	event := e.(*event.Ping)
	log.Info().Msg("Ping event")
	event.Status = &status.StatusResponse{
		Version: status.VersionInfo{
			Name:     "Aurora 1.20.4",
			Protocol: 765,
		},
		Players: status.PlayersInfo{
			Max:    1,
			Online: 0,
			Sample: []status.PlayerSample{},
		},
		Description:         chat.BuildMessage("Glorious Aurora"),
		EnforcesSecureChat:  false,
		PreviewsChat:        false,
		PreventsChatReports: true,
	}
	return e
}
