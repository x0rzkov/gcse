#!/usr/bin/env gosl

import "time"
import "github.com/x0rzkov/gcse/configs"

Printfln("Logging to %q...", configs.LogDir)

for {
  Bash("web -log_dir %s", configs.LogDir)
  time.Sleep(time.Second)
}
