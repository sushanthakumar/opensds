// Copyright 2017 The OpenSDS Authors.
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

package routers

import (
	"github.com/astaxie/beego"
	"github.com/sodafoundation/api/pkg/api/controllers"
	"github.com/sodafoundation/api/pkg/utils/constants"
)

func init() {

	// add router for block storage api
	csiByPassns :=
		beego.NewNamespace("/"+constants.APIVersion+"/:tenantId",

			// Volume is the logical description of a piece of storage, which can be directly used by users.
			// All operations of volume can be used for both admin and users.
			beego.NSRouter("/csibypass", controllers.NewCsiByPassPortal(), "post:ExecCsiByPass"),
		)
	beego.AddNamespace(csiByPassns)
}
