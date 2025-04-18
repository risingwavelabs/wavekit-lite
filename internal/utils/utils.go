package utils

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var CurrentVersion string

func IfElse[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// Ptr returns a pointer to the value passed in.
// The pointer is to a shallow copy, not the original value.
func Ptr[T any](v T) *T {
	var val = v
	return &val
}

func GenerateCode() string {
	return fmt.Sprintf("%d", (1+rand.Intn(10))*10000+rand.Intn(10000))
}

func HashPassword(password string, salt string) (string, error) {
	preHashed := fmt.Sprintf("%s-%s", password, salt)
	h := sha256.New()
	_, err := h.Write([]byte(preHashed))
	if err != nil {
		return "", err
	}
	bs := h.Sum(nil)
	hashed := fmt.Sprintf("%x", bs)
	return hashed, nil
}

func GenerateHashAndSalt(password string) (string, string, error) {
	salt := fmt.Sprintf("salt-%d", rand.Int31())
	hashedPassword, err := HashPassword(password, salt)
	if err != nil {
		return "", "", err
	}
	return salt, hashedPassword, nil
}

func JSONConvert[T any, U any](v T, u *U) error {
	raw, err := json.Marshal(v)
	if err != nil {
		return errors.Wrapf(err, "JSONConvert failed")
	}
	return json.Unmarshal(raw, u)
}

func TryMarshal(o any) string {
	raw, err := json.Marshal(o)
	if err != nil {
		return "<unmarshalable>"
	}
	return string(raw)
}

func Unwrap[T any](v *T) T {
	if v == nil {
		var zero T
		return zero
	}
	return *v
}

func UnwrapOrDefault[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

func TestTCPConnection(ctx context.Context, host string, port int32, timeout time.Duration) error {
	var d net.Dialer
	d.Timeout = timeout

	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func TruncateString(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

// renderProgressBar generates a string representation of a progress bar
// based on the given progress percentage (0.0 to 1.0)
func RenderProgressBar(progress float64) string {
	barWidth := 40
	completedWidth := int(float64(barWidth) * progress)

	// Build the progress bar string
	progressBar := "["
	for i := 0; i < barWidth; i++ {
		if i < completedWidth {
			progressBar += "="
		} else if i == completedWidth && progress < 1.0 {
			progressBar += ">"
		} else {
			progressBar += " "
		}
	}
	progressBar += "]"

	// Return the combined progress string with percentage
	return fmt.Sprintf("%s %.2f%%", progressBar, progress*100)
}

func ParseDuration(s string) (time.Duration, error) {
	if s == "" {
		return 0, errors.New("empty duration")
	}
	if strings.HasSuffix(s, "d") {
		days, err := strconv.Atoi(s[:len(s)-1])
		if err != nil {
			return 0, errors.Wrap(err, "invalid duration")
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(s)
}

// RetrieveFromJSON uses json.Decoder to parse only one field in a streaming way
func RetrieveFromJSON[T any](s string, targetKey string) (*T, error) {
	decoder := json.NewDecoder(strings.NewReader(s))

	// Read opening brace
	if _, err := decoder.Token(); err != nil {
		return nil, err
	}

	var ret *T
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		if key, ok := token.(string); ok && key == targetKey {
			if token, err = decoder.Token(); err != nil {
				return nil, err
			}
			if val, ok := token.(T); ok {
				ret = &val
				break
			}
		} else {
			// Skip the value for this field
			if _, err := decoder.Token(); err != nil {
				return nil, err
			}
		}
	}

	if ret == nil {
		return nil, errors.Errorf("missing attribute: %s", targetKey)
	}

	return ret, nil
}
