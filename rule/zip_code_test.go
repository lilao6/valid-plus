package rule

import "testing"

func TestZipCode_Valid(t *testing.T) {
	type fields struct {
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ZipCode:标准检验成功",
			fields: fields{
				Value: "536000",
			},
			wantErr: false,
		},
		{
			name: "ZipCode:标准检验失败",
			fields: fields{
				Value: "00008888",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ZipCodeRule{
				value: tt.fields.Value,
			}
			if err := r.Valid(); (err != nil) != tt.wantErr {
				t.Errorf("ZipCode.Valid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
