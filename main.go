package main

import "final-project/handler"

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath                    /api
func main() {
	handler.StartApp()
}
