package array

import (
	"testing"
)

func TestFormat(t *testing.T) {
	var id = []int{1, 2, 3, 4}
	d := FormatSliceToMapKey(id)
	if _, ok := d[1]; ok == false {
		t.Error("格式错误")
	}

	if _, ok := d[2]; ok == false {
		t.Error("格式错误")
	}

	if _, ok := d[3]; ok == false {
		t.Error("格式错误")
	}

	if _, ok := d[4]; ok == false {
		t.Error("格式错误")
	}
}

func TestInInt64Slice(t *testing.T) {
	type args struct {
		element int64
		slice   []int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				element: 8,
				slice:   []int64{1, 3, 9, 8, 7},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				element: 10,
				slice:   []int64{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int64{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int64{1, 3, 9, 8, 7, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InInt64Slice(tt.args.element, tt.args.slice); got != tt.want {
				t.Errorf("InIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInIntSlice(t *testing.T) {
	type args struct {
		element int
		slice   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				element: -8,
				slice:   []int{1, 3, 9, -8, 7},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				element: 10,
				slice:   []int{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int{1, 3, 9, 8, 7, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InIntSlice(tt.args.element, tt.args.slice); got != tt.want {
				t.Errorf("InIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInInt16Slice(t *testing.T) {
	type args struct {
		element int16
		slice   []int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				element: -8,
				slice:   []int16{1, 3, 9, -8, 7},
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				element: 10,
				slice:   []int16{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int16{1, 3, 9, 8, 7},
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				element: 0,
				slice:   []int16{1, 3, 9, 8, 7, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InInt16Slice(tt.args.element, tt.args.slice); got != tt.want {
				t.Errorf("InInt16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
