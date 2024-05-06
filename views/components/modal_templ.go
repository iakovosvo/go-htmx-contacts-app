// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Modal() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-show=\"isModalOpen\" x-trap=\"isModalOpen\" x-on:keydown.escape.window=\"isModalOpen = false;clearModalContent(); history.pushState({}, &#39;&#39;, &#39;/&#39;)\" x-on:keydown.enter.window=\"isModalOpen = false;clearModalContent(); history.pushState({}, &#39;&#39;, &#39;/&#39;)\"><div aria-hidden=\"true\" class=\"bg-slate-900/80 bottom-0 fixed left-0 right-0 top-0 z-20\"></div><div role=\"dialog\" aria-modal=\"true\" class=\"fixed bottom-0 p-6 left-0 overflow-auto right-0 top-0 z-30 lg:py-24\"><div class=\"container max-w-xl mx-auto relative shadow-lg z-90\" x-on:click.outside=\"isModalOpen = false;clearModalContent(); history.pushState({}, &#39;&#39;, &#39;/&#39;)\"><div id=\"modal-content\" class=\"p-6 bg-gray-800 rounded-3xl\"></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}