# SM (SuperMemo 2) Go Package

SM is a Golang package that provides an implementation of the SuperMemo 2 (SM2) Algorithm. The SM2 Algorithm is a learning method that enables efficient memorization by scheduling review times at optimally spaced intervals. This package allows developers to integrate the algorithm into their applications with ease.

## Usage

Here is a basic usage example:

```go
package main

import (
	"fmt"
	sm "github.com/mateuszmrw/sm2-go"
)

func main() {
	card := // Initialize your card
	quality := sm.PERFECT_RESPONSE

	// Use SM2 on a card
	details := sm.SM2(card, quality)

	fmt.Println(details.Interval)
}
```

## API

The SM package mainly provides the following API:

- `SM2(card Card, quality Quality) *ReviewDetails`: Returns the review details for a card after applying the SM2 algorithm with a given quality.

And also provides the following types:

- `Card`: An interface for a card. It has methods `Repetitions()`, `Easiness()`, `Interval()`, `SetRepetitions(repetitions uint8)`, `SetEasiness(easiness float64)`, and `SetInterval(interval int)`.
- `Quality`: An enum-like type that indicates the quality of a card's recall. It includes `BLACKOUT`, `CORRECT_REMEMBERED`, `CORRECT_EASY_TO_RECALL`, `CORRECT_WITH_DIFFICULTY`, `CORRECT_AFTER_HESITATION`, and `PERFECT_RESPONSE`.
- `ReviewDetails`: A structure that holds the details of a card review. It includes `dueDate`, `easiness`, `repetitions`, and `interval`.

## Tests
To run the test suite, simply execute:

```bash
    go test
```