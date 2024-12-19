package h

import (
	"encoding/json"
	"strconv"
)

const (
	//the default, puts the content inside the target element
	HxSwapInnerHtml = "innerHTML"
	//replaces the entire target element with the returned content
	HxSwapOuterHtml = "outerHTML"
	//prepends the content before the first child inside the target
	HxSwapAfterBegin = "afterbegin"
	//prepends the content before the target in the target’s parent element
	HxSwapBeforeBegin = "beforebegin"
	//appends the content after the last child inside the target
	HxSwapBeforeEnd = "beforeend"
	//appends the content after the target in the target’s parent element
	HxSwapAfterEnd = "afterend"
	//deletes the target element regardless of the response
	HxSwapDelete = "delete"
	//does not append content from response (Out of Band Swaps and Response Headers will still be processed)
	HxSwapNone = "none"
)

// controls how content will swap in (outerHTML, beforeend, afterend, …)
func (t *TagBuilder) HxSwap(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-swap", string(value)))
	return t
}

// specifies the event that triggers the request
func (t *TagBuilder) HxTrigger(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-trigger", value))
	return t
}

// issues a GET to the specified URL
func (t *TagBuilder) HxGet(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-get", value))
	return t
}

// issues a POST to the specified URL
func (t *TagBuilder) HxPost(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-post", value))
	return t
}

// issues a PATCH to the specified URL
func (t *TagBuilder) HxPatch(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-patch", value))
	return t
}

// issues a DELETE to the specified URL
func (t *TagBuilder) HxDelete(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-delete", value))
	return t
}

// the element to put the htmx-request class on during the request
func (t *TagBuilder) HxIndicator(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-indicator", value))
	return t
}

// specifies the target element to be swapped
func (t *TagBuilder) HxTarget(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-target", value))
	return t
}

// mark element to swap in from a response (out of band)
func (t *TagBuilder) HxSwapOob(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-swap-oob", strconv.FormatBool(value)))
	return t
}

// shows a confirm() dialog before issuing a request
func (t *TagBuilder) HxConfirm(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-confirm", value))
	return t
}

// add progressive enhancement for links and forms
func (t *TagBuilder) HxBoost(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-boost", strconv.FormatBool(value)))
	return t
}

// extensions to use for this element
func (t *TagBuilder) HxExt(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-ext", value))
	return t
}

// handle events with inline scripts on elements
func (t *TagBuilder) HxOn(event string, value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-on:"+event, value))
	return t
}

// select content to swap in from a response
func (t *TagBuilder) HxSelect(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-select", value))
	return t
}

// add values to submit with the request (JSON format)
func (t *TagBuilder) HxVals(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-vals", value))
	return t
}

// add values to submit with the request (JSON format)
func (t *TagBuilder) HxValsJson(obj interface{}) *TagBuilder {
	value, _ := json.Marshal(obj)
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-vals", string(value)))
	return t
}

// select content to swap in from a response, somewhere other than the target (out of band)
func (t *TagBuilder) HxSelectOob(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-select-oob", strconv.FormatBool(value)))
	return t
}

// issues a DELETE to the specified URL
func (t *TagBuilder) HxDisable(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disable", strconv.FormatBool(value)))
	return t
}

// adds the disabled attribute to the specified elements while a request is in flight
func (t *TagBuilder) HxDisabledElt(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disabled-elt", value))
	return t
}

// control and disable automatic attribute inheritance for child nodes
func (t *TagBuilder) HxDisinherit(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disinherit", strconv.FormatBool(value)))
	return t
}

// changes the request encoding type
func (t *TagBuilder) HxEncoding(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-encoding", value))
	return t
}

// adds to the headers that will be submitted with the request
func (t *TagBuilder) HxHeaders(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-headers", value))
	return t
}

// include additional data in requests
func (t *TagBuilder) HxInlude(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-include", value))
	return t
}

// control and enable automatic attribute inheritance for child nodes if it has been disabled by default
func (t *TagBuilder) HxInherit(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-inherit", strconv.FormatBool(value)))
	return t
}

// filters the parameters that will be submitted with a request
func (t *TagBuilder) HxParams(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-params", value))
	return t
}

// shows a prompt() before submitting a request
func (t *TagBuilder) HxPrompt(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-prompt", value))
	return t
}

// specifies elements to keep unchanged between requests
func (t *TagBuilder) HxPreserve(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-preserve", value))
	return t
}

// replace the URL in the browser location bar
func (t *TagBuilder) HxReplaceUrl(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-replace-url", strconv.FormatBool(value)))
	return t
}

// configures various aspects of the request
func (t *TagBuilder) HxRequest(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-request", value))
	return t
}

// control how requests made by different elements are synchronized
func (t *TagBuilder) HxSync(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-sync", value))
	return t
}

// force elements to validate themselves before a request
func (t *TagBuilder) HxValidate(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-validate", value))
	return t
}

// push a URL into the browser location bar to create history
func (t *TagBuilder) HxPushUrl(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-push-url", value))
	return t
}

// prevent sensitive data being saved to the history cache
func (t *TagBuilder) HxHistory(value bool) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-history", strconv.FormatBool(value)))
	return t
}

func (t *TagBuilder) HxHistoryElt() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-history-elt", ""))
	return t
}

// issues a PUT to the specified URL
func (t *TagBuilder) HxPut(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-put", value))
	return t
}
