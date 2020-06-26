// Copyright 2019 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
This module implements a entry into the OpenSDS northbound service.

*/

package controllers

import (
	"encoding/json"
	"fmt"
	log "github.com/golang/glog"
	"github.com/sodafoundation/api/pkg/api/controllerclient"
	"github.com/sodafoundation/api/pkg/api/policy"
	"github.com/sodafoundation/api/pkg/model"
	csi "github.com/sodafoundation/api/pkg/model/csi"
	//"context"
	//c "github.com/sodafoundation/api/pkg/context"
)

func NewCsiByPassPortal() *CsiByPassPortal {
	return &CsiByPassPortal{
		CtrClient: client.NewClient(),
	}
}

type CsiVolumeBuilder struct {
	csiFunc string
	csiModel *csi.CreateVolumeRequest
}

type CsiByPassPortal struct {
	BasePortal

	CtrClient client.Client
}

func (v *CsiByPassPortal) ExecCsiByPass() {
	fmt.Println("[CSIBYPASS] ExecCsiByPass: %+v\", req")

	if !policy.Authorize(v.Ctx, "volume:create") {
		return
	}
	//ctx := c.GetContext(v.Ctx)
	var csiVolume csi.CreateVolumeRequest
	// Unmarshal the request body
	if err := json.NewDecoder(v.Ctx.Request.Body).Decode(&csiVolume); err != nil {
		errMsg := fmt.Sprintf("[CSI]parse volume request body failed: %s", err.Error())
		v.ErrorHandle(model.ErrorBadRequest, errMsg)
		return
	}
	fmt.Println("csiVolume = %v", csiVolume)

	// NOTE:The real volume creation process.
	// Volume creation request is sent to the Dock. Dock will update volume status to "available"
	// after volume creation is completed.
	if err := v.CtrClient.Connect("0.0.0.0:50051"); err != nil {
		log.Error("when connecting dock client:", err)
		return
	}
/*
	csiOpt := &csiVolume

	response, err := v.CtrClient.CreateVolume(context.Background(), csiOpt)
	if err != nil {
		log.Error("create volume failed in Dock service:", err)
		return
	}
 */
	//log.Info(response)

	return
}
