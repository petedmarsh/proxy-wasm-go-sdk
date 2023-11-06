// Copyright 2020-2021 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const (
	bufferOperationAppend  = "append"
	bufferOperationPrepend = "prepend"
	bufferOperationReplace = "replace"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &logContext{}
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	return types.OnPluginStartStatusOK
}

type logContext struct {
	// Embed the default root http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
}

// Override types.DefaultHttpContext.
func (ctx *logContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfo("OnHttpRequestHeaders")
	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *logContext) OnHttpRequestBody(bodySize int, endOfStream bool) types.Action {
	proxywasm.LogInfo("(setBody) OnHttpRequestBody")
	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *logContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	proxywasm.LogInfo("OnHttpResponseHeaders")
	return types.ActionContinue
}

// Override types.DefaultHttpContext.
func (ctx *logContext) OnHttpResponseBody(bodySize int, endOfStream bool) types.Action {
	proxywasm.LogInfo("OnHttpResponseBody")
	return types.ActionContinue
}
