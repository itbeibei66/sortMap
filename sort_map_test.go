package sortMap

import (
	"crypto/rand"
	"github.com/emirpasic/gods/maps/treemap"
	"math/big"
	"sort"
	"testing"
	"time"
)

func TestSortMap(t *testing.T) {
	var allTime1, allTime2 int64
	var allTime3, allTime4 int64
	var allTime5, allTime6 int64
	for tt := 0; tt < 100; tt++ {
		m := NewSortMap()
		m2 := treemap.NewWithIntComparator()
		arr := make([]int, 0)
		for i := 0; i < 100000; i++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(500000))
			arr = append(arr, int(n.Int64()))
		}
		cos1 := time.Now().UnixMilli()
		for i := 0; i < 10000; i++ {
			m.Add(int64(arr[i]), 1)
		}
		end1 := time.Now().UnixMilli()
		allTime1 += end1 - cos1
		cos2 := time.Now().UnixMilli()
		for i := 0; i < 10000; i++ {
			m2.Put(arr[i], 1)
		}
		end2 := time.Now().UnixMilli()
		allTime2 += end2 - cos2

		for i := 0; i < 100000; i++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(500000))

			c1 := time.Now().UnixMilli()
			m.SearchLeftKey(n.Int64())
			c2 := time.Now().UnixMilli()
			allTime3 += c2 - c1

			c1 = time.Now().UnixMilli()
			m2.Floor(int(n.Int64()))
			c2 = time.Now().UnixMilli()
			allTime4 += c2 - c1

			c1 = time.Now().UnixMilli()
			m.SearchRightKey(n.Int64())
			c2 = time.Now().UnixMilli()
			allTime5 += c2 - c1

			c1 = time.Now().UnixMilli()
			m2.Ceiling(int(n.Int64()))
			c2 = time.Now().UnixMilli()
			allTime6 += c2 - c1
		}
	}
	t.Logf("%d, %d", allTime1, allTime2)
	t.Logf("%d, %d", allTime3, allTime4)
	t.Logf("%d, %d", allTime5, allTime6)
}

func TestSortMap2(t *testing.T) {
	m := NewSortMap()
	arr := make([]int, 0)
	for i := 0; i < 10000000; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(500000000))
		arr = append(arr, int(n.Int64()))
	}
	for i := 0; i < len(arr); i++ {
		m.Add(int64(arr[i]), 1)
	}

	c1 := 0
	c2 := 0
	var dd1, dd2 []int64
	for i := 0; i < 10; i++ {
		b2, _ := rand.Int(rand.Reader, big.NewInt(500000000))
		b1, _ := rand.Int(rand.Reader, big.NewInt(b2.Int64()))
		s := time.Now().UnixMilli()
		dd1 = make([]int64, 0)
		each := m.NewIterator()
		each = each.BeginWith(b1.Int64())
		for each.cur != nil && each.Key() <= b2.Int64() {
			dd1 = append(dd1, each.Key())
			each.Next()
		}
		//m.GetRangeKey(b1.Int64(), b2.Int64())
		//fmt.Println(m.GetRangeKey(b1.Int64(), b2.Int64()))
		e := time.Now().UnixMilli()
		c1 += int(e - s)
		s = time.Now().UnixMilli()
		dd2 = make([]int64, 0)
		for _, v := range m.GetRangeKey(b1.Int64(), b2.Int64()) {
			dd2 = append(dd2, v)
		}
		e = time.Now().UnixMilli()
		c2 += int(e - s)
	}
	t.Logf("%d %d", c1, c2)

}

func TestSortMap3(t *testing.T) {
	arr := []int{1, 3, 115, 117,
		119, 121, 123, 125, 127, 131,
		133, 135, 139, 141, 145, 147, 149,
		151, 155, 157, 159, 161, 163, 165, 171, 173, 175, 177, 179,
		181, 183, 185, 187}
	m := NewSortMap()
	for i := 0; i < len(arr); i++ {
		m.Add(int64(arr[i]), 1)
	}
	m.Delete(3)
	t.Logf("%v", m.GetRangeKey(127, 139))
}
func find(arr []int, ta int) (int64, int64, int64, int64) {
	ans1 := -999999
	ans2 := -999999
	for i := 0; i < len(arr); i++ {
		if arr[i] >= ta {
			ans2 = arr[i]
			break
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= ta {
			ans1 = arr[i]
			break
		}
	}
	return int64(ans1), int64(ans2), int64(arr[len(arr)-1]), int64(arr[0])
}

func find2(arr []int, begin int, end int) []int {
	m := make(map[int]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	arr2 := make([]int, 0)
	for k := range m {
		arr2 = append(arr2, k)
	}
	sort.Ints(arr2)
	res := make([]int, 0)
	for _, v := range arr2 {
		if v >= begin && v <= end {
			res = append(res, v)
		}
	}
	return res
}
