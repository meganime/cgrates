/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package apis

import (
	"path"
	"time"

	"github.com/cgrates/birpc/context"
	"github.com/cgrates/cgrates/cores"
	"github.com/cgrates/cgrates/utils"
)

func NewCoreSv1(cS *cores.CoreService) *CoreSv1 {
	return &CoreSv1{cS: cS}
}

// CoreSv1 exports RPC from RLs
type CoreSv1 struct {
	cS *cores.CoreService
	ping
}

func (cS *CoreSv1) Status(_ *context.Context, arg *utils.TenantWithAPIOpts, reply *map[string]interface{}) error {
	return cS.cS.Status(arg, reply)
}

// Sleep is used to test the concurrent requests mechanism
func (cS *CoreSv1) Sleep(_ *context.Context, arg *utils.DurationArgs, reply *string) error {
	time.Sleep(arg.Duration)
	*reply = utils.OK
	return nil
}

func (cS *CoreSv1) Shutdown(_ *context.Context, _ *utils.CGREvent, reply *string) error {
	cS.cS.ShutdownEngine()
	*reply = utils.OK
	return nil
}

// StartCPUProfiling is used to start CPUProfiling in the given path
func (cS *CoreSv1) StartCPUProfiling(_ *context.Context, args *utils.DirectoryArgs, reply *string) error {
	if err := cS.cS.StartCPUProfiling(path.Join(args.DirPath, utils.CpuPathCgr)); err != nil {
		return err
	}
	*reply = utils.OK
	return nil
}

// StopCPUProfiling is used to stop CPUProfiling. The file should be written on the path
// where the CPUProfiling already started
func (cS *CoreSv1) StopCPUProfiling(_ *context.Context, _ *utils.TenantWithAPIOpts, reply *string) error {
	if err := cS.cS.StopCPUProfiling(); err != nil {
		return err
	}
	*reply = utils.OK
	return nil
}

// StartMemoryProfiling is used to start MemoryProfiling in the given path
func (cS *CoreSv1) StartMemoryProfiling(_ *context.Context, args *utils.MemoryPrf, reply *string) error {
	if err := cS.cS.StartMemoryProfiling(args); err != nil {
		return err
	}
	*reply = utils.OK
	return nil
}

// StopMemoryProfiling is used to stop MemoryProfiling. The file should be written on the path
// where the MemoryProfiling already started
func (cS *CoreSv1) StopMemoryProfiling(_ *context.Context, _ *utils.TenantWithAPIOpts, reply *string) error {
	if err := cS.cS.StopMemoryProfiling(); err != nil {
		return err
	}
	*reply = utils.OK
	return nil
}