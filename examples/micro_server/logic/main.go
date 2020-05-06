// MIT License

// Copyright (c) 2019 gonethopper

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// * @Author: ankye
// * @Date: 2019-06-14 19:56:49
// * @Last Modified by:   ankye
// * @Last Modified time: 2019-06-14 19:56:49

package main

import (

	//"github.com/gonethopper/nethopper/cache/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gonethopper/nethopper/examples/micro_server/logic/docs"
	"github.com/gonethopper/nethopper/examples/micro_server/logic/modules/gclient"
	"github.com/gonethopper/nethopper/examples/micro_server/logic/modules/grpc"
	"github.com/gonethopper/nethopper/examples/micro_server/logic/modules/logic"
	"github.com/gonethopper/nethopper/examples/micro_server/logic/modules/redis"
	"github.com/gonethopper/nethopper/log"
	. "github.com/gonethopper/nethopper/server"
)

// @title Nethopper Micro Server - Logic
// @version 1.0.2
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:11080
// @BasePath
func main() {

	//runtime.GOMAXPROCS(1)
	m := map[string]interface{}{
		"filename":       "logs/server.log",
		"level":          DEBUG,
		"maxSize":        5000,
		"maxLines":       10000000,
		"hourEnabled":    false,
		"dailyEnable":    true,
		"queueSize":      1000,
		"driver":         "mysql",
		"dsn":            "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		"address":        ":14000",
		"grpcServerID_0": 0,
		"grpcAddress_0":  ":14001",
		"grpcName_0":     "grpcclient_0",
	}
	RegisterModule("log", log.LogModuleCreate)
	RegisterModule("logic", logic.ModuleCreate)
	RegisterModule("redis", redis.ModuleCreate)
	RegisterModule("grpc", grpc.ModuleCreate)
	RegisterModule("gclient", gclient.ModuleCreate)
	NewNamedModule(ModuleIDLog, "log", nil, m)
	NewNamedModule(ModuleIDRedis, "redis", nil, m)
	NewNamedModule(ModuleIDLogic, "logic", nil, m)
	NewNamedModule(ModuleIDGRPCServer, "grpc", nil, m)
	NewNamedModule(ModuleIDGRPCClient, "gclient", nil, m)
	InitSignal()
	//GracefulExit()
}
