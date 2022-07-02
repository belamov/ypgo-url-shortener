package handlers

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/belamov/ypgo-url-shortener/internal/app/config"
	"github.com/belamov/ypgo-url-shortener/internal/app/mocks"
	"github.com/belamov/ypgo-url-shortener/internal/app/models"
	"github.com/belamov/ypgo-url-shortener/internal/app/responses"
	"github.com/belamov/ypgo-url-shortener/internal/app/services"
	"github.com/belamov/ypgo-url-shortener/internal/app/services/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body string, cookies map[string]string) (*http.Response, string) {
	t.Helper()

	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	req, err = http.NewRequest(method, ts.URL+path, strings.NewReader(body))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if len(cookies) > 0 {
		for name, value := range cookies {
			req.AddCookie(&http.Cookie{
				Name:     name,
				Value:    value,
				Secure:   true,
				HttpOnly: true,
			})
		}
	}
	resp, err = client.Do(req)
	require.NoError(t, err)

	respBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	require.NoError(t, err)

	return resp, string(bytes.TrimSpace(respBody))
}

func testGzippedRequest(t *testing.T, ts *httptest.Server, method, path string, body string) (*http.Response, string) {
	t.Helper()

	var err error
	var req *http.Request
	var resp *http.Response
	var respBody []byte

	var b bytes.Buffer

	w, _ := gzip.NewWriterLevel(&b, gzip.BestCompression)

	_, err = w.Write([]byte(body))
	if err != nil {
		return nil, err.Error()
	}

	err = w.Close()
	if err != nil {
		return nil, err.Error()
	}

	req, err = http.NewRequest(method, ts.URL+path, bytes.NewReader(b.Bytes()))
	require.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err = client.Do(req)
	require.NoError(t, err)

	respBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	require.NoError(t, err)

	return resp, string(bytes.TrimSpace(respBody))
}

func TestHandler_Expand(t *testing.T) {
	type want struct {
		statusCode int
		location   string
		body       string
	}
	tests := []struct {
		name    string
		request string
		want    want
		method  string
	}{
		{
			name: "get with existing id",
			want: want{
				statusCode: http.StatusTemporaryRedirect,
				location:   "/url",
				body:       "",
			},
			request: "/id",
			method:  http.MethodGet,
		},
		{
			name: "get with missing id",
			want: want{
				statusCode: http.StatusNotFound,
				location:   "",
				body:       "cant find full url",
			},
			request: "/missing",
			method:  http.MethodGet,
		},
		{
			name: "get with null id",
			want: want{
				statusCode: http.StatusMethodNotAllowed,
				location:   "",
				body:       "",
			},
			request: "/",
			method:  http.MethodGet,
		},
		{
			name: "not supported method",
			want: want{
				statusCode: http.StatusMethodNotAllowed,
				location:   "",
				body:       "",
			},
			request: "/asd",
			method:  http.MethodDelete,
		},
		{
			name: "it returns 500 when service fails on expanding",
			want: want{
				statusCode: http.StatusInternalServerError,
				location:   "",
				body:       "error text",
			},
			request: "/error",
			method:  http.MethodGet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockRepo)
			mockRepo.On("GetByID", "id").Return("url", nil)
			mockRepo.On("GetByID", "missing").Return("", nil)
			mockRepo.On("GetByID", "error").Return("", errors.New("error text"))

			mockGen := new(mocks.MockGen)

			mockGenID := new(mocks.MockUserIDGenerator)
			mockGenID.On("GenerateUserID").Return("user id")

			cfg := config.New()
			service := services.New(mockRepo, mockGen, cfg)
			r := NewRouter(service, cfg, mockGenID)
			ts := httptest.NewServer(r)
			defer ts.Close()

			result, body := testRequest(t, ts, tt.method, tt.request, "", nil)
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.location, result.Header.Get("Location"))
			assert.Equal(t, tt.want.body, body)
		})
	}
}

func TestHandler_UserURLs(t *testing.T) {
	type want struct {
		body       []responses.UsersShortURL
		statusCode int
	}
	tests := []struct {
		name    string
		request string
		want    want
		userID  string
		method  string
	}{
		{
			name: "get user's urls",
			want: want{
				body: []responses.UsersShortURL{
					{ShortURL: "http://localhost:8080/id", OriginalURL: "url"},
				},
				statusCode: http.StatusOK,
			},
			userID: "user id with urls",
		},
		{
			name: "get another user's urls ",
			want: want{
				body:       []responses.UsersShortURL{},
				statusCode: http.StatusNoContent,
			},
			userID: "user id without urls",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockRepo)
			mockRepo.On("GetUsersUrls", "user id with urls").Return("url", "id")
			mockRepo.On("GetUsersUrls", "user id without urls").Return("", "")

			mockGen := new(mocks.MockGen)
			mockGenID := new(mocks.MockUserIDGenerator)

			cfg := config.New()
			service := services.New(mockRepo, mockGen, cfg)
			r := NewRouter(service, cfg, mockGenID)
			ts := httptest.NewServer(r)
			defer ts.Close()

			cryptographer := crypto.GCMAESCryptographer{Key: cfg.EncryptionKey}
			encryptedCookieValue, _ := cryptographer.Encrypt([]byte(tt.userID))
			cookies := map[string]string{
				UserIDCookieName: hex.EncodeToString(encryptedCookieValue),
			}
			result, body := testRequest(t, ts, http.MethodGet, "/api/user/urls", "", cookies)
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			if tt.want.statusCode == http.StatusNoContent {
				assert.Equal(t, "", body)
				return
			}
			expectedJSON, err := json.Marshal(tt.want.body)
			assert.NoError(t, err)
			assert.JSONEq(t, string(expectedJSON), body)
		})
	}
}

func TestHandler_Shorten(t *testing.T) {
	type want struct {
		statusCode int
		body       string
	}
	tests := []struct {
		name   string
		want   want
		body   string
		method string
	}{
		{
			name: "post with url",
			want: want{
				statusCode: http.StatusCreated,
				body:       "http://localhost:8080/id",
			},
			method: http.MethodPost,
			body:   "url",
		},
		{
			name: "post without url",
			want: want{
				statusCode: http.StatusBadRequest,
				body:       "url required",
			},
			method: http.MethodPost,
			body:   "",
		},
		{
			name: "not supported method",
			want: want{
				statusCode: http.StatusMethodNotAllowed,
				body:       "",
			},
			method: http.MethodGet,
			body:   "",
		},
		{
			name: "it returns 500 when service fails on shortening",
			want: want{
				statusCode: http.StatusInternalServerError,
				body:       "err",
			},
			method: http.MethodPost,
			body:   "error_on_shortening",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockRepo)
			mockRepo.On("Save", models.ShortURL{
				OriginalURL: "url",
				ID:          "id",
				CreatedByID: "user id",
			}).Return(nil)
			mockGen := new(mocks.MockGen)
			mockGen.On("GenerateIDFromString", "url").Return("id", nil)
			mockGen.On("GenerateIDFromString", "").Return("", errors.New("err"))
			mockGen.On("GenerateIDFromString", "error_on_shortening").Return("", errors.New("err"))

			mockGenID := new(mocks.MockUserIDGenerator)
			mockGenID.On("GenerateUserID").Return("user id")

			cfg := config.New()
			service := services.New(mockRepo, mockGen, cfg)
			r := NewRouter(service, cfg, mockGenID)
			ts := httptest.NewServer(r)
			defer ts.Close()

			result, body := testRequest(t, ts, tt.method, "/", tt.body, nil)
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.body, body)

			result, body = testGzippedRequest(t, ts, tt.method, "/", tt.body)
			defer result.Body.Close()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.body, body)
			if tt.want.statusCode == http.StatusCreated {
				assert.NotEmpty(t, result.Header.Get("Set-Cookie"))
			}
		})
	}
}

func TestHandler_ShortenAPI(t *testing.T) {
	type want struct {
		statusCode  int
		body        string
		contentType string
	}
	tests := []struct {
		name   string
		want   want
		body   string
		method string
	}{
		{
			name: "post with url",
			want: want{
				statusCode:  http.StatusCreated,
				body:        "{\"result\":\"http://localhost:8080/id\"}",
				contentType: "application/json",
			},
			method: http.MethodPost,
			body:   "{\"url\":\"url\"}",
		},
		{
			name: "post without url",
			want: want{
				statusCode: http.StatusBadRequest,
				body:       "url required",
			},
			method: http.MethodPost,
			body:   "{\"url\":\"\"}",
		},
		{
			name: "post without url param",
			want: want{
				statusCode: http.StatusBadRequest,
				body:       "url required",
			},
			method: http.MethodPost,
			body:   "{\"some param\":\"some value\"}",
		},
		{
			name: "post without json",
			want: want{
				statusCode: http.StatusBadRequest,
				body:       "cannot decode json",
			},
			method: http.MethodPost,
			body:   "url",
		},
		{
			name: "not supported method",
			want: want{
				statusCode: http.StatusMethodNotAllowed,
				body:       "",
			},
			method: http.MethodGet,
			body:   "",
		},
		{
			name: "it returns 500 when service fails on shortening",
			want: want{
				statusCode: http.StatusInternalServerError,
				body:       "err",
			},
			method: http.MethodPost,
			body:   "{\"url\":\"error_on_shortening\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockRepo)
			mockRepo.On("Save", models.ShortURL{
				OriginalURL: "url",
				ID:          "id",
				CreatedByID: "user id",
			}).Return(nil)
			mockGen := new(mocks.MockGen)
			mockGen.On("GenerateIDFromString", "url").Return("id", nil)
			mockGen.On("GenerateIDFromString", "").Return("", errors.New("err"))
			mockGen.On("GenerateIDFromString", "error_on_shortening").Return("", errors.New("err"))

			mockGenID := new(mocks.MockUserIDGenerator)
			mockGenID.On("GenerateUserID").Return("user id")

			cfg := config.New()
			service := services.New(mockRepo, mockGen, cfg)
			r := NewRouter(service, cfg, mockGenID)
			ts := httptest.NewServer(r)
			defer ts.Close()

			result, body := testRequest(t, ts, tt.method, "/api/shorten", tt.body, nil)
			defer result.Body.Close()

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			}
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.body, body)

			result, body = testGzippedRequest(t, ts, tt.method, "/api/shorten", tt.body)
			defer result.Body.Close()

			if tt.want.contentType != "" {
				assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			}
			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.body, body)
		})
	}
}

func TestHandler_getUserID(t *testing.T) {
	tests := []struct {
		name           string
		cookieRawValue string
		cookieValue    string
		want           string
	}{
		{
			name:           "it returns passed user id correctly",
			cookieRawValue: "user id",
			cookieValue:    "",
			want:           "user id",
		},
		{
			name:           "it generates new user id if no cookie was passed",
			cookieRawValue: "",
			cookieValue:    "",
			want:           "new user id",
		},
		{
			name:           "it generates new user id if no cookie was passed",
			cookieRawValue: "",
			cookieValue:    "75770aa0e5a26e916fc2c4bd76da15aba5284340f1fdea947d25a56dcf9b15354d3bda",
			want:           "new user id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mocks.MockRepo)
			mockGen := new(mocks.MockGen)
			mockGenID := new(mocks.MockUserIDGenerator)
			mockGenID.On("GenerateUserID").Return("new user id")

			cfg := config.New()
			service := services.New(mockRepo, mockGen, cfg)
			r := NewRouter(service, cfg, mockGenID)

			ts := httptest.NewServer(r)
			defer ts.Close()

			req, err := http.NewRequest(http.MethodGet, "", nil)

			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")

			h := NewHandler(service, cfg, mockGenID)

			if tt.cookieRawValue != "" {
				encryptedCookieValue, err := h.crypto.Encrypt([]byte(tt.cookieRawValue))
				encodedCookieValue := hex.EncodeToString(encryptedCookieValue)
				require.NoError(t, err)
				req.AddCookie(&http.Cookie{
					Name:     UserIDCookieName,
					Value:    encodedCookieValue,
					Secure:   true,
					HttpOnly: true,
				})
			}

			if tt.cookieValue != "" {
				req.AddCookie(&http.Cookie{
					Name:     UserIDCookieName,
					Value:    tt.cookieValue,
					Secure:   true,
					HttpOnly: true,
				})
			}

			res := h.getUserID(req)

			assert.Equal(t, tt.want, res)
		})
	}
}
