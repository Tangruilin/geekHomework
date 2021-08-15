package Week5_Go_Limiting

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Reply struct {
	reply bool
}

type SlideWindow struct {
	bucket       int        //设定滑动窗口的桶的个数，这里时需要进行预置的
	data         *list.List //使用链表表示实际的窗口，链表中的的每一个数据表示
	flow         int
	bucketkey    int //当前桶内流量
	maxFlow      int //窗口最大流量
	sync.RWMutex     //支持并发读写
}

func newSlideWindow(bucket int, maxflow int) *SlideWindow {
	return &SlideWindow{
		bucket:  bucket,
		maxFlow: maxflow,
		data:    new(list.List),
	}
}

func (s *SlideWindow) update(space time.Duration) {
	ticker := time.NewTicker(space)
	//窗口定时更新
	go func() {
		for _ = range ticker.C {
			s.Lock()
			if s.data.Len() <= s.bucket {
				s.data.PushBack(s.bucketkey)
				s.bucketkey = 0
			}
			if s.data.Len() > s.bucket {
				for i := 0; s.data.Len()-i > s.bucket; i++ {
					s.flow = s.flow - s.data.Front().Value.(int)
					s.data.Remove(s.data.Front())
				}
			}
			fmt.Println("当前窗口内流量： ", s.flow, "头部桶流量", s.data.Front().Value)
			s.Unlock()
		}
	}()
}

// GetRequest 模拟一个请求的到达, reply表示请求的返回为响应请求还是不响应请求
func (s *SlideWindow) GetRequest() *Reply {
	s.Lock()
	if s.flow >= s.maxFlow {
		return &Reply{reply: false}
	}
	s.flow++
	s.bucketkey++
	s.Unlock()
	return &Reply{reply: true}
}

// MultipleRequest 方便测试
func (s *SlideWindow) MultipleRequest(n int) []*Reply {
	reply := make([]*Reply, 0, n)
	for i := 0; i < n; i++ {
		reply = append(reply, s.GetRequest())
	}
	return reply
}
