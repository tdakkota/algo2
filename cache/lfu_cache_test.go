package cache

import (
	"fmt"
	"testing"

	"github.com/tdakkota/algo2/testutil"
)

type testLFUCache struct {
	cache *LFUCache[int, int]
}

func (c testLFUCache) Put(k, v int) {
	c.cache.Put(k, v)
}

func (c testLFUCache) Get(k int) int {
	v, ok := c.cache.Get(k)
	if !ok {
		return -1
	}
	return v
}

func (c testLFUCache) Delete(k int) (v int) {
	if c.cache.Delete(k) {
		v = 1
	}
	return v
}

func (c testLFUCache) Len() int { return c.cache.Len() }
func (c testLFUCache) Cap() int { return c.cache.Cap() }

func LFUConstructor(capacity int) testLFUCache {
	return testLFUCache{NewLFUCache[int, int](capacity)}
}

// LeetCode-based tests
// First line is array contains function name which should be called
// Second line is argument list (null means there is no args)
// Third line is result list (null means there is no result)
var lfuTests = []string{
	`["LFUCache","put","put","get","put","get","get","put","get","get","get"]
	[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]
	[null,null,null,1,null,-1,3,null,-1,3,4]`,
	`["LFUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	[null,null,null,null,null,null,-1,null,19,17,null,-1,null,null,null,-1,null,-1,5,-1,12,null,null,3,5,5,null,null,1,null,-1,null,30,5,30,null,null,null,-1,null,-1,24,null,null,18,null,null,null,null,14,null,null,18,null,null,11,null,null,null,null,null,18,null,null,-1,null,4,29,30,null,12,11,null,null,null,null,29,null,null,null,null,17,-1,18,null,null,null,-1,null,null,null,20,null,null,null,29,18,18,null,null,null,null,20,null,null,null,null,null,null,null]`,
	// Custom test(not from LeetCode)
	`["LFUCache","put","put","get","delete","get","delete"]
	[[2],[1,1],[2,2],[1],[1],[1],[10]]
	[null,null,null,1,1,-1,0]`,
}

func runLFUCacheTest(t *testing.T, test string) {
	var cache testLFUCache
	for _, c := range parseTest(test) {
		//fmt.Printf("%d %s(%v) = %d\n", i, c.t, c.args, c.result)
		switch c.t {
		case "LFUCache":
			cache = LFUConstructor(c.args[0])
			testutil.Equal(t, c.args[0], cache.Cap())
			testutil.Equal(t, 0, cache.Len())
		case "put":
			cache.Put(c.args[0], c.args[1])
		case "get":
			v := cache.Get(c.args[0])
			if v != c.result {
				t.Fatalf("failed: expected =%d, got =%d", c.result, v)
			}
		case "delete":
			v := cache.Delete(c.args[0])
			if v != c.result {
				t.Fatalf("failed: expected =%d, got =%d", c.result, v)
			}
		default:
			t.Fatal("invalid test")
		}
	}
}

func TestLFUCache(t *testing.T) {
	for i, test := range lfuTests {
		t.Run(fmt.Sprintf("Test #%d", i), func (t *testing.T) {
			runLFUCacheTest(t, test)
		})
	}
	
}