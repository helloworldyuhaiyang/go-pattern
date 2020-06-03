package create

import "testing"

func Test_Builder(t *testing.T) {
	assembly := NewCarBuilder().Color(RedColor)

	familyCar := assembly.Wheels(SteelWheels).TopSpeed(50 * MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(SportsWheels).TopSpeed(150 * MPH).Build()
	sportsCar.Drive()
}
