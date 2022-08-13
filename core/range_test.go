package core_test

import (
	"testing"

	"github.com/whuangz88/calendly/core"
)

func TestStart(t *testing.T) {
	type fields struct {
		StartSecond int
		EndSecond   int
	}

	tests := []struct {
		name      string
		fields    fields
		wantStart string
		wantEnd   string
	}{
		{
			name: "1 o'clock",
			fields: fields{
				StartSecond: 3600,
				EndSecond:   7200,
			},
			wantStart: "1:00",
			wantEnd:   "2:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := core.Range{
				StartSecond: tt.fields.StartSecond,
				EndSecond:   tt.fields.EndSecond,
			}

			assert.Equal(t, tt.wantStart, r.Start())
			assert.Equal(t, tt.wantEnd, r.Start())

		})
	}
}
