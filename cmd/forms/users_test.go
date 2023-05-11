package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"


	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	templateDirKey   string = "TEMPLATE_DIR"
	templateDirValue string = "../modern/templates/"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.ReleaseMode)
	m.Run()
}

func TestPageHandler(t *testing.T) {
	t.Setenv(templateDirKey, templateDirValue)

	type testPage struct {
		testName     string
		pagePath     string
		request      string
		expectedCode int
	}

	testCases := []testPage{
		{testName: "index page Check", pagePath: "/index", request: "GET", expectedCode: 200},
		{testName: "about page Check", pagePath: "/about", request: "GET", expectedCode: 200},
		{testName: "blog-home page Check", pagePath: "/blog-home", request: "GET", expectedCode: 200},
		{testName: "blog-post page Check", pagePath: "/blog-post", request: "GET", expectedCode: 200},
		{testName: "contact page Check", pagePath: "/contact", request: "GET", expectedCode: 200},
		{testName: "faq page Check", pagePath: "/faq", request: "GET", expectedCode: 200},
		{testName: "portfolio-item page Check", pagePath: "/portfolio-item", request: "GET", expectedCode: 200},
		{testName: "portfolio-overview page Check", pagePath: "/portfolio-overview", request: "GET", expectedCode: 200},
		{testName: "pricing page Check", pagePath: "/pricing", request: "GET", expectedCode: 200},
		{testName: "/ page Check", pagePath: "/", request: "GET", expectedCode: 200},
		{testName: "Inexisting /abc page Check", pagePath: "/abc", request: "GET", expectedCode: 404},
		{testName: "Empty page name Check", pagePath: "", request: "GET", expectedCode: 301},
		{testName: "Inexisting abc page Check", pagePath: "abc", request: "GET", expectedCode: 404},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			r := handlers.Handler()
			w := httptest.NewRecorder()

			req, err := http.NewRequest(tc.request, tc.pagePath, nil)
			if err != nil {
				log.Fatal(err)
			}

			r.ServeHTTP(w, req)
			assert.Equal(t, tc.expectedCode, w.Code)
			t.Logf("%v page check, expected code %v, got code %v", tc.pagePath, tc.expectedCode, w.Code)
		})
	}
}

func TestIndexBody(t *testing.T) {
	t.Setenv(templateDirKey, templateDirValue)

	r := handlers.Handler()
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		log.Fatal(err)
	}

	r.ServeHTTP(w, req)
	p, err := ioutil.ReadAll(w.Body)
	pageOK := err == nil && strings.Index(string(p), "<title>Modern Business") > 0
	assert.True(t, pageOK)
	t.Logf("Find string in the request's body, expected True, got %v", pageOK)
}
