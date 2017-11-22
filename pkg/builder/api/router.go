// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
package api

import (
	"github.com/go-chi/chi"
	"github.com/goodrain/rainbond/pkg/builder/api/controller"
)

func APIServer() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/v2/builder", func(r chi.Router) {
		//r.Get("/ping", controller.Ping)
		r.Route("/codecheck", func(r chi.Router) {
			r.Post("/", controller.AddCodeCheck)
			r.Put("/service/{serviceID}", controller.Update)
			r.Get("/service/{serviceID}", controller.GetCodeCheck)
		})
		r.Route("/publish", func(r chi.Router) {
			r.Get("/service/{serviceKey}/version/{appVersion}",controller.GetAppPublish)
			r.Post("/",controller.AddAppPublish)
		})
	})
	return r
}

