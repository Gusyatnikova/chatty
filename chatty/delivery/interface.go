package delivery

type ChattyServer interface {
	Run()
	Shutdown()
	ListenForShutdown()
}
