package clock

import (
	"fmt"
	"math"
)

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	addToHours := m / 60
	m = m % 60
	if m < 0 {
		m += 60
		h = h - 1
	}
	h += addToHours
	if h == 24 {
		h = 0
	}
	h %= 24
	if h < 0 {
		h += 24
	}
	return Clock{hour: h, minute: m}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	if c.minute > 59 {
		c.hour += c.minute / 60
		c.minute = c.minute % 60

		if c.hour > 23 {
			c.hour %= 24
		}
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.minute -= m
	if c.minute < 0 {
		if math.Abs(float64(c.minute/60)) > 0 {
			c.hour += (c.minute / 60)
		}
		c.hour--
		if math.Abs(float64(c.minute%60)) > 0 {
			c.minute = c.minute % 60
		}
		c.minute += 60
		c.hour = c.hour % 24
		if int(math.Abs(float64(c.hour))) == 24 {
			c.hour = 0
		}
		if c.hour < 0 {
			c.hour += 24
		}
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
