package main

type WeaponType int

const (
	Sword WeaponType = iota + 1
	Mace
	Axe
)

type Weapon struct {
	Type   WeaponType
	Name   string
	Damage int
	Speed  float32
	Reach  float32
}

var Weapons = []Weapon{
	{Type: 0, Name: "Sword", Damage: 10, Speed: 10, Reach: 10},
	{Type: 0, Name: "Longsword", Damage: 15, Speed: 5, Reach: 20},
	{Type: 2, Name: "Greataxe", Damage: 20, Speed: 15, Reach: 10},
}
