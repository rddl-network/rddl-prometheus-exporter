package planetmint

func GetActivatedDeviceCount() (count float64, err error) {
	return client.GetActivatedDeviceCount()
}

func GetActiveDeviceCount() (count float64, err error) {
	return client.GetActiveDeviceCount()
}
