// +build darwin
// +build !cgo

package host

import (
	"context"

	"github.com/Hellyna/gopsutil/v3/internal/common"
)

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
