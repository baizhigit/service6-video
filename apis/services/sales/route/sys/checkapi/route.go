package checkapi

import "github.com/baizhigit/service6-video/foundation/web"

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	app.HandleFunc("GET /readiness", readiness)
	app.HandleFunc("GET /liveness", liveness)
}
