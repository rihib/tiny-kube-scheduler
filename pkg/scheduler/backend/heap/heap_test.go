package heap

import "testing"

func testHeapObjectKeyFunc(obj testHeapObject) string {
	return obj.name
}

type testHeapObject struct {
	name string
	val  interface{} // TODO: interface{} can be replaced by any
}

func mkHeapObj(name string, val interface{}) testHeapObject {
	return testHeapObject{name: name, val: val}
}

func compareInts(val1 testHeapObject, val2 testHeapObject) bool {
	first := val1.val.(int)
	second := val2.val.(int)
	return first < second
}

// TODO: add testcases
func TestHeapBasic(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	const amount = 500
	var i int
	var zero testHeapObject

	// Make sure that queue is empty
	if item, ok := h.Peek(); ok || item != zero {
		t.Errorf("expected nil object but got %v", item)
	}

	for i = amount; i > 0; i-- {
		h.AddOrUpdate(mkHeapObj(string([]rune{'a', rune(i)}), i))
		head, ok := h.Peek()
		// e: expected, a: actual
		if e, a := i, head.val; !ok || a != e {
			t.Errorf("expected %d, got %d", e, a)
		}
	}

	// Make sure that the numbers are popped in ascending order.
	prevNum := 0
	for i := 0; i < amount; i++ { // TODO: for loop can be modernized using range over int
		item, err := h.Pop()
		num := item.val.(int)
		// All the items must be sorted.
		if err != nil || prevNum > num {
			t.Errorf("got %v out of order, last was %v", item, prevNum)
		}
		prevNum = num
	}

	_, err := h.Pop()
	if err == nil {
		t.Errorf("expected Pop() to error on empty heap")
	}
}

func TestHeap_AddOrUpdate_Add(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.AddOrUpdate(mkHeapObj("foo", 10))
	h.AddOrUpdate(mkHeapObj("bar", 1))
	h.AddOrUpdate(mkHeapObj("baz", 11))
	h.AddOrUpdate(mkHeapObj("zab", 30))
	h.AddOrUpdate(mkHeapObj("foo", 13)) // This updates "foo".

	item, err := h.Pop()
	if e, a := 1, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item, err = h.Pop()
	if e, a := 11, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	if err := h.Delete(mkHeapObj("baz", 11)); err == nil { // Nothing is deleted.
		t.Fatalf("nothing should be deleted from the heap")
	}
	h.AddOrUpdate(mkHeapObj("foo", 14)) // foo is updated.
	item, err = h.Pop()
	if e, a := 14, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item, err = h.Pop()
	if e, a := 30, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
}

func TestHeap_Delete(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.AddOrUpdate(mkHeapObj("foo", 10))
	h.AddOrUpdate(mkHeapObj("bar", 1))
	h.AddOrUpdate(mkHeapObj("bal", 31))
	h.AddOrUpdate(mkHeapObj("baz", 11))

	// Delete head. Delete should work with "key" and doesn't care about the value.
	if err := h.Delete(mkHeapObj("bar", 200)); err != nil {
		t.Fatalf("Failed to delete head.")
	}
	item, err := h.Pop()
	if e, a := 10, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	h.AddOrUpdate(mkHeapObj("zab", 30))
	h.AddOrUpdate(mkHeapObj("faz", 30))
	len := h.data.Len()
	// Delete non-existing item.
	if err = h.Delete(mkHeapObj("non-existent", 10)); err == nil || len != h.data.Len() {
		t.Fatalf("Didn't expect any item removal")
	}
	// Delete tail.
	if err = h.Delete(mkHeapObj("bal", 31)); err != nil {
		t.Fatalf("Failed to delete tail.")
	}
	// Delete one of the items with value 30.
	if err = h.Delete(mkHeapObj("zab", 30)); err != nil {
		t.Fatalf("Failed to delete item.")
	}
	item, err = h.Pop()
	if e, a := 11, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	item, err = h.Pop()
	if e, a := 30, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	if h.data.Len() != 0 {
		t.Fatalf("expected an empty heap.")
	}
}

func TestHeap_AddOrUpdate_Update(t *testing.T) {
	h := New(testHeapObjectKeyFunc, compareInts)
	h.AddOrUpdate(mkHeapObj("foo", 10))
	h.AddOrUpdate(mkHeapObj("bar", 1))
	h.AddOrUpdate(mkHeapObj("bal", 31))
	h.AddOrUpdate(mkHeapObj("baz", 11))

	// Update an item to a value that should push it to the head.
	h.AddOrUpdate(mkHeapObj("baz", 0))
	if h.data.queue[0] != "baz" || h.data.items["baz"].index != 0 {
		t.Fatalf("expected baz to be at the head")
	}
	item, err := h.Pop()
	if e, a := 0, item.val; err != nil || a != e {
		t.Fatalf("expected %d, got %d", e, a)
	}
	// Update bar to push it farther back in the queue.
	h.AddOrUpdate(mkHeapObj("bar", 100))
	if h.data.queue[0] != "foo" || h.data.items["foo"].index != 0 {
		t.Fatalf("expected foo to be at the head")
	}
}
