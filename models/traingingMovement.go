package models

// programlar ile hareketler arasında 3.tablo
type Training_movent struct {
	TrainingProgramID uint             `json:"training_program_id"`
	TrainingProgram   *TrainingProgram `json:"training_program" gorm:"foreignKey:training_program_id"`
	MovementID        uint             `json:"movement_id"`
	Movement          *Movement        `json:"movement" gorm:"foreignKey:movement_id"`
	RestPeriod        uint16           `json:"rest_period"` // dinlenme süresi
	Reps              int8             `json:"reps"`        // tekrar sayısı
	Sets              int8             `json:"sets"`        // set sayısı
}
