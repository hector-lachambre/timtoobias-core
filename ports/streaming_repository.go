package ports

import (
	"time"

	"gitlab.com/timtoobias-projects/timtoobias-core/entities"
)

type StreamingRepository interface {
	GetStreamingStatusByID(ID string) (*entities.StreamingStatus, error)
	GetLastSync() *time.Time
}
