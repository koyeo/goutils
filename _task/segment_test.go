package _task

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testCalcSegment struct {
	firstId  int64
	lastId   int64
	routines int64
}

func TestNewTask(t *testing.T) {
	testCalcSegments(t, []*testCalcSegment{
		{firstId: 1, lastId: 9, routines: 1},
		{firstId: 1, lastId: 9, routines: 2},
		{firstId: 1, lastId: 9, routines: 3},
		{firstId: 1, lastId: 9, routines: 4},
		{firstId: 1, lastId: 9, routines: 5},
		{firstId: 1, lastId: 9, routines: 6},
		{firstId: 1, lastId: 9, routines: 7},
		{firstId: 1, lastId: 9, routines: 8},
		{firstId: 1, lastId: 9, routines: 9},
		{firstId: 2, lastId: 9, routines: 1},
		{firstId: 2, lastId: 9, routines: 2},
		{firstId: 2, lastId: 9, routines: 3},
		{firstId: 2, lastId: 9, routines: 4},
		{firstId: 2, lastId: 9, routines: 5},
		{firstId: 2, lastId: 9, routines: 6},
		{firstId: 2, lastId: 9, routines: 7},
		{firstId: 2, lastId: 9, routines: 8},
		{firstId: 2, lastId: 9, routines: 9},
		{firstId: 3, lastId: 9, routines: 1},
		{firstId: 3, lastId: 9, routines: 2},
		{firstId: 3, lastId: 9, routines: 3},
		{firstId: 3, lastId: 9, routines: 4},
		{firstId: 3, lastId: 9, routines: 5},
		{firstId: 3, lastId: 9, routines: 6},
		{firstId: 3, lastId: 9, routines: 7},
		{firstId: 3, lastId: 9, routines: 8},
		{firstId: 3, lastId: 9, routines: 9},
		{firstId: 4, lastId: 9, routines: 1},
		{firstId: 4, lastId: 9, routines: 2},
		{firstId: 4, lastId: 9, routines: 3},
		{firstId: 4, lastId: 9, routines: 4},
		{firstId: 4, lastId: 9, routines: 5},
		{firstId: 4, lastId: 9, routines: 6},
		{firstId: 4, lastId: 9, routines: 7},
		{firstId: 4, lastId: 9, routines: 8},
		{firstId: 4, lastId: 9, routines: 9},
		{firstId: 4, lastId: 9, routines: 1},
		{firstId: 4, lastId: 9, routines: 2},
		{firstId: 4, lastId: 9, routines: 3},
		{firstId: 4, lastId: 9, routines: 4},
		{firstId: 4, lastId: 9, routines: 5},
		{firstId: 4, lastId: 9, routines: 6},
		{firstId: 4, lastId: 9, routines: 7},
		{firstId: 4, lastId: 9, routines: 8},
		{firstId: 4, lastId: 9, routines: 9},
		{firstId: 5, lastId: 9, routines: 1},
		{firstId: 5, lastId: 9, routines: 2},
		{firstId: 5, lastId: 9, routines: 3},
		{firstId: 5, lastId: 9, routines: 4},
		{firstId: 5, lastId: 9, routines: 5},
		{firstId: 5, lastId: 9, routines: 6},
		{firstId: 5, lastId: 9, routines: 7},
		{firstId: 5, lastId: 9, routines: 8},
		{firstId: 5, lastId: 9, routines: 9},
		{firstId: 6, lastId: 9, routines: 1},
		{firstId: 6, lastId: 9, routines: 2},
		{firstId: 6, lastId: 9, routines: 3},
		{firstId: 6, lastId: 9, routines: 4},
		{firstId: 6, lastId: 9, routines: 5},
		{firstId: 6, lastId: 9, routines: 6},
		{firstId: 6, lastId: 9, routines: 7},
		{firstId: 6, lastId: 9, routines: 8},
		{firstId: 6, lastId: 9, routines: 9},
		{firstId: 7, lastId: 9, routines: 1},
		{firstId: 7, lastId: 9, routines: 2},
		{firstId: 7, lastId: 9, routines: 3},
		{firstId: 7, lastId: 9, routines: 4},
		{firstId: 7, lastId: 9, routines: 5},
		{firstId: 7, lastId: 9, routines: 6},
		{firstId: 7, lastId: 9, routines: 7},
		{firstId: 7, lastId: 9, routines: 8},
		{firstId: 7, lastId: 9, routines: 9},
		{firstId: 8, lastId: 9, routines: 1},
		{firstId: 8, lastId: 9, routines: 2},
		{firstId: 8, lastId: 9, routines: 3},
		{firstId: 8, lastId: 9, routines: 4},
		{firstId: 8, lastId: 9, routines: 5},
		{firstId: 8, lastId: 9, routines: 6},
		{firstId: 8, lastId: 9, routines: 7},
		{firstId: 8, lastId: 9, routines: 8},
		{firstId: 8, lastId: 9, routines: 9},
		{firstId: 8, lastId: 9, routines: 1},
		{firstId: 9, lastId: 9, routines: 2},
		{firstId: 9, lastId: 9, routines: 3},
		{firstId: 9, lastId: 9, routines: 4},
		{firstId: 9, lastId: 9, routines: 5},
		{firstId: 9, lastId: 9, routines: 6},
		{firstId: 9, lastId: 9, routines: 7},
		{firstId: 9, lastId: 9, routines: 8},
		{firstId: 9, lastId: 9, routines: 9},
	})
}

func testCalcSegments(t *testing.T, tests []*testCalcSegment) {
	for _, test := range tests {
		require.True(t, test.routines > 0, test)
		segments, err := CalcSegments(test.firstId, test.lastId, test.routines)
		require.NoError(t, err)
		
		//d, _ := json.Marshal(segments)
		//fmt.Println(test, string(d))
		
		require.Equal(t, test.firstId-1, segments[0].StartId, test)
		require.Equal(t, test.lastId, segments[len(segments)-1].EndId, test)
		require.Equal(t, test.lastId-test.firstId, segments[len(segments)-1].EndId-segments[0].StartId-1, test)
		for i, v := range segments {
			if i > 0 {
				require.Equal(t, v.StartId, segments[i-1].EndId, test)
			}
		}
		
	}
}
