package _task

import "fmt"

type Segment struct {
	StartId int64 `json:"start_id"`
	EndId   int64 `json:"end_id"`
}

func (p Segment) Copy() *Segment {
	return &Segment{
		StartId: p.StartId,
		EndId:   p.EndId,
	}
}

func CalcSegments(firstId, lastId, routines int64) (segments []*Segment, err error) {
	
	if lastId < firstId {
		err = fmt.Errorf("lastId less than firstId")
	}
	
	if firstId == lastId || (lastId-firstId) <= routines {
		segments = append(segments, &Segment{
			StartId: firstId - 1,
			EndId:   lastId,
		})
		return
	}
	
	a := (lastId - firstId) / routines
	for i := int64(0); i < routines; i++ {
		var startId, endId int64
		if i == 0 {
			startId = firstId - 1
		} else {
			startId = firstId + (i * a)
		}
		
		if i == routines-1 {
			endId = lastId
		} else {
			endId = (firstId) + ((i + 1) * a)
		}
		segments = append(segments, &Segment{
			StartId: startId,
			EndId:   endId,
		})
	}
	
	return
}
