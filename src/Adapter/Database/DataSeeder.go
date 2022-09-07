package Database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Permission"
	"github.com/pascalallen/Baetyl/src/Domain/Auth/Role"
	"gorm.io/gorm"
	"os"
	"path"
	"runtime"
)

type DataSeeder struct {
	permissionsMap       map[string]Permission.Permission
	rolesMap             map[string]Role.Role
	DatabaseConnection   *gorm.DB
	PermissionRepository Permission.PermissionRepository
	RoleRepository       Role.RoleRepository
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

	for _, permissionName := range permissionsToRemove {
		permission := dataSeeder.permissionsMap[permissionName]
		if err := dataSeeder.PermissionRepository.Remove(&permission); err != nil {
			return err
		}
	}

	for _, permissionData := range permissionsData.Permissions {
		id := ulid.MustParse(permissionData.Id)

		permission, err := dataSeeder.PermissionRepository.GetById(id)
		if err != nil {
			return err
		}

		if permission == nil {
			permission = Permission.Define(id, permissionData.Name, permissionData.Description)
			if err := dataSeeder.PermissionRepository.Add(permission); err != nil {
				return err
			}
		}

		if permissionData.Name != permission.Name {
			permission.UpdateName(permissionData.Name)
		}

		if permissionData.Description != permission.Description {
			permission.UpdateDescription(permissionData.Description)
		}

		if err := dataSeeder.PermissionRepository.Save(permission); err != nil {
			return err
		}
	}

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

	for _, roleName := range rolesToRemove {
		role := dataSeeder.rolesMap[roleName]
		if err := dataSeeder.RoleRepository.Remove(&role); err != nil {
			return err
		}
	}

	for _, roleData := range rolesData.Roles {
		id := ulid.MustParse(roleData.Id)

		role, err := dataSeeder.RoleRepository.GetById(id)
		if err != nil {
			return err
		}

		if role == nil {
			role = Role.Define(id, roleData.Name)
			if len(roleData.Permissions) > 0 {
				for _, permissionName := range roleData.Permissions {
					permission, err := dataSeeder.PermissionRepository.GetByName(permissionName)
					if err != nil {
						return err
					}

					if permission != nil && !role.HasPermission(permissionName) {
						role.AddPermission(*permission)
					}
				}
			}

			if err := dataSeeder.RoleRepository.Add(role); err != nil {
				return err
			}
		}

		if roleData.Name != role.Name {
			role.UpdateName(roleData.Name)
		}

		var newRolePermissions []Permission.Permission
		for _, permissionName := range roleData.Permissions {
			permission, err := dataSeeder.PermissionRepository.GetByName(permissionName)
			if err != nil {
				return err
			}

			newRolePermissions = append(newRolePermissions, *permission)
		}

		if err := dataSeeder.DatabaseConnection.Model(&role).Association("Permissions").Replace(newRolePermissions); err != nil {
			return err
		}

		if err := dataSeeder.RoleRepository.Save(role); err != nil {
			return err
		}
	}

	if err := dataSeeder.loadRolesMap(); err != nil {
		return err
	}

	return nil
}

func (dataSeeder *DataSeeder) seedUsers() {
	// TODO
}

func (dataSeeder *DataSeeder) loadPermissionsMap() error {
	permissions, err := dataSeeder.PermissionRepository.GetAll()
	if err != nil {
		return err
	}

	m := make(map[string]Permission.Permission)
	for _, p := range *permissions {
		m[p.Name] = p
	}

	return nil
}

func (dataSeeder *DataSeeder) loadRolesMap() error {
	roles, err := dataSeeder.RoleRepository.GetAll()
	if err != nil {
		return err
	}

	m := make(map[string]Role.Role)
	for _, r := range *roles {
		m[r.Name] = r
	}

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
