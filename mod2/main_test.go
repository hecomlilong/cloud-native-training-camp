package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWebServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Web Server Suite")
}

func ExampleHandler_StdoutHandler() {
	var (
		client *http.Client
		host   = "localhost:1880"
	)
	client = &http.Client{}

	req, err := http.NewRequest("GET", "http://"+host+"/healthz", nil)
	Expect(err).To(BeNil())
	resp, err := client.Do(req)
	Expect(err).To(BeNil())
	fmt.Println(resp.StatusCode)
	// Output: 200
}

var _ = Describe("web server intergration", func() {
	var (
		client *http.Client
		host   = "localhost:1880"
	)

	BeforeEach(func() {
		client = &http.Client{}
	})
	It("header", func() {
		headerCases := map[string]string{
			"X-Request-Id":  "12345678xyz",
			"Authorization": "Token 222",
		}
		req, err := http.NewRequest("GET", "http://"+host+"/healthz", nil)
		Expect(err).To(BeNil())
		for k, v := range headerCases {
			req.Header.Add(k, v)
		}
		resp, err := client.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		Expect(err).To(BeNil())
		for k, v := range headerCases {
			Expect(len(resp.Header[k])).ToNot(Equal(0))
			Expect(resp.Header[k][0]).To(Equal(v))
		}
	})
	It("version", func() {
		version := "Version"
		req, err := http.NewRequest("GET", "http://"+host+"/healthz", nil)
		Expect(err).To(BeNil())
		resp, err := client.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(len(resp.Header[version])).ToNot(Equal(0))
	})
	It("healthz", func() {
		req, err := http.NewRequest("GET", "http://"+host+"/healthz", nil)
		Expect(err).To(BeNil())
		resp, err := client.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).To(Equal(200))
		req, err = http.NewRequest("GET", "http://"+host+"/nobody", nil)
		Expect(err).To(BeNil())
		resp, err = client.Do(req)
		Expect(err).To(BeNil())
		defer resp.Body.Close()
		Expect(resp.StatusCode).ToNot(Equal(200))
	})
})
