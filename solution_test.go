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
		{"Triangle10", args{10.0, SidesTriangle}, 43.30127018922193},
		{"Triangle1", args{1.0, SidesTriangle}, 0.4330127018922193},
		{"Triangle2", args{2.0, SidesTriangle}, 1.7320508075688772},
		{"Square10", args{10.0, SidesSquare}, 100.0},
		{"Circle1", args{1.0, SidesCircle}, 3.141592653589793},
		{"Circle10", args{10.0, SidesCircle}, 314.1592653589793},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSquare := CalcSquare(tt.args.sideLen, tt.args.sidesNum); gotSquare != tt.wantSquare {
				t.Errorf("CalcSquare() = %v, want %v", gotSquare, tt.wantSquare)
			}
		})
	}
}
