package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("Post", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.Form = postData
	form = New(r.Form)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	form = New(postData)
	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("Post", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	postData := url.Values{}
	postData.Add("some_field", "some value")
	form = New(postData)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form shows minlength met when data is shorter")
	}
	isError := form.Errors.Get("some_field")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postData = url.Values{}
	postData.Add("another_field", "abc123")
	form = New(postData)
	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("shows minlength not met when it is")
	}
	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error but get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postData = url.Values{}
	postData.Add("email", "a@a.com")
	form = New(postData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}

	postData = url.Values{}
	postData.Add("email", "")
	form = New(postData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for an invalid email")
	}
}
