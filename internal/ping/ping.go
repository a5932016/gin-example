package ping

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/a5932016/gin_example/internal/ping/pingRequest"
	"github.com/a5932016/gin_example/pkg/layout"
)

func Start(request pingRequest.Start, keepConnectModel *layout.KeepConnectModel) {
	defer keepConnectModel.Cancel()

	cmdList := []string{request.IP, "-c", strconv.Itoa(request.Target)}
	cmd := exec.Command("ping", cmdList...)

	read, err := cmd.StdoutPipe()
	if err != nil {
		keepConnectModel.Err <- err
		return
	}

	err = cmd.Start()
	if err != nil {
		keepConnectModel.Err <- err
		return
	}

	buffer := make([]byte, 1024)
	for {
		select {
		case <-keepConnectModel.Ctx.Done():
			return
		case err := <-keepConnectModel.Err:
			fmt.Printf("error: %v\n", err)
			return
		default:
			n, err := read.Read(buffer)
			if err != nil {
				return
			}

			keepConnectModel.Output <- buffer[:n]
		}
	}
}
