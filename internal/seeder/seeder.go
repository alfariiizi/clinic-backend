package seeder

import (
	"context"
	"fmt"

	"github.com/alfariiizi/vandor/config"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
	"github.com/alfariiizi/vandor/internal/infrastructure/db/user"
	"github.com/alfariiizi/vandor/internal/utils"
)

func GenerateAdmin(db *db.Client) error {
	cfg := config.GetConfig()

	tx, err := db.Tx(context.Background())
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	clinicCreated, err := tx.Clinic.Create().
		SetName("Default Clinic").
		SetAddress("Default Address").
		SetPhone("1234567890").
		SetType("general").
		SetEnabledFeatures([]string{"chat", "appointments", "billing"}).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	password, err := utils.HashPassword(cfg.Superadmin.Password)
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	_, err = tx.User.Create().
		SetEmail(cfg.Superadmin.Email).
		SetName("Admin").
		SetRole(user.RoleADMIN).
		SetPasswordHash(*password).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	userPassword, err := utils.HashPassword("123123123")
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	_, err = tx.User.Create().
		SetEmail("alfariiiziiiii@gmail.com").
		SetName("Dr. Alfarizi").
		SetRole(user.RoleDOCTOR).
		SetPasswordHash(*userPassword).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	_, err = tx.Doctor.Create().
		SetName("Dr. Alfariiizi").
		SetEmail("alfariiiziiiii@gmail.com").
		SetSpecialization("General Practitioner").
		SetPhone("1234567890").
		SetClinicID(clinicCreated.ID).
		SetAvailability(map[string]any{"monday": []string{"09:00-12:00", "13:00-17:00"}, "tuesday": []string{"09:00-12:00", "13:00-17:00"}}).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Println("Seeder Super Admin created successfully")

	return nil
}

func GenerateDefaultTenant(db *db.Client) error {
	tx, err := db.Tx(context.Background())
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}
	}()

	_, err = tx.Clinic.Create().
		SetName("Default Clinic").
		SetAddress("Default Address").
		SetPhone("1234567890").
		SetEnabledFeatures([]string{"chat", "appointments", "billing"}).
		Save(context.Background())
	if err != nil {
		_ = tx.Rollback()
		panic(err)
	}

	return nil
}
