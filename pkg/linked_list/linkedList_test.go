package linked_list

import (
	"bytes"
	"testing"
)

func Test_CreateList(t *testing.T) {
	lst := CreateList()
	if lst == nil {
		t.Errorf("Список не создан")
	}
}

func TestAdd(t *testing.T) {
	lst := CreateList()
	lst.Add(1)
	lst.Add(2)
	lst.Add(3)
	if lst.len != 3 {
		t.Errorf("В список не добавляются элементы")
	}
}

func TestPrint(t *testing.T) {
	lst := CreateList()
	lst.Add(1)
	lst.Add(2)
	lst.Add(3)
	lst.Add(4)
	lst.Add(5)
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	lst.Print(buf)
	expected := "1 2 3 4 5 "
	if expected != buf.String() {
		t.Errorf("expected %s,\n got %s", expected, buf.String())
	}
}

func TestDelete(t *testing.T) {
	lst := CreateList()
	lst.Add(1)
	lst.Add(2)
	lst.Add(3)
	lst.Add(4)
	lst.Add(5)
	expected1 := "1 2 3 5 "
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	lst.Delete(4)
	lst.Add(4)
	lst.Delete(4)
	lst.Print(buf)
	if buf.String() != expected1 {
		t.Errorf("expected %s,\n got %s", expected1, buf.String())
	}
	buf.Reset()

	expected2 := "2 3 5 "
	lst.Delete(1)
	lst.Print(buf)
	if buf.String() != expected2 {
		t.Errorf("expected %s,\n got %s", expected2, buf.String())
	}
	buf.Reset()

	expected3 := "2 3 "
	lst.Delete(5)
	lst.Print(buf)
	if buf.String() != expected3 {
		t.Errorf("expected %s,\n got %s", expected3, buf.String())
	}
	buf.Reset()

	expected4 := ""
	lst.Delete(2)
	lst.Delete(3)
	lst.Print(buf)
	if buf.String() != expected4 {
		t.Errorf("expected %s,\n got %s", expected4, buf.String())
	}
	buf.Reset()
}

func BenchmarkLinkedList_Add(b *testing.B) {
	lst := CreateList()
	for i := 0; i < b.N; i++ {
		lst.Add(i)
	}
}
