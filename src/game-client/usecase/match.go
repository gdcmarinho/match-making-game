package usecase

import (
	"marinho/match-making-game/repository"
)

func FindMatch() {
	repository.Publish()
}