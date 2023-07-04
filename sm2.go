package sm

import (
	"math"
	"time"
)

type Quality int8

const (
	BLACKOUT Quality = iota
	CORRECT_REMEMBERED
	CORRECT_EASY_TO_RECALL
	CORRECT_WITH_DIFFICULTY
	CORRECT_AFTER_HESITATION
	PERFECT_RESPONSE
)

type Card interface {
	Repetitions() uint8
	Easiness() float64
	Interval() int
	SetRepetitions(repetitions uint8)
	SetEasiness(easiness float64)
	SetInterval(interval int)
}

type ReviewDetails struct {
	dueDate     time.Time
	easiness    float64
	repetitions uint8
	interval    int
}

func SM2(card Card, quality Quality) *ReviewDetails {
	repetitions := card.Repetitions()
	easiness := card.Easiness()
	interval := card.Interval()

	if quality < CORRECT_WITH_DIFFICULTY {
		interval = 1
		repetitions = 0
	} else {
		if repetitions == 0 {
			interval = 1
		} else if repetitions == 1 {
			interval = 6
		} else {
			interval = int(math.Round(float64(interval) * easiness))
		}
		repetitions++
	}

	qualityFloat := float64(quality)
	easinessCalculation := easiness + 0.1 - (5.0-qualityFloat)*(0.08+(5.0-qualityFloat)*0.02)

	easiness = math.Max(1.3, easinessCalculation)

	dueDate := time.Now().UTC().AddDate(0, 0, interval)
	reviewDetails := &ReviewDetails{
		repetitions: repetitions,
		easiness:    easiness,
		interval:    interval,
		dueDate:     dueDate,
	}

	card.SetRepetitions(repetitions)
	card.SetEasiness(easiness)
	card.SetInterval(interval)

	return reviewDetails
}
