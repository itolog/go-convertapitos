package timeutils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// ParseDuration converts a string of format “2h,” “7d,” “1w2d3h4m5s,” etc. to timeutils.Duration
// Supported units:
// - s - seconds
// - m - minutes
// - h - hours
// - d - days
// - w - weeks
func ParseDuration(input string) (time.Duration, error) {
	if input == "" {
		return 0, fmt.Errorf("пустая строка времени")
	}

	re := regexp.MustCompile(`(\d+)([smhdw])`)
	matches := re.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 {
		return 0, fmt.Errorf("некорректный формат строки времени: %s", input)
	}

	var totalDuration time.Duration

	for _, match := range matches {
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, fmt.Errorf("ошибка при парсинге числового значения: %w", err)
		}

		unit := match[2]
		var unitDuration time.Duration

		switch unit {
		case "s":
			unitDuration = time.Duration(value) * time.Second
		case "m":
			unitDuration = time.Duration(value) * time.Minute
		case "h":
			unitDuration = time.Duration(value) * time.Hour
		case "d":
			unitDuration = time.Duration(value) * 24 * time.Hour
		case "w":
			unitDuration = time.Duration(value) * 7 * 24 * time.Hour
		default:
			return 0, fmt.Errorf("неподдерживаемая единица измерения: %s", unit)
		}

		totalDuration += unitDuration
	}

	return totalDuration, nil
}

// FormatDuration converts timeutils.Duration into a readable string
// The output format includes:
// - w - weeks
// - d - days
// - h - hours
// - m - minutes
// - s - seconds
func FormatDuration(duration time.Duration) string {
	if duration <= 0 {
		return "0s"
	}

	seconds := int64(duration.Seconds())
	if seconds == 0 {
		return "0s"
	}

	units := []struct {
		label   string
		seconds int64
	}{
		{"w", 60 * 60 * 24 * 7},
		{"d", 60 * 60 * 24},
		{"h", 60 * 60},
		{"m", 60},
		{"s", 1},
	}

	result := ""
	remaining := seconds

	for _, unit := range units {
		value := remaining / unit.seconds
		if value > 0 {
			result += fmt.Sprintf("%d%s", value, unit.label)
			remaining -= value * unit.seconds
		}
	}

	return result
}
