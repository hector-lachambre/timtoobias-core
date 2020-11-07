package ports

import (
	"gitlab.com/timtoobias-projects/timtoobias-core/entities"
)

type VideosRepository interface {
	GetLastVideoByChannelID(ID string) (*entities.Video, error)
}
