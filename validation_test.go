package valid

import "testing"

func TestValidation_Valid(t *testing.T) {
	type args struct {
		obj interface{}
	}
	st := struct {
		A string `valid:"Max(10);ErrorCode(1001)"`
	}{
		A: "11",
	}

	tests := []struct {
		name     string
		args     args
		wantB    bool
		wantCode int64
		wantErr  bool
	}{
		{
			name: "valid:校验成功",
			args: args{
				st,
			},
			wantB:    true,
			wantCode: 0,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validation{}
			gotB, gotCode, err := v.Valid(tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validation.Valid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("Validation.Valid() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotCode != tt.wantCode {
				t.Errorf("Validation.Valid() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
