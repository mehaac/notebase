// caldav implementation for Mac Calendar compatibility with mock data
package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// setupAuth checks Basic Auth and sends a challenge if missing.
// Returns the username and a bool indicating success.
func setupAuth(e *core.RequestEvent, app *pocketbase.PocketBase) (string, bool) {
	authHeader := e.Request.Header.Get("Authorization")
	if authHeader == "" {
		app.Logger().Debug("Sending auth challenge")
		e.Response.Header().Set("WWW-Authenticate", `Basic realm="CalDAV Server"`)
		e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
		e.Response.Header().Set("Connection", "close")
		e.String(http.StatusUnauthorized, "<html><body><h1>401 Unauthorized</h1><p>Authorization required</p></body></html>")
		return "", false
	}
	username, _, ok := e.Request.BasicAuth()
	if !ok {
		e.NoContent(http.StatusUnauthorized)
		return "", false
	}
	app.Logger().Debug("Auth success", "username", username)
	return username, true
}

// handleOptions responds to OPTIONS requests for CalDAV discovery.
func handleOptions(e *core.RequestEvent, app *pocketbase.PocketBase) error {
	_, ok := setupAuth(e, app)
	if !ok {
		return nil
	}
	e.Response.Header().Set("Allow", "OPTIONS, PROPFIND, REPORT, PUT, DELETE")
	e.Response.Header().Set("DAV", "1, 3, calendar-access")
	return e.NoContent(http.StatusOK)
}

// handlePropfind responds to PROPFIND requests with mock calendar properties.
func handlePropfind(e *core.RequestEvent, app *pocketbase.PocketBase) error {
	username, ok := setupAuth(e, app)
	if !ok {
		return nil
	}
	path := e.Request.PathValue("path")
	app.Logger().Debug("PROPFIND request", "username", username, "path", path, "depth", e.Request.Header.Get("Depth"))
	var xmlResponse string
	if path == "" {
		xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
  <d:response>
    <d:href>/caldav/</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>My CalDAV Calendar</d:displayname>
        <d:resourcetype>
          <d:collection/>
          <cal:calendar/>
        </d:resourcetype>
        <d:current-user-principal>
          <d:href>/caldav/principals/` + username + `/</d:href>
        </d:current-user-principal>
        <d:owner>
          <d:href>/caldav/principals/` + username + `/</d:href>
        </d:owner>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
</d:multistatus>`
	} else {
		xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
  <d:response>
    <d:href>/caldav/` + path + `</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>Event: ` + path + `</d:displayname>
        <d:resourcetype>
          <d:collection/>
          <cal:calendar/>
        </d:resourcetype>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
</d:multistatus>`
	}
	e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
	e.Response.Header().Set("DAV", "1, 3, calendar-access")
	return e.String(http.StatusMultiStatus, xmlResponse)
}

// handleReport responds to REPORT requests with mock event data.
func handleReport(e *core.RequestEvent, app *pocketbase.PocketBase) error {
	username, ok := setupAuth(e, app)
	if !ok {
		return nil
	}
	path := e.Request.PathValue("path")
	app.Logger().Debug("REPORT request", "username", username, "path", path)
	buf := make([]byte, 1024)
	n, err := e.Request.Body.Read(buf)
	if err != nil && err != io.EOF {
		app.Logger().Error("Error reading request body", "error", err)
		return e.NoContent(http.StatusBadRequest)
	}
	body := string(buf[:n])
	var xmlResponse string
	if strings.Contains(body, "calendar-query") {
		xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
  <d:response>
    <d:href>/caldav/event1.ics</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>Mock Event 1</d:displayname>
        <cal:calendar-data>BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Mock Calendar//EN
BEGIN:VEVENT
UID:event1@example.com
SUMMARY:Mock Event 1
DTSTART:20231001T100000Z
DTEND:20231001T110000Z
END:VEVENT
END:VCALENDAR</cal:calendar-data>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
</d:multistatus>`
	} else if strings.Contains(body, "calendar-multiget") {
		xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
  <d:response>
    <d:href>/caldav/event1.ics</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>Mock Event 1</d:displayname>
        <cal:calendar-data>BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Mock Calendar//EN
BEGIN:VEVENT
UID:event1@example.com
SUMMARY:Mock Event 1
DTSTART:20231001T100000Z
DTEND:20231001T110000Z
END:VEVENT
END:VCALENDAR</cal:calendar-data>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
  <d:response>
    <d:href>/caldav/event2.ics</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>Mock Event 2</d:displayname>
        <cal:calendar-data>BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Mock Calendar//EN
BEGIN:VEVENT
UID:event2@example.com
SUMMARY:Mock Event 2
DTSTART:20231002T120000Z
DTEND:20231002T130000Z
END:VEVENT
END:VCALENDAR</cal:calendar-data>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
</d:multistatus>`
	} else {
		xmlResponse = `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cal="urn:ietf:params:xml:ns:caldav">
</d:multistatus>`
	}
	e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
	e.Response.Header().Set("DAV", "1, 3, calendar-access")
	return e.String(http.StatusMultiStatus, xmlResponse)
}

// handlePrincipal responds to PROPFIND requests for principal data.
func handlePrincipal(e *core.RequestEvent, app *pocketbase.PocketBase) error {
	username, ok := setupAuth(e, app)
	if !ok {
		return nil
	}
	xmlResponse := `<?xml version="1.0" encoding="UTF-8"?>
<d:multistatus xmlns:d="DAV:" xmlns:cs="http://calendarserver.org/ns/">
  <d:response>
    <d:href>/caldav/principals/` + username + `/</d:href>
    <d:propstat>
      <d:prop>
        <d:displayname>` + username + `</d:displayname>
        <d:current-user-principal>
          <d:href>/caldav/principals/` + username + `/</d:href>
        </d:current-user-principal>
        <d:principal-URL>
          <d:href>/caldav/principals/` + username + `/</d:href>
        </d:principal-URL>
        <d:principal-collection-set>
          <d:href>/caldav/principals/</d:href>
        </d:principal-collection-set>
        <d:current-user-privilege-set>
          <d:privilege><d:read/></d:privilege>
          <d:privilege><d:write/></d:privilege>
          <d:privilege><d:read-acl/></d:privilege>
          <d:privilege><d:write-acl/></d:privilege>
        </d:current-user-privilege-set>
      </d:prop>
      <d:status>HTTP/1.1 200 OK</d:status>
    </d:propstat>
  </d:response>
</d:multistatus>`
	e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
	e.Response.Header().Set("DAV", "1, 3, calendar-access, binding, supersede")
	return e.String(http.StatusMultiStatus, xmlResponse)
}

// initCaldavRoutes registers CalDAV endpoints.
func initCaldavRoutes(app *pocketbase.PocketBase, se *core.ServeEvent) {
	caldavGroup := se.Router.Group("/caldav")
	// Register one OPTIONS endpoint across all paths.
	caldavGroup.Route("OPTIONS", "/{path...}", func(e *core.RequestEvent) error {
		return handleOptions(e, app)
	})
	caldavGroup.Route("PROPFIND", "/{path...}", func(e *core.RequestEvent) error {
		return handlePropfind(e, app)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)
	caldavGroup.Route("PROPFIND", "/principals/{path...}", func(e *core.RequestEvent) error {
		return handlePrincipal(e, app)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)
	caldavGroup.Route("REPORT", "/{path...}", func(e *core.RequestEvent) error {
		return handleReport(e, app)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	// Add well-known endpoint to support discovery (redirect to /caldav/)
	se.Router.Route("PROPFIND", "/.well-known/caldav", func(e *core.RequestEvent) error {
		e.Response.Header().Set("Location", "/caldav/")
		return e.NoContent(http.StatusFound)
	})
}
