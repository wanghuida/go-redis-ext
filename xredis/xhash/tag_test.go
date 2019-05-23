package xhash

import "testing"

func TestHump2underline(t *testing.T) {

	name := "UserAnswer"
	result := Hump2underline(name)
	if result != "user_answer" {
		t.Errorf("conver err name=%s result=%s", name, result)
	}

	name = "Hump2underline"
	result = Hump2underline(name)
	if result != "hump2underline" {
		t.Errorf("conver err name=%s result=%s", name, result)
	}
}
