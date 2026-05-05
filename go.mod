module github.com/go-gst/go-gst

go 1.24.0

toolchain go1.24.4

require (
	github.com/go-gst/go-glib v0.0.0-00010101000000-000000000000
	github.com/go-gst/go-pointer v0.0.0-20241127163939-ba766f075b4c
)

require golang.org/x/exp v0.0.0-20251009144603-d2f985daa21b // indirect

replace github.com/go-gst/go-glib => github.com/vopenia-io/go-glib v0.0.0-20260505152650-eeba9b99a44e
