// vücut ölçüleri
package models

import "gorm.io/gorm"

type BodyMeasurement struct {
	UserID     uint    `json:"user_id"`
	Chest      uint16  `json:"chest"`                                        // göğus ölçüsü
	LeftArm    uint16  `json:"left_arm"`                                     // sol kol
	RigthArm   uint16  `json:"rigth_arm"`                                    // sağ kol
	Abdomen    uint16  `json:"abdomen"`                                      // karın
	Waist      uint16  `json:"waist"`                                        // bel
	Hips       uint16  `json:"hips"`                                         // kalça
	LeftThigh  uint16  `json:"left_thigh"`                                   // sol üst bacak
	RigthThigh uint16  `json:"rigth_thigh"`                                  // sağ üst bacak
	FatRate    uint8   `json:"fat_rate"`                                     // yağ oranı
	Weight     uint16  `json:"weight"`                                       // kilo
	Height     float32 `json:"height" gorm:"type:DECIMAL(13,2);default:'0'"` // boy
	BMI        string  `json:"body_mass_index" gorm:"-"`
}

func (b *BodyMeasurement) AfterFind(db *gorm.DB) error {
	bmi := float32(b.Weight) / (b.Height * b.Height)
	if bmi >= 0 && bmi <= 18.4 {
		b.BMI = "Zayıf"
	} else if bmi >= 18.5 && bmi <= 24.9 {
		b.BMI = "Normal"
	} else if bmi >= 25 && bmi <= 29.9 {
		b.BMI = "Kilolu"
	} else if bmi >= 30 && bmi <= 34.9 {
		b.BMI = "Şişman (1. dereceden obez)"
	} else if bmi >= 35 && bmi <= 44.9 {
		b.BMI = "Şişman (2. dereceden obez)"
	} else if bmi >= 45 {
		b.BMI = "Şişman (3. dereceden obez)"
	}
	return nil
}
