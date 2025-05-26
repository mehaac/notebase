package notebasesync

import (
	"github.com/syncthing/notify"
)

func (h *SyncHandler) WatcherManager() {
	var internalStopCh chan struct{}
	var doneCh chan struct{}

	defer notify.Stop(h.fileChanges)
	go h.fileWatcher(h.fileChanges, h.excludePatters)

	for {
		select {
		case restart, ok := <-h.controlCh:
			if !ok {
				// Channel closed, stop sync job if running and return
				if internalStopCh != nil {
					close(internalStopCh)
					<-doneCh
				}
				return
			}

			if restart {
				h.InitialSync()
			} else {
				if internalStopCh != nil {
					close(internalStopCh)
					<-doneCh
					internalStopCh = nil
					doneCh = nil
				}
			}
		}
	}
}
