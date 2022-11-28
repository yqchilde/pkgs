package sign

const (
	// KeyNameTimeStamp timestamp field name
	KeyNameTimeStamp = "timestamp"

	// KeyNameNonceStr nonce field
	KeyNameNonceStr = "nonce_str"

	// KeyNameAppID app id field
	KeyNameAppID = "app_id"

	// KeyNameSign sign field
	KeyNameSign = "sign"
)

// DefaultKeyName fields required for signature
type DefaultKeyName struct {
	Timestamp string
	NonceStr  string
	AppID     string
	Sign      string
}

func newDefaultKeyName() *DefaultKeyName {
	return &DefaultKeyName{
		Timestamp: KeyNameTimeStamp,
		NonceStr:  KeyNameNonceStr,
		AppID:     KeyNameAppID,
		Sign:      KeyNameSign,
	}
}

func (d *DefaultKeyName) SetKeyNameTimestamp(name string) {
	d.Timestamp = name
}

func (d *DefaultKeyName) SetKeyNameNonceStr(name string) {
	d.NonceStr = name
}

func (d *DefaultKeyName) SetKeyNameAppID(name string) {
	d.AppID = name
}

func (d *DefaultKeyName) SetKeyNameSign(name string) {
	d.Sign = name
}
