module github.com/neoaggelos/juju-go-example

go 1.15

// Juju dependencies
replace github.com/altoros/gosigma => github.com/juju/gosigma v0.0.0-20200420012028-063911838a9e

replace gopkg.in/natefinch/lumberjack.v2 => github.com/juju/lumberjack v2.0.0-20200420012306-ddfd864a6ade+incompatible

replace gopkg.in/mgo.v2 => github.com/juju/mgo v2.0.0-20201106044211-d9585b0b0d28+incompatible

replace github.com/hashicorp/raft => github.com/juju/raft v2.0.0-20200420012049-88ad3b3f0a54+incompatible

replace gopkg.in/yaml.v2 => github.com/juju/yaml v0.0.0-20200420012109-12a32b78de07

replace github.com/dustin/go-humanize v1.0.0 => github.com/dustin/go-humanize v0.0.0-20141228071148-145fabdb1ab7

replace github.com/hashicorp/raft-boltdb => github.com/juju/raft-boltdb v0.0.0-20200518034108-40b112c917c5

replace (
	k8s.io/api v0.0.0 => k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.0.0 => k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.0.0 => k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.0.0 => k8s.io/client-go v0.18.6
)

require (
	github.com/juju/charm/v8 v8.0.0-20200925053015-07d39c0154ac // indirect
	github.com/juju/juju v0.0.0-20201113145451-d16ca94d580c
	github.com/juju/names/v4 v4.0.0-20200929085019-be23e191fee0
	github.com/juju/persistent-cookiejar v0.0.0-20171026135701-d5e5a8405ef9 // indirect
	github.com/juju/usso v1.0.1 // indirect
	github.com/juju/webbrowser v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.7 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b // indirect
	golang.org/x/sys v0.0.0-20201113135734-0a15ea8d9b02 // indirect
	golang.org/x/text v0.3.4 // indirect
	gopkg.in/macaroon-bakery.v2 v2.2.0 // indirect
	gopkg.in/retry.v1 v1.0.3 // indirect
)
