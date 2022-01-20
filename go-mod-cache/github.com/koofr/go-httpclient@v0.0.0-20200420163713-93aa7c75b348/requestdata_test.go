package httpclient_test

import (
	. "github.com/koofr/go-httpclient"
	"net/http"
	"net/url"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RequestData", func() {
	Describe("Copy", func() {
		It("should copy the request", func() {
			params := make(url.Values)
			params.Set("param", "value")

			headers := make(http.Header)
			headers.Set("X-Header", "value")

			reqValue := map[string]string{
				"key": "value",
			}

			respValue := map[string]string{}

			req := &RequestData{
				Method:          "GET",
				Path:            "/path",
				Params:          params,
				Headers:         headers,
				ReqEncoding:     EncodingJSON,
				ReqValue:        reqValue,
				ExpectedStatus:  []int{200},
				IgnoreRedirects: true,
				RespEncoding:    EncodingXML,
				RespValue:       &respValue,
				RespConsume:     true,
			}

			ok, reqCopy := req.Copy()
			Expect(ok).To(BeTrue())

			Expect(req).To(Equal(reqCopy))
			Expect(&req.Params == &reqCopy.Params).To(BeFalse())
			Expect(&req.Headers == &reqCopy.Headers).To(BeFalse())
			Expect(&req.ExpectedStatus == &reqCopy.ExpectedStatus).To(BeFalse())
		})
	})

	Describe("CanCopy", func() {
		It("should not copy request with reader", func() {
			req := &RequestData{
				Method:    "GET",
				Path:      "/path",
				ReqReader: strings.NewReader("data"),
			}

			canCopy := req.CanCopy()
			Expect(canCopy).To(BeFalse())

			ok, _ := req.Copy()
			Expect(ok).To(BeFalse())
		})
	})
})
