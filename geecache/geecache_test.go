package geecache

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
)

// 数据库
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")

	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Fatalf("callback failed")
	}
}

func TestGet(t *testing.T) {
	// 记录某个键调用回调函数的次数,如果次数大于1，表示多次调用回调函数
	loadCounts := make(map[string]int, len(db))
	gee := NewGroup("score", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key]++
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	for k, v := range db {
		if view, err := gee.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of Tom")
		}
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 || loadCounts[k] == 0 {
			t.Fatalf("cache %s miss", k)
		}
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknown should be empty, but %s got", view)
	}
}

func Test(t *testing.T) {

	ss := strings.SplitN("/a/b/c", "/", 2)
	if ss[0] != "" {
		t.Fatalf("error ")
	}

}

func TestGetGroup(t *testing.T) {
	groupName := "score"
	NewGroup(groupName, 2<<10, GetterFunc(func(key string) (bytes []byte, err error) {
		return
	}))
	if group := GetGroup(groupName); group == nil || group.name != groupName {
		t.Fatalf("group %s not exist", groupName)
	}

	if group := GetGroup(groupName + "111"); group != nil {
		t.Fatalf("expect nil,but %s not", group.name)
	}
}
