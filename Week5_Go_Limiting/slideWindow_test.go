package Week5_Go_Limiting

import (
	"testing"
	"time"
)

func TestSlideWindow(t *testing.T) {
	mySliceWindow := newSlideWindow(6, 120)
	mySliceWindow.update(10 * time.Millisecond) //每10ms对窗口进行一次滑动
	mySliceWindow.MultipleRequest(10)
	time.Sleep(5 * time.Millisecond)
	mySliceWindow.MultipleRequest(20)
	time.Sleep(20 * time.Millisecond)
	mySliceWindow.MultipleRequest(10)
	time.Sleep(80 * time.Millisecond)
}
