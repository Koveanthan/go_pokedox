package poke_cache

import (
	"testing"
	"time"
    "fmt"

)

func TestGetAddLocations(t *testing.T) {
    const interval = 5 * time.Second
    cases := []struct {
        url string
        val []byte
    }{
        {
            "example.com/1",
            []byte("This is a sample text"),
        },
        {
            "example.com/2",
            []byte("This is a sample text"),
        },
        {
            "example.com/3",
            []byte("This is a sample text"),
        },
    }
    
    cache := NewCache(interval)

    for i, c := range cases {
        t.Run(fmt.Sprintf("Running case %v", i+1), func (t *testing.T) {
            cache.Add(c.url, c.val) 
            retval, ok := cache.Get(c.url)
            if !ok {
                t.Errorf("value not present")
            }
            if string(retval) != string(c.val) {
                t.Errorf("value is different")
            }
        })         
    }
}

func TestReapLoop(t *testing.T) {
    const interval = 5 * time.Second
    const waitTime = interval + (5 * time.Millisecond)
    cache := NewCache(interval)

    cache.Add("/example/test", []byte("This is a test"))

    if _, ok := cache.Get("/example/test"); !ok {
        t.Errorf("expected to find val")
    } 

    time.Sleep(waitTime)
    
    if _, ok := cache.Get("/example/test"); ok {
        t.Errorf("expected not to find val")
    } 
}
