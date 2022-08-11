package Database

import (
	"errors"
	"fmt"
	GormPermissionRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/Permission"
	GormRoleRepository "github.com/pascalallen/Baetyl/src/Adapter/Repository/Auth/Role"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"log"
	"os"
	"path"
	"runtime"
)

type DataSeeder struct {
	permissionsMap map[string]Permission.Permission
	rolesMap       map[string]Role.Role
}

func (dataSeeder *DataSeeder) Seed() error {
	if err := dataSeeder.seedPermissions(); err != nil {
		return err
	}

	if err := dataSeeder.seedRoles(); err != nil {
		return err
	}

	dataSeeder.seedUsers()

	return nil
}

func (dataSeeder *DataSeeder) seedPermissions() error {
	if err := dataSeeder.loadPermissionsMap(); err != nil {
		return err
	}

	_, filename, _, ok := runtime.Caller(2)
	if !ok {
		return errors.New("error getting filename")
	}

	rootDir := path.Dir(filename)
	permissionsFile := fmt.Sprintf("%s/database/seeds/auth.permissions.json", rootDir)

	contents, err := os.ReadFile(permissionsFile)
	if err != nil {
		return fmt.Errorf("error reading permissions file: %s", err.Error())
	}

	// TODO
	log.Printf("PERMISSIONS FILE CONTENTS: %s", contents)

	return nil
}

func (dataSeeder *DataSeeder) seedRoles() error {
	if err := dataSeeder.loadRolesMap(); err != nil {
		return err
	}

	_, filename, _, ok := runtime.Caller(2)
	if !ok {
		return errors.New("error getting filename")
	}

	rootDir := path.Dir(filename)
	rolesFile := fmt.Sprintf("%s/database/seeds/auth.roles.json", rootDir)

	contents, err := os.ReadFile(rolesFile)
	if err != nil {
		return fmt.Errorf("error reading roles file: %s", err.Error())
	}

	// TODO
	log.Printf("ROLES FILE CONTENTS: %s", contents)

	return nil
}

func (dataSeeder *DataSeeder) seedUsers() {
	// TODO
}

func (dataSeeder *DataSeeder) loadPermissionsMap() error {
	permissionRepository := GormPermissionRepository.GormPermissionRepository{}
	permissions, err := permissionRepository.GetAll()
	if err != nil {
		return err
	}

	for _, p := range *permissions {
		dataSeeder.permissionsMap[p.Name] = p
	}

	return nil
}

func (dataSeeder *DataSeeder) loadRolesMap() error {
	roleRepository := GormRoleRepository.GormRoleRepository{}
	roles, err := roleRepository.GetAll()
	if err != nil {
		return err
	}

	for _, r := range *roles {
		dataSeeder.rolesMap[r.Name] = r
	}

	return nil
}
