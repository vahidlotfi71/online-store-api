package Providers

func ErrorProvider(err error) string {
	message := err.Error()

	return message
}
