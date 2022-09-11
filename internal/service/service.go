package service

import (
	"github.com/Gym-Apps/gym-backend/internal/utils"
	"github.com/Gym-Apps/gym-backend/pkg"
)

type Service struct {
	Utils   utils.IUtils
	Service pkg.ILogger
}
