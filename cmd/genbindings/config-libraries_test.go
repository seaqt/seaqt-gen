package main

import (
	"reflect"
	"testing"
)

func TestSelectModules(t *testing.T) {
	modules := []string{"Core", "Gui", "Network", "RemoteObjects"}

	cases := []struct {
		filter string
		want   []string
	}{
		{"", modules},
		{"RemoteObjects,Network,Core", []string{"Core", "Network", "RemoteObjects"}},
		{" Core , Gui,", []string{"Core", "Gui"}},
	}
	for _, c := range cases {
		if got := selectModules(modules, c.filter); !reflect.DeepEqual(got, c.want) {
			t.Errorf("selectModules(%q) = %v, want %v", c.filter, got, c.want)
		}
	}

	defer func() {
		if recover() == nil {
			t.Errorf("selectModules with an unknown module did not panic")
		}
	}()
	selectModules(modules, "Core,Nope")
}
