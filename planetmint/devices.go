package planetmint

func GetActivatedDeviceCount() (count float64, err error) {
	return GetClient().GetActivatedDeviceCount()
}

func GetActiveDeviceCount() (count float64, err error) {
	return GetClient().GetActiveDeviceCount()
}
