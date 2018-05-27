package drawing

import (
	"math"
	"testing"
)

func TestForward(t *testing.T) {

	type args struct {
		startX         float64
		startY         float64
		startDirection float64
		numSpaces      int
	}
	tests := []struct {
		name  string
		args  args
		wantX float64
		wantY float64
	}{
		{
			name: "straight up",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 0,
				numSpaces:      4,
			},
			wantX: 10,
			wantY: 6,
		},
		{
			name: "straight down",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 180,
				numSpaces:      4,
			},
			wantX: 10,
			wantY: 14,
		},
		{
			name: "straight right",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 90,
				numSpaces:      4,
			},
			wantX: 14,
			wantY: 10,
		},
		{
			name: "straight left",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 270,
				numSpaces:      4,
			},
			wantX: 6,
			wantY: 10,
		},
		{
			name: "NE",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 45,
				numSpaces:      5,
			},
			wantX: 13.53,
			wantY: 6.47,
		},
		{
			name: "NNE",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 15,
				numSpaces:      5,
			},
			wantX: 1.29,
			wantY: -4.83,
		},
		{
			name: "NEE",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 75,
				numSpaces:      5,
			},
			wantX: 4.83,
			wantY: -1.29,
		},

		{
			name: "SE",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 135,
				numSpaces:      5,
			},
			wantX: 13.53,
			wantY: 13.53,
		},
		{
			name: "SEE",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 105,
				numSpaces:      5,
			},
			wantX: 4.83,
			wantY: 1.29,
		},
		{
			name: "SSE",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 165,
				numSpaces:      5,
			},
			wantX: 1.29,
			wantY: 4.83,
		},

		{
			name: "NW",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 315,
				numSpaces:      5,
			},
			wantX: 6.47,
			wantY: 6.47,
		},
		{
			name: "NNW",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 345,
				numSpaces:      5,
			},
			wantX: 8.71,
			wantY: 5.17,
		},
		{
			name: "NWW",
			args: args{
				startX:         10,
				startY:         10,
				startDirection: 285,
				numSpaces:      5,
			},

			wantX: 5.17,
			wantY: 8.71,
		},

		{
			name: "SW",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 225,
				numSpaces:      5,
			},
			wantX: -3.54,
			wantY: 3.54,
		},
		{
			name: "SSW",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 195,
				numSpaces:      5,
			},
			wantX: -1.29,
			wantY: 4.83,
		},
		{
			name: "SWW",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 255,
				numSpaces:      5,
			},
			wantX: -4.83,
			wantY: 1.29,
		},

		{
			name: "line",
			args: args{
				startX:         0,
				startY:         0,
				startDirection: 150,
				numSpaces:      5,
			},
			wantX: 2.5,
			wantY: 4.33,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(1000, 1000)
			c.multiplier = 1

			c.currentX = tt.args.startX
			c.currentY = tt.args.startY
			c.currentDirection = tt.args.startDirection

			c.Forward(tt.args.numSpaces)

			if !floatEquals(tt.wantX, c.currentX) {
				t.Errorf("Expected x %.2f got %.2f", tt.wantX, c.currentX)
			}

			if !floatEquals(tt.wantY, c.currentY) {
				t.Errorf("Expected y %.2f got %.2f", tt.wantY, c.currentY)
			}
		})
	}
}

var eps float64 = 0.01

func floatEquals(a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func TestTurn(t *testing.T) {
	type args struct {
		startDirection float64
		degrees        int
	}
	tests := []struct {
		name       string
		args       args
		wantDegree float64
	}{
		{
			name: "turn right",
			args: args{
				startDirection: 0,
				degrees:        45,
			},
			wantDegree: 45,
		},
		{
			name: "turn left",
			args: args{
				startDirection: 0,
				degrees:        -45,
			},
			wantDegree: 315,
		},
		{
			name: "turn right far",
			args: args{
				startDirection: 0,
				degrees:        115,
			},
			wantDegree: 115,
		},
		{
			name: "turn left far",
			args: args{
				startDirection: 0,
				degrees:        -115,
			},
			wantDegree: 245,
		},
		{
			name: "turn right from other angle",
			args: args{
				startDirection: 90,
				degrees:        10,
			},
			wantDegree: 100,
		},
		{
			name: "turn left from other angle",
			args: args{
				startDirection: 90,
				degrees:        -10,
			},
			wantDegree: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCanvas(1000, 1000)

			c.currentDirection = tt.args.startDirection

			c.Turn(tt.args.degrees)

			if !floatEquals(tt.wantDegree, c.currentDirection) {
				t.Errorf("want directiong %.2f, got %.2f", tt.wantDegree, c.currentDirection)
			}
		})
	}
}
