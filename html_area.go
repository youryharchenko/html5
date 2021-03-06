package html5

// HTMLArea represents HTML <area> tag
type HTMLArea struct {
	HTMLElement
}

// Area creates an HTML <area> tag element
func Area() *HTMLArea {
	e := &HTMLArea{}
	e.a = make(map[string]interface{})
	e.tagName = "area"
	return e
}

// S sets the element's CSS properties
func (e *HTMLArea) S(style StyleMap) *HTMLArea {
	e.HTMLElement.S(style)
	return e
}

// Key sets virtual dom's special property to instruct the diffing mechanism
// to reorder the node instead of replacing it
func (e *HTMLArea) Key(key interface{}) *HTMLArea {
	e.key = F(key)
	return e
}

// Ref marks the dest pointer to receive the real DOM element on render.
// Useful for getting live value of an input element, for example.
func (e *HTMLArea) Ref(dest *DOMElement) *HTMLArea {
	e.ref = dest
	return e
}

// Alt sets the element's "alt" attribute
func (e *HTMLArea) Alt(v string) *HTMLArea {
	e.a["alt"] = v
	return e
}

// Coords sets the element's "coords" attribute
func (e *HTMLArea) Coords(v string) *HTMLArea {
	e.a["coords"] = v
	return e
}

// Shape sets the element's "shape" attribute
func (e *HTMLArea) Shape(v string) *HTMLArea {
	e.a["shape"] = v
	return e
}

// Target sets the element's "target" attribute
func (e *HTMLArea) Target(v string) *HTMLArea {
	e.a["target"] = v
	return e
}

// Download sets the element's "download" attribute
func (e *HTMLArea) Download(v string) *HTMLArea {
	e.a["download"] = v
	return e
}

// Rel sets the element's "rel" attribute
func (e *HTMLArea) Rel(v string) *HTMLArea {
	e.a["rel"] = v
	return e
}

// ReferrerPolicy sets the element's "referrerpolicy" attribute
func (e *HTMLArea) ReferrerPolicy(v string) *HTMLArea {
	e.a["referrerpolicy"] = v
	return e
}

// ID sets the element's "id" attribute
func (e *HTMLArea) ID(v string) *HTMLArea {
	e.a["id"] = v
	return e
}

// Class sets the element's "class" attribute
func (e *HTMLArea) Class(v string) *HTMLArea {
	e.a["class"] = v
	return e
}

// Title sets the element's "title" attribute
func (e *HTMLArea) Title(v string) *HTMLArea {
	e.a["title"] = v
	return e
}

// Lang sets the element's "lang" attribute
func (e *HTMLArea) Lang(v string) *HTMLArea {
	e.a["lang"] = v
	return e
}

// Translate sets the element's "translate" attribute
func (e *HTMLArea) Translate(v bool) *HTMLArea {
	if v {
		e.a["translate"] = ""
	} else {
		delete(e.a, "translate")
	}
	return e
}

// Dir sets the element's "dir" attribute
func (e *HTMLArea) Dir(v string) *HTMLArea {
	e.a["dir"] = v
	return e
}

// Hidden sets the element's "hidden" attribute
func (e *HTMLArea) Hidden(v bool) *HTMLArea {
	if v {
		e.a["hidden"] = ""
	} else {
		delete(e.a, "hidden")
	}
	return e
}

// TabIndex sets the element's "tabindex" attribute
func (e *HTMLArea) TabIndex(v int) *HTMLArea {
	e.a["tabindex"] = v
	return e
}

// AccessKey sets the element's "accesskey" attribute
func (e *HTMLArea) AccessKey(v string) *HTMLArea {
	e.a["accesskey"] = v
	return e
}

// Draggable sets the element's "draggable" attribute
func (e *HTMLArea) Draggable(v bool) *HTMLArea {
	if v {
		e.a["draggable"] = ""
	} else {
		delete(e.a, "draggable")
	}
	return e
}

// Spellcheck sets the element's "spellcheck" attribute
func (e *HTMLArea) Spellcheck(v bool) *HTMLArea {
	if v {
		e.a["spellcheck"] = ""
	} else {
		delete(e.a, "spellcheck")
	}
	return e
}

// Href sets the element's "href" attribute
func (e *HTMLArea) Href(v string) *HTMLArea {
	e.a["href"] = v
	return e
}

// Protocol sets the element's "protocol" attribute
func (e *HTMLArea) Protocol(v string) *HTMLArea {
	e.a["protocol"] = v
	return e
}

// Username sets the element's "username" attribute
func (e *HTMLArea) Username(v string) *HTMLArea {
	e.a["username"] = v
	return e
}

// Password sets the element's "password" attribute
func (e *HTMLArea) Password(v string) *HTMLArea {
	e.a["password"] = v
	return e
}

// Host sets the element's "host" attribute
func (e *HTMLArea) Host(v string) *HTMLArea {
	e.a["host"] = v
	return e
}

// Hostname sets the element's "hostname" attribute
func (e *HTMLArea) Hostname(v string) *HTMLArea {
	e.a["hostname"] = v
	return e
}

// Port sets the element's "port" attribute
func (e *HTMLArea) Port(v string) *HTMLArea {
	e.a["port"] = v
	return e
}

// Pathname sets the element's "pathname" attribute
func (e *HTMLArea) Pathname(v string) *HTMLArea {
	e.a["pathname"] = v
	return e
}

// Search sets the element's "search" attribute
func (e *HTMLArea) Search(v string) *HTMLArea {
	e.a["search"] = v
	return e
}

// Hash sets the element's "hash" attribute
func (e *HTMLArea) Hash(v string) *HTMLArea {
	e.a["hash"] = v
	return e
}
