package rule

import (
	"testing"
)

func TestTelRule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Tel:标准检验成功",
			fields: fields{
				Value: "022-00008888",
			},
			wantErr: false,
		},
		{
			name: "Tel:标准检验成功",
			fields: fields{
				Value: "02270008888",
			},
			wantErr: false,
		},
		{
			name: "Tel:标准检验失败",
			fields: fields{
				Value: "222-70008888",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TelRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("TelRule.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
