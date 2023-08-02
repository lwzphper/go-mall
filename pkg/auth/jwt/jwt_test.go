package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	jwtSignKey = []byte("aomKPd2HzKTNuEiX5tc5IlriiUIb9IEBvAI0jMGzOQEp38yfL0cJibvQMYTLpxoX")
	wantStr    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnby1tYWxsIiwic3ViIjoiMSIsImV4cCI6MTY5MDk3MDM1NiwiaWF0IjoxNjkwOTcwMzU2fQ.zYOgu0uCwCx9D-UZqxuQSJrWZlp60LiwtEWSv09NSXo"
)

func TestGenerateToken(t *testing.T) {

	cases := []struct {
		name      string
		signKey   []byte
		now       time.Time
		want      string
		wantEqual bool
	}{
		{
			name:      "valid_token",
			signKey:   jwtSignKey,
			now:       time.Unix(1690970356, 0), // 固定时间，防止随机值生成的token变化
			want:      wantStr,
			wantEqual: true,
		},
		{
			name:      "valid_token",
			signKey:   jwtSignKey,
			now:       time.Now(),
			want:      wantStr,
			wantEqual: false,
		},
	}

	for _, c := range cases {
		jwtGen := NewJwtTokenGen("go-mall", c.signKey)
		jwtGen.nowFunc = func() time.Time {
			return c.now
		}
		token, err := jwtGen.GenerateToken("1")
		if err != nil {
			t.Errorf("token generate error: %v", err)
		}

		if c.wantEqual {
			assert.Equal(t, c.want, token)
		}

		if !c.wantEqual {
			assert.NotEqual(t, c.want, token)
		}
	}

}
