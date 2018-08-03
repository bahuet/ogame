package ogame

// LightLaser ...
type lightLaser struct {
	BaseDefense
}

// NewLightLaser ...
func NewLightLaser() *lightLaser {
	d := new(lightLaser)
	d.ID = LightLaserID
	d.Price = Resources{Metal: 1500, Crystal: 500}
	d.StructuralIntegrity = 2000
	d.ShieldPower = 25
	d.WeaponPower = 100
	d.RapidfireFrom = map[ID]int{DestroyerID: 10, BomberID: 20, DeathstarID: 200}
	d.Requirements = map[ID]int{ShipyardID: 2, LaserTechnologyID: 3}
	return d
}