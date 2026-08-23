package enum

type DriverStatus string

const (
	StatusOnline  DriverStatus = "online"
	StatusOffline DriverStatus = "offline"
)

func (s DriverStatus) IsValid() bool {
	switch s {
	case StatusOnline, StatusOffline:
		return true
	}
	return false
}
