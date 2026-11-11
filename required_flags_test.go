package cobra

import "testing"

func TestValidateRequiredFlagsEmptyAnnotation(t *testing.T) {
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")
	assertNoErr(t, c.Flags().SetAnnotation("name", BashCompOneRequiredFlag, []string{}))

	assertNoErr(t, c.ValidateRequiredFlags())
}

func TestMissingRequiredFlags(t *testing.T) {
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")
	c.Flags().String("other", "", "")
	assertNoErr(t, c.MarkFlagRequired("name"))

	missing := c.MissingRequiredFlags()
	if len(missing) != 1 || missing[0] != "name" {
		t.Fatalf("Expected [name], got %v", missing)
	}

	assertNoErr(t, c.Flags().Set("name", "x"))
	if missing := c.MissingRequiredFlags(); len(missing) != 0 {
		t.Errorf("Expected no missing flag, got %v", missing)
	}
}

func TestValidateRequiredFlagsChain(t *testing.T) {
	root := &Command{Use: "root", Run: emptyRun}
	root.Flags().String("token", "", "")
	assertNoErr(t, root.MarkFlagRequired("token"))
	child := &Command{Use: "child", Run: emptyRun}
	root.AddCommand(child)

	if err := child.ValidateRequiredFlagsChain(); err == nil {
		t.Error("Expected an error for the missing required flag of the parent")
	}
}
