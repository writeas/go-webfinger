[![GoDoc](https://godoc.org/code.as/writeas/go-webfinger?status.svg)](https://godoc.org/code.as/writeas/go-webfinger)

# go-webfinger

go-webfinger is a golang webfinger server implementation. See [v1.0](https://github.com/writeas/go-webfinger/releases/tag/1.0) for the latest stable version, and our [Code.as repo](https://code.as/writeas/go-webfinger) for the Write.as-specific implementation.

Past v1.0, this fork was made especially for federation support on [Write.as](https://write.as), which includes users across write.as, \*.writeas.com, and custom domains we host. The `master` branch contains changes specific to our implementation, and will change without notification.

## Status

Project has had no input/work by me (sheenobu) in the past 2 years. Marking
the package as 1.0 to match [github.com/writeas/go-webfinger](https://github.com/writeas/go-webfinger).

## Usage

`webfinger.Service` is implemented as a net/http handler, which means
usage is simply registering the object with your http service.

Using the webfinger service as the main ServeHTTP:

```go
myResolver = ...
wf := webfinger.Default(myResolver{})
wf.NotFoundHandler = // the rest of your app
http.ListenAndService(":8080", wf)
```

Using the webfinger service as a route on an existing HTTP router:

```go
myResolver = ...
wf := webfinger.Default(myResolver{})
http.Handle(webfinger.WebFingerPath, http.HandlerFunc(wf.Webfinger))
http.ListenAndService(":8080", nil)
```

## Defaults

The webfinger service is installed with a few defaults. Some of these
defaults ensure we stick closely to the webfinger specification (tls-only, CORS, Content-Type)
and other defaults are simply useful for development (no-cache)

The full list of defaults can be found in the godoc for `webfinger.Service`. They are exposed
as public variables which can be overriden.

`PreHandlers` are the list of preflight HTTP handlers to run. You can add your own via `wf.PreHandlers["my-custom-name"] = ...`, however,
execution order is not guaranteed.

### TLS-Only

Handler which routes to the TLS version of the page. Disable via `wf.NoTLSHandler = nil`.

### No-Cache

A PreFlight handler which sets no-cache headers on anything under `/.well-known/webfinger`. Disable or override via `wf.PreHandlers[webfinger.NoCacheMiddleware]`

### Content Type as application/jrd+json

A PreFlight handler which sets the Content-Type to `application/jrd+json`. Disable or override via `wf.PreHandlers[webfinger.ContentTypeMiddleware]`.

### CORS 

A PreFlight handler which adds the CORS headers. Disable or override via `wf.PreHandlers[webfinger.CorsMiddleware].`
