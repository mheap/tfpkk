// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Emphasis struct {
	// List of children nodes of the current node
	Children []AnyNode `json:"children,omitempty"`
	// Level of emphasis
	// 1 - italic
	// 2 - bold
	//
	Level *float64 `json:"level,omitempty"`
	Type  string   `json:"type"`
}
