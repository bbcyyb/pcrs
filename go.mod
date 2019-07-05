module pcrs

go 1.12

require (
	github.com/bbcyyb/pcrs v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.4.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/golang-migrate/migrate v0.0.0-20190527172536-8437fe6dc6ae
	golang.org/x/crypto v0.0.0-20190617133340-57b3e21c3d56
	golang.org/x/sys v0.0.0-20190616124812-15dcb6c0061f

)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0

	github.com/bbcyyb/pcrs => ./
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190618222545-ea8f1a30c443
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190618124811-92942e4437e2
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20190619014844-b5b0513f8c1b
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190619223125-e40ef342dc56
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190619215442-4adf7a708c2d
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190611190212-a7e196e89fd3
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
)
