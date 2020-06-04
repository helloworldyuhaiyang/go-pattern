package create

import "testing"

func Test_Builder(t *testing.T) {
	carBuilder := NewCarBuilder()

	familyCar, err := carBuilder.Color(BlackColor).Wheels(SteelWheels).TopSpeed(50 * MPH).Build()
	if err != nil {
		t.Fail()
	}
	_ = familyCar.Drive()

	sportsCar, err := carBuilder.Color(RedColor).Wheels(SportsWheels).TopSpeed(150 * MPH).Build()
	if err != nil {
		t.Fail()
	}
	_ = sportsCar.Drive()
}
