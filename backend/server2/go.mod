module github.com/par1ram/server2

go 1.23.2

replace github.com/par1ram/common => ../common

require (
	github.com/par1ram/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
