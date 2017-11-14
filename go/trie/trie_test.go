package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewNode()

	trie.Load("adc", "adc")
	trie.Load("abc", "abc")
	trie.Load("add", "add")
	trie.Load("adsfa", "adsfa")

	str := "adcxxxxxxxxx"

	val, ok := trie.Find(str)
	t.Errorf("%v %v %v ", str, val, ok)
	str = "ab"
	val, ok = trie.Find(str)
	t.Errorf("%v %v %v", str, val, ok)

	str = "adsfaaaaadfadfadfa"
	val, ok = trie.Find(str)
	t.Errorf("%v %v %v", str, val, ok)
}
