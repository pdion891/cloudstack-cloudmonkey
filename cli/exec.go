// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cli

import (
	"cmk/cmd"
	"cmk/config"
)

// ExecCmd executes a single provided command
func ExecCmd(cfg *config.Config, args []string) error {
	if len(args) < 1 {
		return nil
	}

	command := cmd.FindCommand(args[0])
	if command != nil {
		return command.Handle(cmd.NewRequest(command, cfg, args[1:]))
	}

	catchAllHandler := cmd.GetAPIHandler()
	return catchAllHandler.Handle(cmd.NewRequest(catchAllHandler, cfg, args))
}
