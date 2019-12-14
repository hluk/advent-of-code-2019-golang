package adv

import "testing"

func TestMin(t *testing.T) {
	AssertEq(Min(0, 0), 0, t)
	AssertEq(Min(1, 1), 1, t)
	AssertEq(Min(0, 1), 0, t)
	AssertEq(Min(1, 0), 0, t)
	AssertEq(Min(0, -1), -1, t)
	AssertEq(Min(-1, 0), -1, t)
}

func TestMax(t *testing.T) {
	AssertEq(Max(0, 0), 0, t)
	AssertEq(Max(1, 1), 1, t)
	AssertEq(Max(0, 1), 1, t)
	AssertEq(Max(1, 0), 1, t)
	AssertEq(Max(0, -1), 0, t)
	AssertEq(Max(-1, 0), 0, t)
}

func TestAbs(t *testing.T) {
	AssertEq(Abs(0), 0, t)
	AssertEq(Abs(1), 1, t)
	AssertEq(Abs(-1), 1, t)
}

func TestSign(t *testing.T) {
	AssertEq(Sign(0), 0, t)
	AssertEq(Sign(1), 1, t)
	AssertEq(Sign(-1), -1, t)
	AssertEq(Sign(0), 0, t)
	AssertEq(Sign(99), 1, t)
	AssertEq(Sign(-99), -1, t)
}

func TestGDC(t *testing.T) {
	AssertEq(GCD(0, 0), 0, t)
	AssertEq(GCD(1, 1), 1, t)
	AssertEq(GCD(99, 99), 99, t)
	AssertEq(GCD(1*2*3*4*5*6*7*8, 5*6*7*8), 5*6*7*8, t)
	AssertEq(GCD(13, 7), 1, t)
}

func TestLCM(t *testing.T) {
	AssertEq(LCM(0, 0), 0, t)
	AssertEq(LCM(3, 12), 12, t)
	AssertEq(LCM(3, 13), 3*13, t)
	AssertEq(LCM(186028, LCM(135024, 193052)), 303070460651184, t)
}
