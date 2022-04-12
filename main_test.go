package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsWeekend(t *testing.T) {
	day, _ := time.Parse("2006/01/02", "2022/04/13")
	ret := isWeekend(day)
	assert.Equal(t, false, ret)

	day, _ = time.Parse("2006/01/02", "2022/04/16")
	ret = isWeekend(day)
	assert.Equal(t, true, ret)

	day, _ = time.Parse("2006/01/02", "2022/04/17")
	ret = isWeekend(day)
	assert.Equal(t, true, ret)
}

func TestIsHoliday(t *testing.T) {
	day, _ := time.Parse("2006/01/02", "2022/04/13")
	ret := isHoliday(day)
	assert.Equal(t, false, ret)

	day, _ = time.Parse("2006/01/02", "2022/04/16")
	ret = isHoliday(day)
	assert.Equal(t, false, ret)

	day, _ = time.Parse("2006/01/02", "2022/04/29")
	ret = isHoliday(day)
	assert.Equal(t, true, ret)

	day, _ = time.Parse("2006/01/02", "2022/05/03")
	ret = isHoliday(day)
	assert.Equal(t, true, ret)
}
