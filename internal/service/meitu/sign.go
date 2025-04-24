package meitu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	BasicDateFormat     = "20060102T150405Z"
	Algorithm           = "SDK-HMAC-SHA256"
	HeaderXDate         = "X-Sdk-Date"
	HeaderHost          = "Host"
	HeaderAuthorization = "Authorization"
	HeaderContentSha256 = "X-Sdk-Content-Sha256"
)

type Signer struct {
	Key    string
	Secret string
}

func NewSigner(key, secret string) *Signer {
	return &Signer{
		Key:    key,
		Secret: secret,
	}
}

func (s *Signer) signStringToSign(stringToSign string, signingKey []byte) (string, error) {
	h := hmac.New(sha256.New, signingKey)
	if _, err := h.Write([]byte(stringToSign)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s *Signer) authHeaderValue(signature, accessKey string, signedHeaders []string) string {
	signedHeadersStr := strings.Join(signedHeaders, ";")
	headerValue := fmt.Sprintf("%s Access=%s, SignedHeaders=%s, Signature=%s", Algorithm, accessKey, signedHeadersStr, signature)
	encodedVal := base64.StdEncoding.EncodeToString([]byte(headerValue))
	return "Bearer " + encodedVal
}

func (s *Signer) canonicalRequest(method, url string, headers http.Header, body string, signedHeaders []string) string {
	parsedURL := parseURL(url)
	canonicalURI := s.canonicalURI(parsedURL.Path)
	canonicalQueryString := parsedURL.Query().Encode()
	canonicalHeaders := s.canonicalHeaders(headers, signedHeaders)
	signedHeadersStr := strings.Join(signedHeaders, ";")

	var hexencode string
	if hex := headers.Get(HeaderContentSha256); hex != "" {
		hexencode = hex
	} else {
		hexencode = hashSHA256(body)
	}

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		method, canonicalURI, canonicalQueryString, canonicalHeaders,
		signedHeadersStr, hexencode)
}

func (s *Signer) canonicalURI(path string) string {
	if len(path) == 0 || !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}
func (s *Signer) canonicalHeaders(headers http.Header, signedHeaders []string) string {
	lowHeaders := make(map[string]string)
	for key, value := range headers {
		lowHeaders[strings.ToLower(key)] = strings.TrimSpace(value[0])
	}

	var a []string
	for _, key := range signedHeaders {
		value := lowHeaders[strings.ToLower(key)]
		if strings.EqualFold(strings.ToLower(key), HeaderHost) {
			value = headers.Get(HeaderHost)
		}
		a = append(a, key+":"+value)
	}

	sort.Strings(a)
	return strings.Join(a, "\n")
}

func (s *Signer) signedHeaders(headers http.Header) []string {
	var signedHeaders []string
	for key := range headers {
		signedHeaders = append(signedHeaders, strings.ToLower(key))
	}
	sort.Strings(signedHeaders)
	return signedHeaders
}

func (s *Signer) Sign(url, method string, headers http.Header, body string) (*http.Request, error) {
	dt := headers.Get(HeaderXDate)
	var t time.Time
	var err error
	if dt == "" {
		t = time.Now().UTC()
		headers.Set(HeaderXDate, t.Format(BasicDateFormat))
	} else {
		t, err = time.Parse(BasicDateFormat, dt)
		if err != nil {
			return nil, err
		}
	}

	signedHeaders := s.signedHeaders(headers)
	canonicalRequest := s.canonicalRequest(method, url, headers, body, signedHeaders)
	stringToSign := s.stringToSign(canonicalRequest, t.Format(BasicDateFormat))
	signature, err := s.signStringToSign(stringToSign, []byte(s.Secret))
	if err != nil {
		return nil, err
	}
	authValue := s.authHeaderValue(signature, s.Key, signedHeaders)
	headers.Set(HeaderAuthorization, authValue)

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return req, nil
}

func (s *Signer) stringToSign(canonicalRequest, timeFormat string) string {
	hash := hashSHA256(canonicalRequest)
	return fmt.Sprintf("%s\n%s\n%s",
		Algorithm, timeFormat, hash)
}

func hashSHA256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func hashHMAC(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func parseURL(urlStr string) *url.URL {
	parsedURL, _ := url.Parse(urlStr)
	return parsedURL
}
