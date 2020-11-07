package usecases

import (
	"gitlab.com/timtoobias-projects/timtoobias-core/entities"
	"gitlab.com/timtoobias-projects/timtoobias-core/ports"
)

type GetMediaInfosOutputBoundary interface {
	Error(e error)
	Success(output *GetMediaInfosOutputMessage)
	GetViewModel() interface{}
}

type GetMediaInfosInteractor struct {
	StreamingRepository ports.StreamingRepository
	VideosRepository    ports.VideosRepository
	Presenter           GetMediaInfosOutputBoundary
}

type GetMediaInfosInputMessage struct {
	StreamingChannelID       string
	MainVideosChannelID      string
	SecondaryVideosChannelID string
}

type GetMediaInfosOutputMessage struct {
	Status                      *entities.StreamingStatus
	LastVideoOnMainChannel      *entities.Video
	LastVideoOnSecondaryChannel *entities.Video
}

// Execute the usecase
func (interactor *GetMediaInfosInteractor) Execute(input *GetMediaInfosInputMessage) {
	status, err := interactor.StreamingRepository.GetStreamingStatusByID(input.StreamingChannelID)

	if err != nil {
		interactor.Presenter.Error(err)
		return
	}

	lastVideoOnMainChannel, err := interactor.VideosRepository.
		GetLastVideoByChannelID(input.MainVideosChannelID)

	if err != nil {
		interactor.Presenter.Error(err)
		return
	}

	lastVideoOnSecondaryChannel, err := interactor.VideosRepository.
		GetLastVideoByChannelID(input.SecondaryVideosChannelID)

	if err != nil {
		interactor.Presenter.Error(err)
		return
	}

	output := &GetMediaInfosOutputMessage{
		Status:                      status,
		LastVideoOnMainChannel:      lastVideoOnMainChannel,
		LastVideoOnSecondaryChannel: lastVideoOnSecondaryChannel,
	}

	interactor.Presenter.Success(output)
}
