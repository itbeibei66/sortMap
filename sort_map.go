package sortMap

import "strconv"

// 排序数据结构
type SortMap struct {
	head []*Node
}

func NewSortMap() *SortMap {
	res := &SortMap{
		head: make([]*Node, 20),
	}
	for k := range res.head {
		res.head[k] = newNode(0, 0, nil, k)
	}
	return res
}

// Node
type Node struct {
	key   int64
	value interface{}
	son   []*Node
	par   *Node
	index int
}

func newNode(key int64, val interface{}, par *Node, c int) *Node {
	return &Node{
		key:   key,
		value: val,
		par:   par,
		index: c,
	}
}

type dfsCon struct {
	data []int64
}

func (s *SortMap) Add(key int64, value interface{}) {
	str := strconv.FormatInt(key, 10)
	if key >= 0 {
		s.add(key, value, str, s.head[len(str)+9], 0)
	} else {
		s.add(key, value, str, s.head[10-len(str)], 0)
	}
}

func (s *SortMap) add(key int64, value interface{}, str string, node *Node, tag int) {
	cur := node
	for i := tag; i < len(str); i++ {
		c := int(str[i] - '0')
		if cur.son == nil {
			cur.son = make([]*Node, 10)
		}
		if cur.son[c] == nil {
			cur.son[c] = newNode(key, value, cur, c)
			return
		}
		if key > cur.son[c].key {
			cur = cur.son[c]
		} else if key < cur.son[c].key {
			compareKey := cur.son[c].key
			compareStr := strconv.FormatInt(compareKey, 10)
			compareValue := cur.son[c].value
			cur.son[c].key = key
			cur.son[c].value = value
			s.add(compareKey, compareValue, compareStr, cur.son[c], i+1)
			return
		} else {
			cur.son[c].value = value
			return
		}
	}
}

func (s *SortMap) Search(key int64) (interface{}, bool) {
	data, ok := s.search(key)
	if !ok {
		return nil, ok
	}
	return data.value, ok
}

func (s *SortMap) search(key int64) (*Node, bool) {
	str := strconv.FormatInt(key, 10)
	var cur *Node
	if key >= 0 {
		cur = s.head[len(str)+9]
	} else {
		str = str[1:]
		cur = s.head[10-len(str)]
	}
	for i := 0; i < len(str); i++ {
		c := int(str[i] - '0')
		if cur.son == nil || cur.son[c] == nil {
			return nil, false
		}
		if cur.son[c].key == key {
			return cur.son[c], true
		} else if cur.son[c].key > key {
			return nil, false
		}
		cur = cur.son[c]
	}
	return nil, false
}

func (s *SortMap) SearchLeftKey(key int64) (int64, bool) {
	return s.searchLeftKey(key)
}

func (s *SortMap) searchLeftKey(key int64) (int64, bool) {
	str := strconv.FormatInt(key, 10)
	var index int64
	if key >= 0 {
		index = int64(len(str) + 9)
	} else {
		str = str[1:]
		index = int64(10 - len(str))
	}
	cur := s.head[index]
	var lastFoundNum int64
	var lastFound bool
	for i := 0; i < len(str); i++ {
		c := int(str[i] - '0')
		if cur.son == nil {
			return s.searchLeftKey1(lastFoundNum, lastFound, index, cur)
		}
		if cur.son[c] == nil {
			return s.searchLeftKey2(lastFoundNum, lastFound, index, cur, c)
		}
		if cur.son[c].key == key {
			return cur.son[c].key, true
		} else if cur.son[c].key > key {
			return s.searchLeftKey2(lastFoundNum, lastFound, index, cur, c)
		} else {
			lastFoundNum = cur.son[c].key
			lastFound = true
			cur = cur.son[c]
		}
	}
	return lastFoundNum, lastFound
}

func (s *SortMap) searchLeftKey1(lastFoundNum int64, lastFound bool, index int64, cur *Node) (int64, bool) {
	if cur == s.head[index] {
		for j := index - 1; j >= 0; j-- {
			if s.head[j].son != nil {
				return s.peekMaxWithNode(s.head[j], j)
			}
		}
	} else {
		return cur.key, true
	}
	return lastFoundNum, lastFound
}

func (s *SortMap) searchLeftKey2(lastFoundNum int64, lastFound bool, index int64, cur *Node, c int) (int64, bool) {
	for k := c - 1; k >= 0; k-- {
		if cur.son[k] != nil {
			return s.peekMaxWithNode(cur.son[k], int64(k))
		}
	}
	if cur == s.head[index] {
		for j := index - 1; j >= 0; j-- {
			if s.head[j].son != nil {
				return s.peekMaxWithNode(s.head[j], j)
			}
		}
		return lastFoundNum, lastFound
	}
	return cur.key, true
}

func (s *SortMap) SearchRightKey(key int64) (int64, bool) {
	return s.searchRightKey(key)
}

func (s *SortMap) searchRightKey(key int64) (int64, bool) {
	str := strconv.FormatInt(key, 10)
	var index int64
	if key >= 0 {
		index = int64(len(str) + 9)
	} else {
		str = str[1:]
		index = int64(10 - len(str))
	}
	var cur *Node
	cur = s.head[index]
	var lastFoundNum int64
	var lastFound bool
	for i := 0; i < len(str); i++ {
		c := int(str[i] - '0')
		if cur.son == nil {
			return s.searchRightKey1(lastFoundNum, lastFound, index, cur)
		}
		if cur.son[c] == nil {
			return s.searchRightKey2(lastFoundNum, lastFound, index, cur, c)
		}
		if cur.son[c].key == key {
			return cur.son[c].key, true
		} else if cur.son[c].key > key {
			lastFoundNum = cur.son[c].key
			lastFound = true
			return lastFoundNum, lastFound
		} else {
			cur = cur.son[c]
		}
	}
	return lastFoundNum, lastFound
}

func (s *SortMap) searchRightKey1(lastFoundNum int64, lastFound bool, index int64, cur *Node) (int64, bool) {
	if cur == s.head[index] {
		for j := index + 1; j <= 19; j++ {
			if s.head[j].son != nil {
				return s.peekMinWithNode(s.head[j], j)
			}
		}
		return lastFoundNum, lastFound
	}
	for j := 0; j < len(cur.par.son); j++ {
		if cur.par.son[j] != nil && cur.par.son[j].key > cur.key {
			return s.peekMinWithNode(cur.par.son[j], int64(j))
		}
	}
	return s.searchRightKey1(lastFoundNum, lastFound, index, cur.par)
}

func (s *SortMap) searchRightKey2(lastFoundNum int64, lastFound bool, index int64, cur *Node, c int) (int64, bool) {
	for k := c + 1; k < 10; k++ {
		if cur.son[k] != nil {
			return s.peekMinWithNode(cur.son[k], int64(k))
		}
	}
	if cur == s.head[index] {
		for j := index + 1; j <= 19; j++ {
			if s.head[j].son != nil {
				return s.peekMinWithNode(s.head[j], j)
			}
		}
		return lastFoundNum, lastFound
	}
	for j := 0; j < len(cur.par.son); j++ {
		if cur.par.son[j] != nil && cur.par.son[j].key > cur.key {
			return s.peekMinWithNode(cur.par.son[j], int64(j))
		}
	}
	return s.searchRightKey1(lastFoundNum, lastFound, index, cur.par)
}

func (s *SortMap) Delete(key int64) {
	str := strconv.FormatInt(key, 10)
	var cur *Node
	if key >= 0 {
		cur = s.head[len(str)+9]
	} else {
		str = str[1:]
		cur = s.head[10-len(str)]
	}
	for i := 0; i < len(str); i++ {
		c := int(str[i] - '0')
		if cur.son == nil {
			return
		}
		if cur.son[c] == nil {
			return
		}
		if cur.son[c].key == key {
			b := key < 0
			s.up(cur, cur.son[c], c, b)
			return
		} else if cur.son[c].key > key {
			return
		}
		cur = cur.son[c]
	}
}

func (s *SortMap) PollMin() {
	s.pollMin(0, 19)
}

func (s *SortMap) pollMin(left, right int64) {
	for i := left; i <= right; i++ {
		if p := s.head[i]; p != nil {
			j := (1 - i/10) * (int64(len(p.son)) - (1 - i/10))
			for j >= 0 && j < int64(len(p.son)) {
				if p.son[j] != nil {
					s.up(p, p.son[j], int(j), i <= 9)
					return
				}
				j += 2*(i/10) - 1
			}
		}
	}
}

func (s *SortMap) PeekMin() (int64, bool) {
	return s.peekMin(0, 19)
}

func (s *SortMap) peekMin(left, right int64) (int64, bool) {
	for i := left; i <= right; i++ {
		if p := s.head[i]; p != nil {
			j := (1 - i/10) * (int64(len(p.son)) - (1 - i/10))
			for j >= 0 && j < int64(len(p.son)) {
				if p.son[j] != nil {
					return p.son[j].key, true
				}
				j += 2*(i/10) - 1
			}
		}
	}
	return 0, false
}

func (s *SortMap) peekMinWithNode(cur *Node, index int64) (int64, bool) {
	if cur != s.head[index] {
		return cur.key, true
	}
	if cur.son != nil {
		j := (int64(len(cur.son)) - 1) * (1 - divAbs(cur.key)) / 2
		for j >= 0 && j < int64(len(cur.son)) {
			if cur.son[j] != nil {
				return cur.son[j].key, true
			}
			j += divAbs(cur.key)
		}
	}
	return 0, false
}

func (s *SortMap) up(par *Node, cur *Node, c int, f bool) {
	if par == nil {
		return
	}
	if cur == nil {
		cur = newNode(0, 0, nil, 0)
	}
	if cur.son == nil {
		cur = nil
		par.son[c] = nil
		s.deleteSon(par)
		return
	}
	k := (len(cur.son) - 1) * converseBool(f)
	for k >= 0 && k <= len(cur.son)-1 {
		if cur.son[k] != nil {
			cur.key = cur.son[k].key
			cur.value = cur.son[k].value
			s.up(cur, cur.son[k], k, f)
			return
		}
		k -= 2*converseBool(f) - 1
	}
	cur = nil
	par.son[c] = nil
	s.deleteSon(par)
}

func (s *SortMap) deleteSon(par *Node) {
	flag := true
	for i := 0; i < len(par.son); i++ {
		if par.son[i] != nil {
			flag = false
			break
		}
	}
	if flag {
		par.son = nil
	}
}

func (s *SortMap) PollMax() {
	s.pollMax(0, 19)
}

func (s *SortMap) pollMax(left, right int64) {
	for i := right; i >= left; i-- {
		if p := s.head[i]; p != nil {
			j := (i / 10) * (int64(len(p.son)) - i/10)
			for j >= 0 && j < int64(len(p.son)) {
				if p.son[j] != nil {
					s.down(p, p.son[j], int(j), i < 10, true)
					return
				}
				j -= 2*(i/10) - 1
			}
		}
	}
}

func (s *SortMap) PeekMax() (int64, bool) {
	return s.peekMax(0, 19)
}

func (s *SortMap) peekMax(left, right int64) (int64, bool) {
	for i := right; i >= left; i-- {
		if p := s.head[i]; p.son != nil {
			j := (i / 10) * (int64(len(p.son)) - i/10)
			for j >= 0 && j < int64(len(p.son)) {
				if p.son[j] != nil {
					return s.down(p, p.son[j], int(j), i < 10, true), true
				}
				j -= 2*(i/10) - 1
			}
		}
	}
	return 0, false
}

func (s *SortMap) peekMaxWithNode(cur *Node, index int64) (int64, bool) {
	if cur.son != nil {
		j := (int64(len(cur.son)) - 1) * (1 + divAbs(cur.key)) / 2
		for j >= 0 && j < int64(len(cur.son)) {
			if cur.son[j] != nil {
				return s.down(cur, cur.son[j], int(j), cur.key < 0, false), true
			}
			j -= divAbs(cur.key)
		}
	}
	if cur != s.head[index] {
		return cur.key, true
	}
	return 0, false
}

func (s *SortMap) down(p *Node, cur *Node, c int, flag bool, needDelete bool) int64 {
	if cur.son != nil {
		i := (len(cur.son) - 1) * (1 - converseBool(flag))
		for i >= 0 && i < len(cur.son) {
			if cur.son[i] != nil {
				return s.down(cur, cur.son[i], i, flag, needDelete)
			}
			i += 2*converseBool(flag) - 1
		}
	}
	ans := cur.key
	if needDelete {
		cur = nil
		p.son[c] = nil
	}
	return ans
}

func (s *SortMap) GetRangeKey(begin, end int64) []int64 {
	beginNum, ok1 := s.SearchRightKey(begin)
	endNum, ok2 := s.SearchLeftKey(end)
	if !(ok1 && ok2) || endNum < beginNum {
		return nil
	}
	var sNode, eNode *Node
	var ok bool
	sNode, ok = s.search(beginNum)
	eNode, ok = s.search(endNum)
	if !ok {
		return nil
	}
	dfsCon := &dfsCon{}
	s.dfs(sNode, sNode, eNode, (int64(len(sNode.son))-1)*(1-divAbs(sNode.key))/2, dfsCon)
	return dfsCon.data
}

func (s *SortMap) dfs(cur, startNode, endNode *Node, index int64, con *dfsCon) {
	if con == nil || (len(con.data) > 0 && con.data[len(con.data)-1] == endNode.key) {
		return
	}
	// 到了最上层节点
	if cur == nil {
		if index < 19 {
			s.head[index+1].key = startNode.key
			s.dfs(s.head[index+1], s.head[index+1], endNode, int64(len(s.head)-1)*(1-(index+1)/10), con)
			s.head[index+1].key = 0
		}
		return
	}
	// 先把自己和子孙节点遍历
	if cur != s.head[cur.index] && cur.key >= startNode.key && cur.key <= endNode.key {
		con.data = append(con.data, cur.key)
	}
	if cur.son != nil {
		i := index
		for i >= 0 && i < int64(len(cur.son)) {
			if cur.son[i] != nil {
				s.dfs(cur.son[i], startNode, endNode, (int64(len(cur.son[i].son))-1)*(1-divAbs(cur.son[i].key))/2, con)
			}
			if cur.par == nil {
				i += 2*(int64(cur.index)/10) - 1
			} else {
				i += divAbs(cur.key)
			}
		}
	}
	// 再遍历往上的邻居节点
	if s.isAncestor(cur, startNode) {
		if cur.par == nil {
			s.dfs(cur.par, startNode, endNode, int64(cur.index), con)
		} else {
			s.dfs(cur.par, startNode, endNode, int64(cur.index)+divAbs(cur.key), con)
		}
	}
}

func (s *SortMap) isAncestor(cur *Node, target *Node) bool {
	for target != nil {
		if target == cur {
			return true
		}
		target = target.par
	}
	return false
}

func converseBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func divAbs(a int64) int64 {
	if a == 0 {
		return 1
	}
	return a / abs(a)
}

func getNextMin(key int64, idx int) int64 {
	if key == 0 {
		return -1
	}
	if key < 0 {
		return -getNextMax(-key, idx)
	}
	bStr := []byte(strconv.FormatInt(key, 10))
	flag := idx
	for i := idx; i >= 0; i-- {
		if bStr[i] != byte('0') {
			bStr[i] = bStr[i] - 1
			flag = i
			break
		}
	}
	for i := flag + 1; i < len(bStr); i++ {
		bStr[i] = byte('9')
	}
	res, _ := strconv.ParseInt(string(bStr), 10, 64)
	return res
}

func getNextMax(key int64, idx int) int64 {
	if key == 0 {
		return 1
	}
	if key < 0 {
		return -getNextMin(-key, idx)
	}
	bStr := []byte("0" + strconv.FormatInt(key, 10))
	idx++
	flag := idx
	for i := idx; i >= 0; i-- {
		if bStr[i] != byte('9') {
			bStr[i] = bStr[i] + 1
			flag = i
			break
		}
	}
	for i := flag + 1; i < len(bStr); i++ {
		bStr[i] = byte('0')
	}
	res, _ := strconv.ParseInt(string(bStr), 10, 64)
	return res
}
