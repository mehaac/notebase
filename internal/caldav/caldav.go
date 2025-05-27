package caldav

import (
	"log"
	"net/http"
	"time"

	"github.com/biozz/wow/notebase/internal/config"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type CaldavHandler struct {
	app  *pocketbase.PocketBase
	root string
}

func NewHandler(app *pocketbase.PocketBase, root string, config *config.NotebaseConfig) *CaldavHandler {
	return &CaldavHandler{
		app:  app,
		root: root,
	}
}

// initCaldavRoutes registers minimal CalDAV endpoints.
func (h *CaldavHandler) Routes(se *core.ServeEvent) {
	// Handle all common discovery paths
	discoveryPaths := []string{
		"/.well-known/caldav",
		"/",
		"/principals/",
		"/calendar/dav/",
	}

	for _, path := range discoveryPaths {
		se.Router.Route("PROPFIND", path, func(e *core.RequestEvent) error {
			log.Printf("CalDAV Discovery Request: %s %s", e.Request.Method, e.Request.URL.Path)
			log.Printf("Request Headers: %v", e.Request.Header)

			xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:"><D:response><D:href>/</D:href><D:propstat><D:prop><D:current-user-principal><D:href>/caldav/principals/me/</D:href></D:current-user-principal><D:resourcetype><D:collection/></D:resourcetype></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

			e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
			e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
			log.Printf("Response Headers: %v", e.Response.Header())
			return e.String(http.StatusMultiStatus, xmlResponse)
		}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)
	}

	// Add .well-known/caldav redirect
	se.Router.Route("GET", "/.well-known/caldav", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		e.Response.Header().Set("Location", "/caldav/")
		return e.NoContent(http.StatusPermanentRedirect)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	caldav := se.Router.Group("/caldav")

	// Root PROPFIND - Discovery endpoint
	caldav.Route("PROPFIND", "/", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		log.Printf("Request Headers: %v", e.Request.Header)

		xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:"><D:response><D:href>/</D:href><D:propstat><D:prop><D:current-user-principal><D:href>/caldav/principals/me/</D:href></D:current-user-principal><D:resourcetype><D:collection/></D:resourcetype></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

		e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		log.Printf("Response Headers: %v", e.Response.Header())
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	// Principal PROPFIND
	caldav.Route("PROPFIND", "/principals/me/", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		log.Printf("Request Headers: %v", e.Request.Header)

		xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav"><D:response><D:href>/caldav/principals/me/</D:href><D:propstat><D:prop><C:calendar-home-set><D:href>/caldav/calendars/</D:href></C:calendar-home-set><D:resourcetype><D:principal/></D:resourcetype></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

		e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		log.Printf("Response Headers: %v", e.Response.Header())
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	// Calendar home PROPFIND
	caldav.Route("PROPFIND", "/calendars/", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		log.Printf("Request Headers: %v", e.Request.Header)

		xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav"><D:response><D:href>/caldav/calendars/</D:href><D:propstat><D:prop><D:resourcetype><D:collection/></D:resourcetype><D:displayname>Calendars</D:displayname></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

		e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		log.Printf("Response Headers: %v", e.Response.Header())
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	// Calendar collection PROPFIND
	caldav.Route("PROPFIND", "/calendars/default/", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		log.Printf("Request Headers: %v", e.Request.Header)

		xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav"><D:response><D:href>/caldav/calendars/default/</D:href><D:propstat><D:prop><D:resourcetype><D:collection/><C:calendar/></D:resourcetype><D:displayname>Default Calendar</D:displayname><C:calendar-description>Default calendar</C:calendar-description><C:calendar-timezone>UTC</C:calendar-timezone><C:supported-calendar-component-set><C:comp name="VEVENT"/></C:supported-calendar-component-set></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

		e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		log.Printf("Response Headers: %v", e.Response.Header())
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)

	// Calendar REPORT
	caldav.Route("REPORT", "/calendars/default/", func(e *core.RequestEvent) error {
		log.Printf("CalDAV Request: %s %s", e.Request.Method, e.Request.URL.Path)
		log.Printf("Request Headers: %v", e.Request.Header)

		now := time.Now()
		xmlResponse := `<?xml version="1.0" encoding="utf-8"?><D:multistatus xmlns:D="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav"><D:response><D:href>/caldav/calendars/default/event1.ics</D:href><D:propstat><D:prop><D:getetag>"1"</D:getetag><C:calendar-data>BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//Example Corp.//CalDAV Client//EN
BEGIN:VEVENT
UID:event1@example.com
DTSTAMP:` + now.Format("20060102T150405Z") + `
DTSTART:` + now.Format("20060102T150405Z") + `
DTEND:` + now.Add(time.Hour).Format("20060102T150405Z") + `
SUMMARY:Test Event
DESCRIPTION:This is a test event
END:VEVENT
END:VCALENDAR</C:calendar-data></D:prop><D:status>HTTP/1.1 200 OK</D:status></D:propstat></D:response></D:multistatus>`

		e.Response.Header().Set("DAV", "1, 2, 3, calendar-access")
		e.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
		log.Printf("Response Headers: %v", e.Response.Header())
		return e.String(http.StatusMultiStatus, xmlResponse)
	}).Unbind(apis.DefaultLoadAuthTokenMiddlewareId)
}
