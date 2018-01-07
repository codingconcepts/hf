package command

import (
	"context"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

const (
	// HostFile is the full path to the host file.
	HostFile = `%windir%\system32\drivers\etc\hosts`
}

// FlushDNS refreshes the local DNS cache and is
// useful after a host file modification.
func FlushDNS(c *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "killall", "-HUP", "mDNSResponder")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("running command: %v", err)
	}
}
