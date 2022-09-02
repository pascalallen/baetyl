package Database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"os"
	"path"
	"runtime"
)

type DataSeeder struct {
	permissionsMap map[string]Permission.Permission
	rolesMap       map[string]Role.Role
}

type PermissionData struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionsData struct {
	Permissions []PermissionData `json:"permissions"`
}

type RoleData struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type RolesData struct {
	Roles []RoleData `json:"roles"`
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

	var permissionsData PermissionsData
	if err := json.Unmarshal(contents, &permissionsData); err != nil {
		return fmt.Errorf("error parsing permissions json: %s", err.Error())
	}

	var currentPermissions []string
	for permissionName := range dataSeeder.permissionsMap {
		currentPermissions = append(currentPermissions, permissionName)
	}

	var seedPermissions []string
	for _, permissionData := range permissionsData.Permissions {
		seedPermissions = append(seedPermissions, permissionData.Name)
	}

	var permissionsToRemove []string
	for _, permissionName := range seedPermissions {
		if len(currentPermissions) > 0 && !contains(currentPermissions, permissionName) {
			permissionsToRemove = append(permissionsToRemove, permissionName)
		}
	}

	// TODO: Remove/add permissions when permission repository is implemented

	if err := dataSeeder.loadPermissionsMap(); err != nil {
		return err
	}

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

	var rolesData RolesData
	if err := json.Unmarshal(contents, &rolesData); err != nil {
		return fmt.Errorf("error parsing roles json: %s", err.Error())
	}

	var currentRoles []string
	for roleName := range dataSeeder.rolesMap {
		currentRoles = append(currentRoles, roleName)
	}

	var seedRoles []string
	for _, roleData := range rolesData.Roles {
		seedRoles = append(seedRoles, roleData.Name)
	}

	var rolesToRemove []string
	for _, roleName := range seedRoles {
		if len(currentRoles) > 0 && !contains(currentRoles, roleName) {
			rolesToRemove = append(rolesToRemove, roleName)
		}
	}

	// TODO: Remove/add roles when role repository is implemented

	if err := dataSeeder.loadRolesMap(); err != nil {
		return err
	}

	return nil
}

func (dataSeeder *DataSeeder) seedUsers() {
	// TODO
}

func (dataSeeder *DataSeeder) loadPermissionsMap() error {
	// TODO: Fetch all permissions and make a hash map where permission name is the key, when permission repository is implemented

	return nil
}

func (dataSeeder *DataSeeder) loadRolesMap() error {
	// TODO: Fetch all roles and make a hash map where role name is the key, when role repository is implemented

	return nil
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
