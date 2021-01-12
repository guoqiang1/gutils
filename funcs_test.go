package gutils

import (
	"reflect"
	"testing"
)

func TestMd5(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{data: "123456"},
			want: "e10adc3949ba59abbe56e057f20f883e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5(tt.args.data); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrc32(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{data: "123456"},
			want: 158520161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc32(tt.args.data); got != tt.want {
				t.Errorf("Crc32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandString10(t *testing.T) {
	t.Log(RandString10(6))
}

func TestRandString62(t *testing.T) {
	t.Log(RandString62(16))
}

func TestFirstDayOfNowMonth(t *testing.T) {
	t.Log(FirstDayOfNowMonth())
}

func TestDateTimePower(t *testing.T) {
	t.Log(DateTimePower("Y-m-d H:i:s", 0))
}

func TestFirstDayOfNowWeek(t *testing.T) {
	t.Log(FirstDayOfNowWeek())
}

func TestArrayUniqueInt64(t *testing.T) {
	type args struct {
		data []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{data: []int64{2, 2, 0, 12, 32, 12}},
			want: []int64{2, 0, 12, 32},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUniqueInt64(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUniqueInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUniqueString(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{data: []string{"abc", "ef", "你好", "d", "你", "你好", "abc"}},
			want: []string{"abc", "ef", "你好", "d","你"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUniqueString(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUniqueString() = %v, want %v", got, tt.want)
			}
		})
	}
}
