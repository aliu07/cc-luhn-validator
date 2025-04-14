package constants

type DataSource int

const (
	Handler DataSource = iota
	Cache
	Server
)

func (d DataSource) String() string {
	switch d {
	case Handler:
		return "handler"
	case Cache:
		return "cache"
	case Server:
		return "server"
	default:
		return "unkown"
	}
}
