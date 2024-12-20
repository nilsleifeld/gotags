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
func (t *TagBuilder[T]) HxSwap(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-swap", string(value)))
	return t.AsT()
}

// specifies the event that triggers the request
func (t *TagBuilder[T]) HxTrigger(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-trigger", value))
	return t.AsT()
}

// issues a GET to the specified URL
func (t *TagBuilder[T]) HxGet(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-get", value))
	return t.AsT()
}

// issues a POST to the specified URL
func (t *TagBuilder[T]) HxPost(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-post", value))
	return t.AsT()
}

// issues a PATCH to the specified URL
func (t *TagBuilder[T]) HxPatch(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-patch", value))
	return t.AsT()
}

// issues a DELETE to the specified URL
func (t *TagBuilder[T]) HxDelete(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-delete", value))
	return t.AsT()
}

// the element to put the htmx-request class on during the request
func (t *TagBuilder[T]) HxIndicator(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-indicator", value))
	return t.AsT()
}

// specifies the target element to be swapped
func (t *TagBuilder[T]) HxTarget(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-target", value))
	return t.AsT()
}

// mark element to swap in from a response (out of band)
func (t *TagBuilder[T]) HxSwapOob(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-swap-oob", strconv.FormatBool(value)))
	return t.AsT()
}

// shows a confirm() dialog before issuing a request
func (t *TagBuilder[T]) HxConfirm(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-confirm", value))
	return t.AsT()
}

// add progressive enhancement for links and forms
func (t *TagBuilder[T]) HxBoost(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-boost", strconv.FormatBool(value)))
	return t.AsT()
}

// extensions to use for this element
func (t *TagBuilder[T]) HxExt(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-ext", value))
	return t.AsT()
}

// handle events with inline scripts on elements
func (t *TagBuilder[T]) HxOn(event string, value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-on:"+event, value))
	return t.AsT()
}

// select content to swap in from a response
func (t *TagBuilder[T]) HxSelect(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-select", value))
	return t.AsT()
}

// add values to submit with the request (JSON format)
func (t *TagBuilder[T]) HxVals(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-vals", value))
	return t.AsT()
}

// add values to submit with the request (JSON format)
func (t *TagBuilder[T]) HxValsJson(obj interface{}) *T {
	value, _ := json.Marshal(obj)
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-vals", string(value)))
	return t.AsT()
}

// select content to swap in from a response, somewhere other than the target (out of band)
func (t *TagBuilder[T]) HxSelectOob(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-select-oob", strconv.FormatBool(value)))
	return t.AsT()
}

// issues a DELETE to the specified URL
func (t *TagBuilder[T]) HxDisable(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disable", strconv.FormatBool(value)))
	return t.AsT()
}

// adds the disabled attribute to the specified elements while a request is in flight
func (t *TagBuilder[T]) HxDisabledElt(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disabled-elt", value))
	return t.AsT()
}

// control and disable automatic attribute inheritance for child nodes
func (t *TagBuilder[T]) HxDisinherit(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-disinherit", strconv.FormatBool(value)))
	return t.AsT()
}

// changes the request encoding type
func (t *TagBuilder[T]) HxEncoding(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-encoding", value))
	return t.AsT()
}

// adds to the headers that will be submitted with the request
func (t *TagBuilder[T]) HxHeaders(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-headers", value))
	return t.AsT()
}

// include additional data in requests
func (t *TagBuilder[T]) HxInlude(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-include", value))
	return t.AsT()
}

// control and enable automatic attribute inheritance for child nodes if it has been disabled by default
func (t *TagBuilder[T]) HxInherit(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-inherit", strconv.FormatBool(value)))
	return t.AsT()
}

// filters the parameters that will be submitted with a request
func (t *TagBuilder[T]) HxParams(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-params", value))
	return t.AsT()
}

// shows a prompt() before submitting a request
func (t *TagBuilder[T]) HxPrompt(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-prompt", value))
	return t.AsT()
}

// specifies elements to keep unchanged between requests
func (t *TagBuilder[T]) HxPreserve(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-preserve", value))
	return t.AsT()
}

// replace the URL in the browser location bar
func (t *TagBuilder[T]) HxReplaceUrl(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-replace-url", strconv.FormatBool(value)))
	return t.AsT()
}

// configures various aspects of the request
func (t *TagBuilder[T]) HxRequest(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-request", value))
	return t.AsT()
}

// control how requests made by different elements are synchronized
func (t *TagBuilder[T]) HxSync(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-sync", value))
	return t.AsT()
}

// force elements to validate themselves before a request
func (t *TagBuilder[T]) HxValidate(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-validate", value))
	return t.AsT()
}

// push a URL into the browser location bar to create history
func (t *TagBuilder[T]) HxPushUrl(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-push-url", value))
	return t.AsT()
}

// prevent sensitive data being saved to the history cache
func (t *TagBuilder[T]) HxHistory(value bool) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-history", strconv.FormatBool(value)))
	return t.AsT()
}

func (t *TagBuilder[T]) HxHistoryElt() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-history-elt", ""))
	return t.AsT()
}

// issues a PUT to the specified URL
func (t *TagBuilder[T]) HxPut(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("data-hx-put", value))
	return t.AsT()
}
