package drivercore_test

import (
	"testing"

	"github.com/kuttiproject/drivercore"
	"github.com/kuttiproject/drivercore/drivercoretest"
	"github.com/kuttiproject/drivercore/drivercoretest/drivermock"
)

const TESTK8SVERSION = "1.22"

func init() {
	mockdriver1 := drivermock.New("mock1", "Mock driver with NAT Networking", true, true)
	drivercore.RegisterDriver("mock1", mockdriver1)

	mockdriver1.UpdateRemoteImage(TESTK8SVERSION, false)
}

func TestDriverCore(t *testing.T) {
	cnt := drivercore.DriverCount()
	if cnt < 1 {
		t.Fatal("no drivers found")
	}

	rd := drivercore.RegisteredDrivers()
	if len(rd) < 1 {
		t.Fatal("no drivers found")
	}

	drivercore.ForEachDriver(func(d drivercore.Driver) bool {
		if d.Status() != "Ready" {
			t.Fatalf("Driver status was %v instead of Ready", d.Status())
		}
		return true
	})

	ok := drivercore.RegisteredDriver("mock1")
	if !ok {
		t.Fatal("Driver mock1 not registered")
	}

	_, ok = drivercore.GetDriver("mock1")
	if !ok {
		t.Fatal("Driver mock1 not found")
	}

	drivercoretest.TestDriver(t, "mock1", TESTK8SVERSION)
}
