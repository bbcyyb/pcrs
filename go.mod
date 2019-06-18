module pcrs

go 1.12

require (
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	golang.org/x/crypto v0.0.0-20190617133340-57b3e21c3d56
	golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f
)

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
