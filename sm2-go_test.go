package sm

import (
	"math"
	"testing"
	"time"
)

func almostFloat64Equal(a, b float64) bool {
	return math.Abs(a-b) < .00001
}

type TestCard struct {
	repetitions uint8
	easiness    float64
	interval    int
}

func (c *TestCard) Repetitions() uint8 {
	return c.repetitions
}

func (c *TestCard) Easiness() float64 {
	return c.easiness
}

func (c *TestCard) Interval() int {
	return c.interval
}

func (c *TestCard) SetRepetitions(repetitions uint8) {
	c.repetitions = repetitions
}

func (c *TestCard) SetEasiness(easiness float64) {
	c.easiness = easiness
}

func (c *TestCard) SetInterval(interval int) {
	c.interval = interval
}

func TestSM2(t *testing.T) {
	testCases := []struct {
		card             *TestCard
		quality          Quality
		expectedReps     uint8
		expectedEasiness float64
		expectedInterval int
	}{
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: BLACKOUT, expectedReps: 0, expectedEasiness: 1.5, expectedInterval: 1},
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: CORRECT_REMEMBERED, expectedReps: 0, expectedEasiness: 1.76, expectedInterval: 1},
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: CORRECT_EASY_TO_RECALL, expectedReps: 0, expectedEasiness: 1.98, expectedInterval: 1},
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: CORRECT_WITH_DIFFICULTY, expectedReps: 4, expectedEasiness: 2.16, expectedInterval: 28},
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: CORRECT_AFTER_HESITATION, expectedReps: 4, expectedEasiness: 2.3, expectedInterval: 28},
		{card: &TestCard{repetitions: 3, easiness: 2.3, interval: 12}, quality: PERFECT_RESPONSE, expectedReps: 4, expectedEasiness: 2.4, expectedInterval: 28},
		{card: &TestCard{repetitions: 3, easiness: 1.3, interval: 12}, quality: BLACKOUT, expectedReps: 0, expectedEasiness: 1.3, expectedInterval: 1},
		{card: &TestCard{repetitions: 1, easiness: 2.5, interval: 1}, quality: CORRECT_AFTER_HESITATION, expectedReps: 2, expectedEasiness: 2.5, expectedInterval: 6},
	}

	for _, tc := range testCases {
		reviewDetails := SM2(tc.card, tc.quality)
		if tc.card.repetitions != tc.expectedReps {
			t.Errorf("Unexpected card repetitions: expected %d, got %d", tc.expectedReps, tc.card.repetitions)
		}
		if !almostFloat64Equal(tc.expectedEasiness, tc.card.easiness) {
			t.Errorf("Unexpected card easiness: expected %f, got %f", tc.expectedEasiness, tc.card.easiness)
		}
		if reviewDetails.repetitions != tc.expectedReps {
			t.Errorf("Unexpected reviewDetails repetitions: expected %d, got %d", tc.expectedReps, reviewDetails.repetitions)
		}
		if !almostFloat64Equal(tc.expectedEasiness, reviewDetails.easiness) {
			t.Errorf("Unexpected reviewDetails easiness: expected %f, got %f", tc.expectedEasiness, reviewDetails.easiness)
		}
		if reviewDetails.interval != tc.expectedInterval {
			t.Errorf("Expected reviewDetails interval %d, got %d", tc.expectedInterval, reviewDetails.interval)
		}

		// Allow up to 2 seconds for time difference between dueDate and now
		dueDateExpected := time.Now().UTC().AddDate(0, 0, tc.expectedInterval)
		dueDateDiff := reviewDetails.dueDate.Sub(dueDateExpected)
		if dueDateDiff > 2*time.Second || dueDateDiff < -2*time.Second {
			t.Errorf("Expected due date to be within 2 seconds of %s, got %s", dueDateExpected, reviewDetails.dueDate)
		}
	}
}
