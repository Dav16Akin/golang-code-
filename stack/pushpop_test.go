package stack

import "testing"

func TestPushPop(t *testing.T) {
	c:= new(Stack)
	c.Push(5)

	if i,_ := c.Pop(); i != 5 {
		t.Log("Pop doesn't give 5")
		t.Fail()
	}
}
