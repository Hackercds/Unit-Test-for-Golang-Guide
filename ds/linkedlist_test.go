package ds_test

import (
	"DeepTest/ds"
	"reflect"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := ds.NewLinkedList()
	if !ll.IsEmpty() {
		t.Error("new list should be empty")
	}
	if ll.Size() != 0 {
		t.Errorf("size = %d, want 0", ll.Size())
	}
	if len(ll.ToSlice()) != 0 {
		t.Error("ToSlice() should return empty slice for new list")
	}
}

func TestInsertAtHead(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtHead(1)
	ll.InsertAtHead(2)
	ll.InsertAtHead(3)

	expected := []int{3, 2, 1}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ToSlice() = %v, want %v", got, expected)
	}
	if ll.Size() != 3 {
		t.Errorf("size = %d, want 3", ll.Size())
	}
}

func TestInsertAtTail(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	expected := []int{1, 2, 3}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ToSlice() = %v, want %v", got, expected)
	}
}

func TestDeleteHead(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	if !ll.Delete(1) {
		t.Error("Delete(1) should return true")
	}
	expected := []int{2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("ToSlice() = %v, want %v", ll.ToSlice(), expected)
	}
}

func TestDeleteMiddle(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	if !ll.Delete(2) {
		t.Error("Delete(2) should return true")
	}
	expected := []int{1, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("ToSlice() = %v, want %v", ll.ToSlice(), expected)
	}
}

func TestDeleteTail(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	if !ll.Delete(3) {
		t.Error("Delete(3) should return true")
	}
	expected := []int{1, 2}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("ToSlice() = %v, want %v", ll.ToSlice(), expected)
	}
}

func TestDeleteNotFound(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)

	if ll.Delete(99) {
		t.Error("Delete(99) should return false")
	}
	if ll.Size() != 2 {
		t.Errorf("size should not change, got %d", ll.Size())
	}
}

func TestDeleteEmptyList(t *testing.T) {
	ll := ds.NewLinkedList()
	if ll.Delete(1) {
		t.Error("Delete on empty list should return false")
	}
}

func TestSearch(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtTail(1)
	ll.InsertAtTail(2)
	ll.InsertAtTail(3)

	if !ll.Search(2) {
		t.Error("Search(2) should return true")
	}
	if ll.Search(99) {
		t.Error("Search(99) should return false")
	}
}

func TestSearchEmptyList(t *testing.T) {
	ll := ds.NewLinkedList()
	if ll.Search(1) {
		t.Error("Search on empty list should return false")
	}
}

func TestLinkedListSize(t *testing.T) {
	ll := ds.NewLinkedList()
	ll.InsertAtHead(1)
	ll.InsertAtTail(2)
	ll.InsertAtHead(3)
	if ll.Size() != 3 {
		t.Errorf("size = %d, want 3", ll.Size())
	}
	ll.Delete(1)
	if ll.Size() != 2 {
		t.Errorf("size = %d, want 2", ll.Size())
	}
}
