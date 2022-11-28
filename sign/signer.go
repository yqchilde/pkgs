package sign

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/yqchilde/pkgs/utils"
)

type CryptoFunc func(secretKey string, args string) []byte

type Signer struct {
	*DefaultKeyName

	body       url.Values // signature parameter body
	bodyPrefix string     // parameter body prefix
	bodySuffix string     // parameter body suffix
	splitChar  string     // prefix, suffix separator

	secretKey  string
	cryptoFunc CryptoFunc
}

func NewSigner(cryptoFunc CryptoFunc) *Signer {
	return &Signer{
		DefaultKeyName: newDefaultKeyName(),
		body:           make(url.Values),
		bodyPrefix:     "",
		bodySuffix:     "",
		splitChar:      "",
		cryptoFunc:     cryptoFunc,
	}
}

func NewSignerMd5() *Signer {
	return NewSigner(Md5Sign)
}

func NewSignerHmac() *Signer {
	return NewSigner(HmacSign)
}

// SetBody Set the entire parameter body object
func (s *Signer) SetBody(body url.Values) {
	for k, v := range body {
		s.body[k] = v
	}
}

// GetBody Return body content
func (s *Signer) GetBody() url.Values {
	return s.body
}

// AddBody Add signature body field and value
func (s *Signer) AddBody(key string, value string) *Signer {
	return s.AddBodies(key, []string{value})
}

// AddBodies add value to body
func (s *Signer) AddBodies(key string, value []string) *Signer {
	s.body[key] = value
	return s
}

// SetTimeStamp Set timestamp parameters
func (s *Signer) SetTimeStamp(ts int64) *Signer {
	return s.AddBody(s.Timestamp, strconv.FormatInt(ts, 10))
}

// GetTimeStamp Get timestamp
func (s *Signer) GetTimeStamp() string {
	return s.body.Get(s.Timestamp)
}

// SetNonceStr Set random string parameters
func (s *Signer) SetNonceStr(nonce string) *Signer {
	return s.AddBody(s.NonceStr, nonce)
}

// GetNonceStr Return nonce string
func (s *Signer) GetNonceStr() string {
	return s.body.Get(s.NonceStr)
}

// SetAppID Set app parameters
func (s *Signer) SetAppID(appID string) *Signer {
	return s.AddBody(s.AppID, appID)
}

// GetAppID get app id
func (s *Signer) GetAppID() string {
	return s.body.Get(s.AppID)
}

// RandNonceStr Automatically generate 16-bit random string parameters
func (s *Signer) RandNonceStr() *Signer {
	return s.SetNonceStr(utils.RandomStr(16))
}

// SetSignBodyPrefix Set the prefix string of the signature string
func (s *Signer) SetSignBodyPrefix(prefix string) *Signer {
	s.bodyPrefix = prefix
	return s
}

// SetSignBodySuffix Set the suffix string of the signature string
func (s *Signer) SetSignBodySuffix(suffix string) *Signer {
	s.bodySuffix = suffix
	return s
}

// SetSplitChar Set the separator between prefix, suffix and signature body. The default is an empty string
func (s *Signer) SetSplitChar(split string) *Signer {
	s.splitChar = split
	return s
}

// SetAppSecret Set the signing key
func (s *Signer) SetAppSecret(appSecret string) *Signer {
	s.secretKey = appSecret
	return s
}

// SetAppSecretWrapBody t the head and tail of the signature parameter body, splice the App Secret string
func (s *Signer) SetAppSecretWrapBody(appSecret string) *Signer {
	s.SetSignBodyPrefix(appSecret)
	s.SetSignBodySuffix(appSecret)
	return s.SetAppSecret(appSecret)
}

// GetSignBodyString Get the original string used for signing
func (s *Signer) GetSignBodyString() string {
	return s.MakeRawBodyString()
}

// MakeRawBodyString Get the original string used for signing
func (s *Signer) MakeRawBodyString() string {
	return s.bodyPrefix + s.splitChar + s.getSortedBodyString() + s.splitChar + s.bodySuffix
}

// GetSignedQuery Get the query string with signed parameters
func (s *Signer) GetSignedQuery() string {
	return s.MakeSignedQuery()
}

// MakeSignedQuery Get a string with signed parameters
func (s *Signer) MakeSignedQuery() string {
	body := s.getSortedBodyString()
	sign := s.GetSignature()
	return body + "&" + s.Sign + "=" + sign
}

// GetSignature Get signature
func (s *Signer) GetSignature() string {
	return s.MakeSign()
}

// MakeSign Generate signature
func (s *Signer) MakeSign() string {
	sign := fmt.Sprintf("%x", s.cryptoFunc(s.secretKey, s.GetSignBodyString()))
	return sign
}

func (s *Signer) getSortedBodyString() string {
	return SortKVPairs(s.body)
}

// SortKVPairs Concatenate the key-value pairs of the Map into strings in lexicographical order
func SortKVPairs(m url.Values) string {
	size := len(m)
	if size == 0 {
		return ""
	}
	keys := make([]string, size)
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	pairs := make([]string, size)
	for i, key := range keys {
		pairs[i] = key + "=" + strings.Join(m[key], ",")
	}
	return strings.Join(pairs, "&")
}
