package main

import (
	"database/sql"
)

var DB *sql.DB

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

func AllWeapons() ([]Weapon, error) {
	rows, err := DB.Query("SELECT * FROM weapons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weapons []Weapon

	for rows.Next() {
		var weapon Weapon

		err := rows.Scan(&weapon.Type, &weapon.Name, &weapon.Damage, &weapon.Speed, &weapon.Reach)
		if err != nil {
			return nil, err
		}

		weapons = append(weapons, weapon)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return weapons, nil
}

var Weapons = []Weapon{
	{Type: 0, Name: "Sword", Damage: 10, Speed: 10, Reach: 10},
	{Type: 0, Name: "Longsword", Damage: 15, Speed: 5, Reach: 20},
	{Type: 2, Name: "Greataxe", Damage: 20, Speed: 15, Reach: 10},
}
