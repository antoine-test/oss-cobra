package cobra

import "testing"

func TestBindEnv(t *testing.T) {
	t.Setenv("COBRA_TEST_NAME", "from-env")
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")

	assertNoErr(t, c.BindEnv("name", "COBRA_TEST_NAME"))

	if got, _ := c.Flags().GetString("name"); got != "from-env" {
		t.Errorf("Expected %q, got %q", "from-env", got)
	}
}

func TestBindEnvKeepsChangedFlag(t *testing.T) {
	t.Setenv("COBRA_TEST_NAME", "from-env")
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")
	assertNoErr(t, c.Flags().Set("name", "from-cli"))

	assertNoErr(t, c.BindEnv("name", "COBRA_TEST_NAME"))

	if got, _ := c.Flags().GetString("name"); got != "from-cli" {
		t.Errorf("Expected %q, got %q", "from-cli", got)
	}
}

func TestBindEnvUnknownFlag(t *testing.T) {
	c := &Command{Use: "c", Run: emptyRun}
	if err := c.BindEnv("missing", "COBRA_TEST_NAME"); err == nil {
		t.Error("Expected an error")
	}
}

func TestBindEnvCommandValue(t *testing.T) {
	t.Setenv("COBRA_TEST_NAME", "cmd:echo resolved")
	c := &Command{Use: "c", Run: emptyRun}
	c.Flags().String("name", "", "")

	assertNoErr(t, c.BindEnv("name", "COBRA_TEST_NAME"))

	if got, _ := c.Flags().GetString("name"); got != "resolved" {
		t.Errorf("Expected %q, got %q", "resolved", got)
	}
}
