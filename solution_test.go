package square

import "testing"

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen  float64
		sidesNum numSides
	}
	tests := []struct {
		name       string
		args       args
		wantSquare float64
	}{
		{"Triangle", args{10.0, SidesTriangle}, 43.30127018922193},
		{"Square", args{10.0, SidesSquare}, 100.0},
		{"Circle", args{10.0, SidesCircle}, 314.1592653589793},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSquare := CalcSquare(tt.args.sideLen, tt.args.sidesNum); gotSquare != tt.wantSquare {
				t.Errorf("CalcSquare() = %v, want %v", gotSquare, tt.wantSquare)
			}
		})
	}
}
