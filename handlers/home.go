package handlers

import (
	"net/http"
)

// Home to display available endpoints
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Azure Api Gateway</title><style>*,html,body{box-sizing:border-box;font-family:'Segoe UI',Tahoma,Geneva,Verdana,sans-serif}body{display:flex;align-items:center;justify-content:center;min-height:100vh;background-color:#eee}.wrapper{background-color:#fff;padding:35px 50px;border-radius:20px}code{background-color:#eee;padding:7px 10px;margin-bottom:15px;display:block}h1{color:#1167e7;margin-top:0;text-align:center}p{color:#555;margin:0}small{color:#b66027;display:block;margin-bottom:5px}</style></head><body><div class="wrapper"><h1>Welcome to azure API Gateway</h1><p>These are available endpoints which you can use</p><br/><p><small>To get user details, [send Username header]</small></p><code>https://azureapigateway.azurewebsites.net/user/profile</code><p><small>To get Micro-service name</small></p><code>https://azureapigateway.azurewebsites.net/microservice/name</code></div></body></html>`))
	w.WriteHeader(http.StatusOK)
}
