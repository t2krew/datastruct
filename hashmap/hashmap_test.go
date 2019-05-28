package hashmap

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		if reflect.Indirect(reflect.ValueOf(New())).Type().Name() != "HashMap" {
			t.Errorf("new instance must be type HashMap")
		}
	})
}

func TestHashMap_Set(t *testing.T) {
	hm := New()
	cases := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "test1",
			key:   "key1",
			value: "value1",
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			hm.Set(item.key, item.value)
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	hm := New()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test1",
			input: "key1",
			want:  "value1",
		},
		{
			name:  "test2",
			input: "key2",
			want:  "value2",
		},
		{
			name:  "test3",
			input: "key3",
			want:  "value3",
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			if value, ok := hm.Get(item.input); ok {
				if value != item.want {
					t.Errorf("want %s, but got %s", item.want, value)
				}
			}
		})
	}
}

func TestHashMap_Len(t *testing.T) {
	hm := New()
	hm.Set("key1", "value1")
	hm.Set("key2", "value2")
	hm.Set("key3", "value3")

	t.Run("get hashmap length", func(t *testing.T) {
		if hm.Len() != 3 {
			t.Errorf("want 3, but got %d", hm.Len())
		}
	})
}

func TestHashMap_Remove(t *testing.T) {
	hm := New()
	hm.Set("key1", "value1")
	t.Run("remove key", func(t *testing.T) {
		hm.Remove("key1")
		if _, ok := hm.Get("key1"); ok {
			t.Errorf("want false, but got %v", ok)
		}
	})
}
