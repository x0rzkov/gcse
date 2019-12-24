#!/usr/bin/env gosl

import "time"
import "github.com/x0rzkov/gcse/configs"

Printfln("Logging to %q...", configs.LogDir)

for {
  Bash("stored -log_dir %s", configs.LogDir)
  time.Sleep(time.Second)
}
