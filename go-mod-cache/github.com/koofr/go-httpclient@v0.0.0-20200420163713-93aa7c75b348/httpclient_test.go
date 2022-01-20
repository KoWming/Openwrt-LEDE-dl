package httpclient_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/koofr/go-httpclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ExampleStruct struct {
	Key string `xml:"Key"`
}

type InvalidStruct struct {
	Key complex128 `xml:"Key"`
}

type errorReader struct {
	err error
}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, r.err
}

var ts *httptest.Server
var client *HTTPClient
var handler func(http.ResponseWriter, *http.Request)

var _ = Describe("HTTPClient", func() {
	BeforeEach(func() {
		handler = nil

		ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer GinkgoRecover()

			if handler == nil {
				fmt.Fprintln(w, "ok")
			} else {
				handler(w, r)
			}
		}))

		u, _ := url.Parse(ts.URL)

		client = New()
		client.Client = ts.Client()
		client.BaseURL = u
	})

	AfterEach(func() {
		ts.Close()
	})

	Describe("New", func() {
		It("should create new default client", func() {
			res, err := client.Request(&RequestData{
				Method:  "GET",
				FullURL: ts.URL,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.StatusCode).To(Equal(200))
		})

		It("should fail to make request with invalid certificate", func() {
			ts.Close()
			ts = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}))
			client.Client = ts.Client()

			u, _ := url.Parse(ts.URL)
			client = New()
			client.BaseURL = u

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).To(HaveOccurred())
		})

		It("should get error if remote server is not reachable", func() {
			ts.Close()

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).To(HaveOccurred())
		})

		It("should get error if FullURL is invalid", func() {
			_, err := client.Request(&RequestData{
				Method:  "GET",
				FullURL: "://???",
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(`missing protocol scheme`))
		})
	})

	Describe("Insecure", func() {
		It("should create new insecure http client (ignoring TLS errors)", func() {
			ts.Close()
			ts = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}))
			client.Client = ts.Client()

			u, _ := url.Parse(ts.URL)
			client = Insecure()
			client.BaseURL = u

			res, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(res.StatusCode).To(Equal(200))
		})
	})

	Describe("Request", func() {
		It("should pass the context", func() {
			requestCtx, requestCtxCancel := context.WithCancel(context.Background())

			handler = func(w http.ResponseWriter, r *http.Request) {
				defer GinkgoRecover()
				requestCtxCancel()
				Eventually(r.Context().Err).Should(HaveOccurred())
			}

			_, err := client.Request(&RequestData{
				Context: requestCtx,
				Method:  "GET",
				Path:    "/",
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("context canceled"))
		})
	})

	Describe("Headers", func() {
		It("should set global headers", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Header.Get("X-Header")).To(Equal("value"))
				fmt.Fprintln(w, "ok")
			}

			client.Headers.Set("X-Header", "value")

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("SetPostHook", func() {
		It("should add post response hook for status code", func() {
			postHookError := fmt.Errorf("Post hook error")

			client.SetPostHook(http.StatusOK, func(req *http.Request, res *http.Response) error {
				return postHookError
			})

			_, err := client.Request(&RequestData{
				Method:  "GET",
				FullURL: ts.URL,
			})
			Expect(err).To(Equal(postHookError))
		})
	})

	Describe("SetErrorHandler", func() {
		It("should set the error handler", func() {
			ts.Close()

			errorHandled := false

			client.SetErrorHandler(func(res *http.Response, err error) error {
				Expect(res).To(BeNil())
				Expect(err).To(HaveOccurred())
				errorHandled = true
				return err
			})

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).To(HaveOccurred())

			Expect(errorHandled).To(BeTrue())
		})
	})

	Describe("SetRateLimit", func() {
		It("should rate limit requests without timeout", func() {
			responseLock := &sync.RWMutex{}
			responseLock.Lock()
			var counter int32 = 0

			handler = func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&counter, 1)
				responseLock.RLock()
				fmt.Fprintln(w, "ok")
			}

			client.SetRateLimit(10, 0)

			for i := 0; i < 20; i++ {
				go func() {
					client.Request(&RequestData{
						Method:  "GET",
						FullURL: ts.URL,
					})
				}()
			}

			time.Sleep(1 * time.Second)

			c := counter

			responseLock.Unlock()

			Expect(int(c)).To(Equal(10))

			time.Sleep(1 * time.Second)

			Expect(int(counter)).To(Equal(20))
		})

		It("should rate limit requests with timeout", func() {
			responseLock := &sync.RWMutex{}
			responseLock.Lock()

			handler = func(w http.ResponseWriter, r *http.Request) {
				responseLock.RLock()
				fmt.Fprintln(w, "ok")
			}

			client.SetRateLimit(10, 1*time.Second)

			var errors int32 = 0

			for i := 0; i < 15; i++ {
				go func() {
					_, err := client.Request(&RequestData{
						Method:  "GET",
						FullURL: ts.URL,
					})

					if err == RateLimitTimeoutError {
						atomic.AddInt32(&errors, 1)
					}
				}()
			}

			time.Sleep(2 * time.Second)

			responseLock.Unlock()

			Expect(int(errors)).To(Equal(5))
		})
	})

	Describe("Path", func() {
		It("should set request path", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Path).To(Equal("/foo+bar"))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/foo+bar",
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Params", func() {
		It("should set request query string", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Query().Encode()).To(Equal("key=value"))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
				Params: url.Values{"key": {"value"}},
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Headers", func() {
		It("should set request headers", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Header.Get("X-Header")).To(Equal("value"))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:  "GET",
				Path:    "/",
				Headers: http.Header{"X-Header": {"value"}},
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should override global headers", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				Expect(r.Header.Get("X-Header")).To(Equal("override"))
				fmt.Fprintln(w, "ok")
			}

			client.Headers.Set("X-Header", "value")

			_, err := client.Request(&RequestData{
				Method:  "GET",
				Path:    "/",
				Headers: http.Header{"X-Header": {"override"}},
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("ReqReader", func() {
		It("should set request body reader", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(Equal([]byte("body")))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:    "POST",
				Path:      "/",
				ReqReader: bytes.NewReader([]byte("body")),
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("ReqValue", func() {
		It("should set JSON request body with EncodingJSON", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(Equal([]byte(`{"key":"value"}`)))
				Expect(r.Header.Get("content-type")).To(Equal("application/json"))
				fmt.Fprintln(w, "ok")
			}

			data := map[string]string{"key": "value"}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingJSON,
				ReqValue:    data,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not set JSON content-type for empty body with EncodingJSON", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(BeEmpty())
				Expect(r.Header.Get("content-type")).NotTo(Equal("application/json"))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingJSON,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not set JSON request body with EncodingJSON and invalid data", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			data := InvalidStruct{
				Key: complex(42, 42),
			}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingJSON,
				ReqValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("json: unsupported type: complex128"))
		})

		It("should set XML request body with EncodingXML", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(Equal([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<ExampleStruct><Key>value</Key></ExampleStruct>`)))
				Expect(r.Header.Get("content-type")).To(Equal("application/xml"))
				fmt.Fprintln(w, "ok")
			}

			data := ExampleStruct{
				Key: "value",
			}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingXML,
				ReqValue:    data,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not set XML content-type for empty body with EncodingXML", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(BeEmpty())
				Expect(r.Header.Get("content-type")).NotTo(Equal("application/xml"))
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingXML,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not set XML request body with EncodingXML and invalid data", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			data := InvalidStruct{
				Key: complex(42, 42),
			}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingXML,
				ReqValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("xml: unsupported type: complex128"))
		})

		It("should set urlencoded request body with EncodingForm", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				body, _ := ioutil.ReadAll(r.Body)
				Expect(body).To(Equal([]byte(`key=value`)))
				Expect(r.Header.Get("content-type")).To(Equal("application/x-www-form-urlencoded"))
				fmt.Fprintln(w, "ok")
			}

			data := make(url.Values)
			data.Set("key", "value")

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingForm,
				ReqValue:    data,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not set urlencoded request body with EncodingForm and invalid data", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			data := "data"

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: EncodingForm,
				ReqValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("HTTPClient: invalid ReqValue type string"))
		})

		It("should not set request body with invalid ReqEncoding", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			data := map[string]string{"key": "value"}

			_, err := client.Request(&RequestData{
				Method:      "POST",
				Path:        "/",
				ReqEncoding: Encoding("invalid"),
				ReqValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("HTTPClient: invalid ReqEncoding: invalid"))
		})
	})

	Describe("ExpectedStatus", func() {
		It("should filter response status", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(201)
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:         "GET",
				Path:           "/",
				ExpectedStatus: []int{201},
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return error for wrong response status", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(400)
				fmt.Fprintln(w, "fail")
			}

			res, err := client.Request(&RequestData{
				Method:         "GET",
				Path:           "/",
				ExpectedStatus: []int{201},
			})
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(InvalidStatusError{
				Expected: []int{201},
				Got:      400,
				Headers:  res.Header,
				Content:  "fail\n",
			}))
			Expect(strings.HasPrefix(err.Error(), "Invalid response status! Got 400, expected [201]; headers: ")).To(BeTrue())
			Expect(strings.HasSuffix(err.Error(), ", content: fail\n")).To(BeTrue())
		})
	})

	Describe("IgnoreRedirects", func() {
		It("should follow redirects redirects by default", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/redirect" {
					w.WriteHeader(201)
				} else {
					w.Header().Set("Location", "/redirect")
					w.WriteHeader(301)
				}
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:         "GET",
				Path:           "/",
				ExpectedStatus: []int{201},
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should ignore redirects", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/redirect" {
					w.WriteHeader(201)
				} else {
					w.Header().Set("Location", "/redirect")
					w.WriteHeader(301)
				}
				fmt.Fprintln(w, "ok")
			}

			_, err := client.Request(&RequestData{
				Method:          "GET",
				Path:            "/",
				ExpectedStatus:  []int{301},
				IgnoreRedirects: true,
			})
			Expect(err).NotTo(HaveOccurred())
		})

		It("should ignore redirects with missing http client transport", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/redirect" {
					w.WriteHeader(201)
				} else {
					w.Header().Set("Location", "/redirect")
					w.WriteHeader(301)
				}
				fmt.Fprintln(w, "ok")
			}

			client.Client = &http.Client{}

			_, err := client.Request(&RequestData{
				Method:          "GET",
				Path:            "/",
				ExpectedStatus:  []int{301},
				IgnoreRedirects: true,
			})
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("RespValue", func() {
		It("should unmarshal JSON response with EncodingJSON", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/json")
				fmt.Fprint(w, `{"key":"value"}`)
			}

			data := map[string]string{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingJSON,
				RespValue:    &data,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(Equal(map[string]string{"key": "value"}))
		})

		It("should not unmarshal invalid JSON response with EncodingJSON", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/json")
				fmt.Fprint(w, `{"key":"value"`)
			}

			data := map[string]string{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingJSON,
				RespValue:    &data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unexpected end of JSON input"))
		})

		It("should not unmarshal JSON response with EncodingJSON and non-pointer RespValue", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/json")
				fmt.Fprint(w, `{"key":"value"}`)
			}

			data := map[string]string{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingJSON,
				RespValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("json: Unmarshal(non-pointer map[string]string)"))
		})

		It("should not unmarshal JSON response with EncodingJSON and server error", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/json")
				w.Header().Set("content-length", "42")
				fmt.Fprint(w, `{"key":"value"}`)
			}

			data := map[string]string{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingJSON,
				RespValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unexpected EOF"))
		})

		It("should unmarshal XML response with EncodingXML", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/xml")
				fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
<ExampleStruct><Key>value</Key></ExampleStruct>`)
			}

			data := ExampleStruct{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingXML,
				RespValue:    &data,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data.Key).To(Equal("value"))
		})

		It("should not unmarshal invalid XML response with EncodingXML", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/xml")
				fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
<ExampleStruct><Key>value</Key></ExampleStruct`)
			}

			data := ExampleStruct{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingXML,
				RespValue:    &data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("XML syntax error on line 2: unexpected EOF"))
		})

		It("should not unmarshal XML response with EncodingXML and non-pointer RespValue", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/xml")
				fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
<ExampleStruct><Key>value</Key></ExampleStruct>`)
			}

			data := ExampleStruct{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingXML,
				RespValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("non-pointer passed to Unmarshal"))
		})

		It("should not unmarshal XML response with EncodingXML and server error", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-type", "application/xml")
				w.Header().Set("content-length", "42")
				fmt.Fprint(w, `<?xml version="1.0" encoding="UTF-8"?>
<ExampleStruct><Key>value</Key></ExampleStruct>`)
			}

			data := ExampleStruct{}

			_, err := client.Request(&RequestData{
				Method:       "GET",
				Path:         "/",
				RespEncoding: EncodingXML,
				RespValue:    data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unexpected EOF"))
		})

		It("should read byte slice", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			data := []byte{}

			_, err := client.Request(&RequestData{
				Method:    "GET",
				Path:      "/",
				RespValue: &data,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(data).To(Equal([]byte("ok\n")))
		})

		It("should not read byte slice with server error", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("content-length", "42")
				fmt.Fprintln(w, "ok")
			}

			data := []byte{}

			_, err := client.Request(&RequestData{
				Method:    "GET",
				Path:      "/",
				RespValue: &data,
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unexpected EOF"))
		})
	})

	Describe("RespConsume", func() {
		It("should not consume response body by default", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			res, err := client.Request(&RequestData{
				Method: "GET",
				Path:   "/",
			})
			Expect(err).NotTo(HaveOccurred())

			n, err := res.Body.Read([]byte{0, 0, 0, 0})

			Expect(n).To(Equal(3))
			Expect(err).To(Equal(io.EOF))
		})

		It("should consume response", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}

			res, err := client.Request(&RequestData{
				Method:      "GET",
				Path:        "/",
				RespConsume: true,
			})
			Expect(err).NotTo(HaveOccurred())

			n, err := res.Body.Read([]byte{0, 0, 0, 0})

			Expect(n).To(Equal(0))
			Expect(err.Error()).To(Equal("http: read on closed response body"))
		})
	})

	Describe("UploadFile", func() {
		It("should upload a file", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				reader, err := r.MultipartReader()
				Expect(err).NotTo(HaveOccurred())
				p, err := reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				body, err := ioutil.ReadAll(p)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte("body")))
				Expect(p.FormName()).To(Equal("file"))
				Expect(p.FileName()).To(Equal("filename.txt"))

				fmt.Fprintln(w, "ok")
			}

			req := &RequestData{
				Method: "POST",
				Path:   "/",
			}

			err := req.UploadFile("file", "filename.txt", bytes.NewReader([]byte("body")))
			Expect(err).NotTo(HaveOccurred())

			_, err = client.Request(req)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should not upload a file with broken reader", func() {
			waitCh := make(chan bool)

			handler = func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					waitCh <- true
				}()
				defer GinkgoRecover()
				reader, err := r.MultipartReader()
				Expect(err).NotTo(HaveOccurred())
				p, err := reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				_, err = ioutil.ReadAll(p)
				Expect(err).To(HaveOccurred())

				fmt.Fprintln(w, "ok")
			}

			req := &RequestData{
				Method: "POST",
				Path:   "/",
			}

			buf := make([]byte, 1*1024*1024, 1*1024*1024) // must be greater than 759K

			bodyReader := io.MultiReader(bytes.NewReader(buf), &errorReader{fmt.Errorf("broken body")})

			err := req.UploadFile("file", "filename.txt", bodyReader)
			Expect(err).NotTo(HaveOccurred())

			_, err = client.Request(req)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("broken body"))

			<-waitCh
		})

		It("should not upload a file with broken form", func() {
			req := &RequestData{
				Method: "POST",
				Path:   "/",
			}

			err := req.UploadFile("file", "filename.txt", bytes.NewReader([]byte("body")))
			Expect(err).NotTo(HaveOccurred())

			req.ReqReader.(*io.PipeReader).CloseWithError(fmt.Errorf("broken form"))
		})
	})

	Describe("UploadFileExtra", func() {
		It("should upload a file with extra fields", func() {
			handler = func(w http.ResponseWriter, r *http.Request) {
				reader, err := r.MultipartReader()
				Expect(err).NotTo(HaveOccurred())
				p, err := reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				body, err := ioutil.ReadAll(p)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte("bar")))
				Expect(p.FormName()).To(Equal("foo"))
				Expect(p.FileName()).To(Equal(""))
				p, err = reader.NextPart()
				Expect(err).NotTo(HaveOccurred())
				body, err = ioutil.ReadAll(p)
				Expect(err).NotTo(HaveOccurred())
				Expect(body).To(Equal([]byte("body")))
				Expect(p.FormName()).To(Equal("file"))
				Expect(p.FileName()).To(Equal("filename.txt"))

				fmt.Fprintln(w, "ok")
			}

			req := &RequestData{
				Method: "POST",
				Path:   "/",
			}

			extra := map[string]string{
				"foo": "bar",
			}

			err := req.UploadFileExtra("file", "filename.txt", bytes.NewReader([]byte("body")), extra)
			Expect(err).NotTo(HaveOccurred())

			_, err = client.Request(req)
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
