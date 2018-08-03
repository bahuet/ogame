package ogame

// SmallShieldDome ...
type smallShieldDome struct {
	BaseDefense
}

// NewSmallShieldDome ...
func NewSmallShieldDome() *smallShieldDome {
	d := new(smallShieldDome)
	d.ID = SmallShieldDomeID
	d.Price = Resources{Metal: 10000, Crystal: 10000}
	d.StructuralIntegrity = 20000
	d.ShieldPower = 2000
	d.WeaponPower = 1
	d.Requirements = map[ID]int{ShipyardID: 1, ShieldingTechnologyID: 2}
	return d
}