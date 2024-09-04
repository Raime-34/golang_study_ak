package main

import "testing"

func Test_4TV(t *testing.T) {
	samsung := Samsunger{
		TV{
			model: "Samsung",
		},
	}

	lg := LGer{
		TV{
			model: "LG",
		},
	}

	if samsung.GetStatus() || lg.GetStatus() {
		t.Errorf("GetStatus() = Wrong state")
	}

	samsung.switchON()
	lg.switchON()

	if !(samsung.GetStatus() && lg.GetStatus()) {
		t.Errorf("TV not turning ON")
	}

	samsung.switchOFF()
	lg.switchOFF()

	if samsung.GetStatus() || lg.GetStatus() {
		t.Errorf("TV not turning OFF")
	}

	if samsung.GetModel() != "Samsung" || lg.GetModel() != "LG" {
		t.Errorf("GetModel() = Wrong model name")
	}

	if samsung.SamsungHub() != "SamsungHub is opened" || lg.LGHub() != "LGHub is opened" {
		t.Errorf("SamsungHub() and LGHub() = return wrong string")
	}
}
