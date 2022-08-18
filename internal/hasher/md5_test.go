package hasher

import (
	"encoding/hex"
	"reflect"
	"testing"
)

const (
	test1 = "kratos is a good framework"
)

func TestMd5HashOutput(t *testing.T) {
	hash := Md5Hash([]byte(""))
	encoded := hex.EncodeToString(hash)
	t.Log(encoded)
}

func TestMd5Hash(t *testing.T) {
	want, _ := hex.DecodeString("ceb1afafca05ef9cd4c11e3d042b76bc")
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"tes1", args{data: []byte(test1)}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Hash(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Md5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}
