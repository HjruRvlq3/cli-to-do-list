package app

import (
	"github.com/HjruRvlq3/cli-to-do-list/internal/input"
	"github.com/HjruRvlq3/cli-to-do-list/internal/service"
	"github.com/HjruRvlq3/cli-to-do-list/internal/storage"
)

func Run() {
	repo := storage.JSONRepository{}
	srv := service.New(repo)
	input.MovManager(srv)
}
