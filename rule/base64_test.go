package rule

import "testing"

func TestBase64Rule_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "base64标准检验成功",
			fields: fields{
				Value: "c3VjaHVhbmdqaUBnbWFpbC5jb20=",
			},
			wantErr: false,
		},
		{
			name: "base64标准检验失败",
			fields: fields{
				Value: "suchuangji@gmail.com",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Base64Rule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("Base64.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
