package webcrawler

import (
	"sort"
	"testing"
)

func Test_muxEntrySlice_sortable(t *testing.T) {
	val1 := muxEntry{
		pattern: "1",
	}
	val2 := muxEntry{
		pattern: "123",
	}
	slice := []muxEntry{val2, val1}
	sort.Sort(muxEntrySlice(slice))
	expect(t, slice[0].pattern, "1")
	expect(t, slice[1].pattern, "123")
	//结构体判等
	expect(t, slice[0], val1)
	expect(t, slice[1], val2)
}

func Test_Spiderlair(t *testing.T) {
	var spl = NewSpiderlair()

	val1 := muxEntry{
		pattern: "http://www.baidu.com",
	}
	val2 := muxEntry{
		pattern: "123",
	}

	val3 := muxEntry{
		pattern: "1",
	}
	val4 := muxEntry{
		pattern: "123",
	}
}
