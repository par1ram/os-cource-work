module github.com/par1ram/client

go 1.23.2

require (
	github.com/par1ram/common v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

replace github.com/par1ram/common => ../common
