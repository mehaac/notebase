// caldav is not even close to being ready, this is mostly boilerplate LLM code, which doesn't work with Mac calendar
package main

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func initCaldavRoutes(app *pocketbase.PocketBase, se *core.ServeEvent) {
	caldavGroup := se.Router.Group("/caldav")

	caldavGroup.Route("PROPFIND", "/{path...}", func(e *core.RequestEvent) error {
		// Check for auth header
		authHeader := e.Request.Header.Get("Authorization")
		if authHeader == "" {
			app.Logger().Debug("Sending auth challenge")

			// Set exact headers used by CalDAV servers
			e.Response.Header().Set("WWW-Authenticate", `Basic realm="CalDAV Server"`)
			e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
			e.Response.Header().Set("Connection", "close")

			// Return 401 with a body (some clients expect this)
			return e.String(http.StatusUnauthorized, "<html><body><h1>401 Unauthorized</h1><p>Authorization required</p></body></html>")
		}

		// Process authenticated requests...
		username, _, ok := e.Request.BasicAuth()
		if !ok {
			return e.NoContent(http.StatusUnauthorized)
		}

		app.Logger().Debug("Auth success", "username", username)

		path := e.Request.PathValue("path")
		app.Logger().Debug("PROPFIND request", "path", path, "depth", e.Request.Header.Get("Depth"))

		xmlResponse := ""
		if path == "" {
			xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
    <d:response>
        <d:href>/caldav/</d:href>
        <d:propstat>
            <d:prop>
                <d:displayname>My CalDAV Server</d:displayname>
                <d:resourcetype>
                    <d:collection/>
                    <cal:calendar/>
                </d:resourcetype>
                <d:current-user-principal>
                    <d:href>/caldav/principals/user/</d:href>
                </d:current-user-principal>
                <d:principal-collection-set>
                    <d:href>/caldav/principals/</d:href>
                </d:principal-collection-set>
                <d:owner>
                    <d:href>/caldav/principals/user/</d:href>
                </d:owner>
                <d:supported-report-set>
                    <d:report>
                        <d:calendar-multiget/>
                    </d:report>
                    <d:report>
                        <d:calendar-query/>
                    </d:report>
                </d:supported-report-set>
            </d:prop>
            <d:status>HTTP/1.1 200 OK</d:status>
        </d:propstat>
    </d:response>
</d:multistatus>`
		}

		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		e.Response.Header().Set("DAV", "1, 3, calendar-access")
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

}
