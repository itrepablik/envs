package envs

import "testing"

func TestLoadEnvFromFile(t *testing.T) {
	err := LoadEnvFromFile("server.env", "server1.env")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Loaded environment variables from server.env and server1.env")
}
