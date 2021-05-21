// Package drivercoretest provides utilities for driver testing.
// It should be used only by driver developers, and only in tests.
package drivercoretest

import (
	"testing"

	"github.com/kuttiproject/drivercore"
)

// TestDriver runs a set of tests on any driver.
// It first tests for the presence of the specified driver.
// Then, it calls UpateImageList on the driver.
// Then it checks if an image exists for the specified version.
// If so, and the image is already downloaded, then it is purged.
// Thereafer, Fetch is called on the image to download it.
// Then TestDriver creates a network, and a machine connected to it.
// The machine will be created with the specified k8sversuon.
// It then tries to start the machine, forward and unforward
// ports, and finally stop and delete the machine and network.
func TestDriver(t *testing.T, drivername string, k8sversion string) {
	// Test for presence of driver
	drv, ok := drivercore.GetDriver(drivername)
	if !ok {
		t.Logf("Driver '%v' not registered.", drivername)
		t.FailNow()
	}
	// Test for UpateImageList
	err := drv.UpdateImageList()
	if err != nil {
		t.Logf("Driver '%v' could not update its image list: %v.", drivername, err)
		t.FailNow()
	}

	img, err := drv.GetImage(k8sversion)
	if err != nil {
		t.Logf(
			"Driver '%v' could not get image for k8s version %v: %v.",
			drivername,
			k8sversion,
			err,
		)
		t.FailNow()
	}

	if img.Status() == drivercore.ImageStatusDownloaded {
		t.Logf("Testing image local purge.")
		err = img.PurgeLocal()
		if err != nil {
			t.Logf("Image.PurgeLocal returned an error: %v", err)
			t.Fail()
		}
	}

	err = img.Fetch()
	if err != nil {
		t.Logf("Image.Fetch returned an error: %v", err)
		t.FailNow()
	}

	if img.Status() != drivercore.ImageStatusDownloaded {
		t.Logf(
			"Image.Status shows wrong result after apparently successful fetch: %v.",
			img.Status(),
		)
		t.FailNow()
	}

	t.Log("Attempting to create network for a cluster called 'zintakova'.")
	nw, err := drv.NewNetwork("zintakova")
	if err != nil {
		t.Logf("Error in NewNetwork: %v\n", err)
		t.FailNow()
	}

	nwqname := drv.QualifiedNetworkName("zintakova")
	if nw.Name() != nwqname {
		t.Logf(
			"Wrong name returned. Wanted %v, got %v, which should now be deleted manually.",
			nwqname,
			nw.Name(),
		)
		t.FailNow()
	}

	t.Log("NewNetwork worked as expected. Calling again with same parameters...")
	_, err = drv.NewNetwork("zintakova")
	if err == nil {
		t.Logf(
			"The second call to NewNetwork should have failed. Remember to clean network called %v",
			nwqname,
		)
		t.FailNow()
	}

	t.Logf("NewNetwork second call errored as expected, with %v.", err)
	t.Log("Now calling NewMachine...")
	newnode, err := drv.NewMachine("champu", "zintakova", k8sversion)
	if err != nil {
		t.Logf("Error from NewMachine: %v.", err)
		t.FailNow()
	} else {

		t.Log(newnode)

		t.Logf("NewMachine seems to have worked. Now starting the new host...")
		//newnode.WaitForStateChange(20)
		err = newnode.Start()
		if err != nil {
			t.Logf("Host starting failed with error: %v.", err)
			t.Fail()
		} else {
			if drv.UsesNATNetworking() {
				t.Logf("Host starting worked. Now waiting twenty seconds, and forwarding SSH port...")
				newnode.WaitForStateChange(20)
				err = newnode.ForwardSSHPort(10001)
				if err != nil {
					t.Logf("SSH port forwarding failed with error: %v", err)
					t.Fail()
				}

				sshaddr := newnode.SSHAddress()
				if sshaddr != "localhost:10001" {
					t.Logf("SSH port mapping does not appear to be successful. Value was '%s'", sshaddr)
					t.Fail()
				} else {
					t.Log("SSH Port mapping successful.")
				}

				t.Log("Now trying to map a non-SSH port...")
				err = newnode.ForwardPort(10080, 80)
				if err != nil {
					t.Logf("Port forwarding host 10080 to vm 80 failed with:%v", err)
					t.Fail()
				} else {
					t.Logf("Now trying that again with a different hostport. Should fail...")
					err = newnode.ForwardPort(10081, 80)
					if err == nil {
						t.Log("Second forward should have failed. It didn't.")
						t.Fail()
					}

					t.Logf("Now trying that again with a different port, but already allocated hostport. Should fail...")
					err = newnode.ForwardPort(10080, 81)
					if err == nil {
						t.Log("Third forward should have failed. It didn't.")
						t.Fail()
					}

					t.Log("Port forwarding tests were successful. Now unforwarding...")
					err = newnode.UnforwardPort(80)
					if err != nil {
						t.Logf("Unforwarding vm port 80 failed with:%v", err)
						t.Fail()
					}
				} // ForwardPort
			} // UsesNATNetworking
		} // Start

		t.Log("Now stopping host...")
		err = newnode.Stop()
		if err != nil {
			t.Logf("Error stopping host: %v", err)
			t.Fail()
		}
	} // NewMachine

	t.Logf("NewMachine seems to have created node with name %s and status %s. Now waiting 20 seconds and calling DeleteHost...", newnode.Name(), newnode.Status())
	newnode.WaitForStateChange(25)
	err = drv.DeleteMachine("champu", "zintakova")
	if err != nil {
		t.Logf("Error from DeleteMachine: %v\n", err)
		t.FailNow()
	}

	t.Log("DeleteMachine seems to have worked. Now calling DeleteNetwork...")
	err = drv.DeleteNetwork("zintakova")
	if err != nil {
		t.Logf("Error from DeleteNetwork: %v\n", err)
		t.FailNow()
	}
}
