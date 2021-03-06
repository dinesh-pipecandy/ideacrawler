// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package internal

// PageResourceType Resource type as it was perceived by the rendering engine.
//
// This type cannot be used directly. Use page.ResourceType instead.
type PageResourceType string

func (e PageResourceType) Valid() bool {
	switch e {
	case "Document", "Stylesheet", "Image", "Media", "Font", "Script", "TextTrack", "XHR", "Fetch", "EventSource", "WebSocket", "Manifest", "Other":
		return true
	default:
		return false
	}
}

func (e PageResourceType) String() string {
	return string(e)
}

// PageFrameID Unique frame identifier.
//
// This type cannot be used directly. Use page.FrameID instead.
type PageFrameID string
