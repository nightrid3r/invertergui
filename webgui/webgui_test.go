/*
Copyright (c) 2015, Hendrik van Wyk
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

* Neither the name of invertergui nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package webgui

import (
	"github.com/hpdvanwyk/invertergui/datasource"
	"reflect"
	"testing"
)

type mockSource struct {
}

func NewMockSource() datasource.DataSource {
	return &mockSource{}
}

func (s *mockSource) GetData(status *datasource.MultiplusStatus) error {
	status.OutCurrent = 2.0
	status.InCurrent = 2.3
	status.OutVoltage = 230.0
	status.InVoltage = 230.1
	status.BatVoltage = 25
	status.BatCurrent = -10
	status.InFreq = 50
	status.OutFreq = 50
	status.Leds = []int{0, 0, 0, 0, 1, 0, 0, 1}
	return nil
}

func TestWebGui(t *testing.T) {
	t.Skip("Not yet implimented")
	//TODO figure out how to test template output.
}

type templateTest struct {
	input  *statusProcessed
	output *TemplateInput
}

var templateInputTests = []templateTest{
	{
		input: &statusProcessed{
			status: datasource.MultiplusStatus{
				OutCurrent: 2.0,
				InCurrent:  2.3,
				OutVoltage: 230.0,
				InVoltage:  230.1,
				BatVoltage: 25,
				BatCurrent: -10,
				InFreq:     50,
				OutFreq:    50,
				Leds:       []int{0, 0, 0, 0, 1, 0, 0, 1}},
			err: nil,
		},
		output: &TemplateInput{
			Error:      nil,
			OutCurrent: "2.000",
			OutVoltage: "230.000",
			OutPower:   "460.000",
			InCurrent:  "2.300",
			InVoltage:  "230.100",
			InPower:    "529.230",
			InMinOut:   "69.230",
			BatVoltage: "25.000",
			BatCurrent: "-10.000",
			BatPower:   "-250.000",
			InFreq:     "50.000",
			BatCharge:  "0.000",
			Leds:       []string{"Mains", "Float"}},
	},
}

func TestTemplateInput(t *testing.T) {
	for i := range templateInputTests {
		templateInput := buildTemplateInput(templateInputTests[i].input)
		if !reflect.DeepEqual(templateInput, templateInputTests[i].output) {
			t.Errorf("buildTemplateInput not producing expected results")
		}
	}
}
