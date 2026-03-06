package token

import (
	"testing"
)

func TestLookupIdent(t *testing.T) {

	token := LookupIdent("resource")
	if token != KEYWORD {
		t.Errorf("LookupIdent did not return KEYWORD")
	}

	token = LookupIdent("file")
	if token != IDENT {
		t.Errorf("LookupIdent did not return IDENT")
	}
}
