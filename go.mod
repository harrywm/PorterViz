module porter/donut

go 1.15

replace porter/donut/networks v0.0.1 => ./networks

replace porter/donut/gui v0.0.1 => ./gui

require (
	fyne.io/fyne v1.4.3 // indirect
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.1+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/jroimartin/gocui v0.4.0 // indirect
	github.com/kr/pretty v0.2.1
	github.com/nsf/termbox-go v0.0.0-20201124104050-ed494de23a00 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	google.golang.org/grpc v1.34.0 // indirect
	porter/donut/gui v0.0.1
	porter/donut/networks v0.0.1
)
