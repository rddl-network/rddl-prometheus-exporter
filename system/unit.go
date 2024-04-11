package system

import (
	"context"
	"log"

	"github.com/coreos/go-systemd/v22/dbus"
)

// CheckUnitActiveState takes the (unescaped) unit name and returns 1 if the unit is active or 0 if inactive
func CheckUnitActiveState(ctx context.Context, unit string) float64 {
	logger := log.Default()

	conn, err := dbus.NewWithContext(ctx)
	if err != nil {
		logger.Print(err)
		return 0
	}

	unitProps, err := conn.GetUnitPropertiesContext(ctx, unit)
	if err != nil {
		logger.Print(err)
		return 0
	}

	state, ok := unitProps["ActiveState"]
	if !ok {
		logger.Printf("%s properties do not contain ActiveState", unit)
	}

	if state != "active" {
		return 0
	}

	return 1
}
