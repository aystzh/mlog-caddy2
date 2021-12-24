// Copyright 2021 ZLIU.ORG
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

package zauth

import (
	"net/http"
	"sync"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"

	//"github.com/juju/errors"
	"github.com/liuzl/store"
	//"github.com/rs/zerolog"
)

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective("mlog", parseCaddyfile)
}

var once, onceAdmin sync.Once
var authDB *store.LevelStore

// Middleware implements an HTTP handler that implements the
// ak, sk auth.
type Middleware struct {
	AuthDBDir     string `json:"log_dir,omitempty"`
	AuthAdminAddr string `json:"split_by,omitempty`
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mlog",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Middleware) Provision(ctx caddy.Context) error {
	if m.AuthDBDir == "" {
		m.AuthDBDir = "./zhb_search_log"
	}
	if m.AuthAdminAddr == "" {
		m.AuthAdminAddr = "day"
	}
	//TODO web api
	/*onceAdmin.Do(func() {
		go m.admin()
	}) */
	return nil
}

// Validate implements caddy.Validator.
func (m *Middleware) Validate() error {
	return nil
}

func (m *Middleware) getAuthDB() *store.LevelStore {
	once.Do(func() {
		var err error
		if authDB, err = store.NewLevelStore(m.AuthDBDir); err != nil {
			panic(err)
		}
	})
	return authDB
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	caddy.Log().Info("msg", zap.String("msg", "ServeHTTP"))

	m.authorize(r)

	/* if pass, err := m.authorize(r); !pass {
		if err != nil {
			ctx := r.Context()
			zlog := zerolog.Ctx(ctx)
			hashlog := GetZlog(ctx)
			zlog.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("error", errors.ErrorStack(err))
			})
			hashlog.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("error", errors.ErrorStack(err))
			})
		}
		ret := MakeResponse(err, "", nil)
		MustEncode(w, ret)
		return nil
	} */
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "log_dir":
				if d.NextArg() {
					caddy.Log().Info("msg", zap.String("log_dir", d.Val()))
					m.AuthDBDir = d.Val()
				}
			case "split_by":
				if d.NextArg() {
					caddy.Log().Info("msg", zap.String("split_by", d.Val()))
					m.AuthAdminAddr = d.Val()
				}
			}
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
