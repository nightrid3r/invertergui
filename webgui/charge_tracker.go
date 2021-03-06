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
	"time"
)

type ChargeTracker struct {
	fullLevel    float64
	currentLevel float64
	lastUpdate   time.Time
}

// Creates a new charge tracker. fullLevel is in A h.
func NewChargeTracker(fullLevel float64) *ChargeTracker {
	return &ChargeTracker{
		fullLevel:    fullLevel,
		currentLevel: fullLevel, // Have to start somewhere.
		lastUpdate:   time.Now(),
	}
}

func (c *ChargeTracker) Update(amp float64, timestamp time.Time) {
	newNow := timestamp
	elapsed := newNow.Sub(c.lastUpdate).Hours()
	c.lastUpdate = newNow
	c.currentLevel -= elapsed * amp
}

func (c *ChargeTracker) CurrentLevel() float64 {
	return c.currentLevel
}

func (c *ChargeTracker) Reset() {
	c.currentLevel = c.fullLevel
}
