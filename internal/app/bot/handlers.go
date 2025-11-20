package bot

type Handlers interface {
	HandleMessages(evt any)
}
