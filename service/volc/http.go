/**
 * Package repo
 * @file      : http.go
 * @author    : xaoyaoyao
 * @contact   : xaoyaoyao@aliyun.com
 * @version   : 1.0.0
 * @time      : 2025/4/22 15:54
 **/

package volc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/volcengine/skd/internal/config"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func hmacSHA256(key []byte, content string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return mac.Sum(nil)
}

func getSignedKey(secretKey, date, region, service string) []byte {
	kDate := hmacSHA256([]byte(secretKey), date)
	kRegion := hmacSHA256(kDate, region)
	kService := hmacSHA256(kRegion, service)
	kSigning := hmacSHA256(kService, "request")
	return kSigning
}

func hashSHA256(data []byte) []byte {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		log.Printf("input hash err:%s", err.Error())
	}

	return hash.Sum(nil)
}

func DoRequest(action, version, method string, queries url.Values, body []byte) ([]byte, []byte, int, error) {
	// 构建请求
	queries.Set("Action", action)
	queries.Set("Version", version)

	Path := config.Get().VolcEnginePath
	Addr := config.Get().VolcEngineAddr

	requestAddr := fmt.Sprintf("%s%s?%s", Addr, Path, queries.Encode())
	request, err := http.NewRequest(method, requestAddr, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, 0, fmt.Errorf("bad request: %w", err)
	}

	// 构建签名材料
	date := time.Now().UTC().Format("20060102T150405Z")
	authDate := date[:8]
	request.Header.Set("X-Date", date)

	payload := hex.EncodeToString(hashSHA256(body))
	request.Header.Set("X-Content-Sha256", payload)
	request.Header.Set("Content-Type", "application/json")

	queryString := strings.Replace(queries.Encode(), "+", "%20", -1)
	signedHeaders := []string{"host", "x-date", "x-content-sha256", "content-type"}
	var headerList []string
	for _, header := range signedHeaders {
		if header == "host" {
			headerList = append(headerList, header+":"+request.Host)
		} else {
			v := request.Header.Get(header)
			headerList = append(headerList, header+":"+strings.TrimSpace(v))
		}
	}
	headerString := strings.Join(headerList, "\n")

	Region := config.Get().VolcEngineRegion
	Service := config.Get().VolcEngineService
	SecretAccessKey := config.Get().SecretAccessKey
	AccessKeyID := config.Get().AccessKeyId

	canonicalString := strings.Join([]string{
		method,
		Path,
		queryString,
		headerString + "\n",
		strings.Join(signedHeaders, ";"),
		payload,
	}, "\n")

	hashedCanonicalString := hex.EncodeToString(hashSHA256([]byte(canonicalString)))

	credentialScope := authDate + "/" + Region + "/" + Service + "/request"
	signString := strings.Join([]string{
		"HMAC-SHA256",
		date,
		credentialScope,
		hashedCanonicalString,
	}, "\n")

	// 构建认证请求头
	signedKey := getSignedKey(SecretAccessKey, authDate, Region, Service)
	signature := hex.EncodeToString(hmacSHA256(signedKey, signString))

	authorization := "HMAC-SHA256" +
		" Credential=" + AccessKeyID + "/" + credentialScope +
		", SignedHeaders=" + strings.Join(signedHeaders, ";") +
		", Signature=" + signature
	request.Header.Set("Authorization", authorization)

	// 发起请求
	requestRaw, err := httputil.DumpRequest(request, true)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("dump request err: %w", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("do request err: %w", err)
	}
	// response.Body
	defer response.Body.Close()

	responseRaw, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("dump response err: %w", err)
	}
	return requestRaw, responseRaw, response.StatusCode, err
}
