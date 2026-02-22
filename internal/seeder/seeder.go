// Package seeder initializes required MongoDB collections and default records
// (permissions, roles, admin user) on the first run.
// It is idempotent: records that already exist are skipped.
package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/hoag/go-social-feed/internal/models"
	pkgMongo "github.com/hoag/go-social-feed/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	colUsers       = "social_users"
	colRoles       = "social_roles"
	colPermissions = "social_permissions"

	adminPhone    = "0328923189"
	adminPassword = "123456"
	adminUsername = "Admin"
)

// Run seeds the database. Safe to call on every startup.
func Run(ctx context.Context, db pkgMongo.Database) error {
	fmt.Println("[seeder] Starting database seed...")

	permIDs, err := seedPermissions(ctx, db)
	if err != nil {
		return fmt.Errorf("seeder.seedPermissions: %w", err)
	}

	adminRoleID, err := seedRoles(ctx, db, permIDs)
	if err != nil {
		return fmt.Errorf("seeder.seedRoles: %w", err)
	}

	if err := seedAdminUser(ctx, db, adminRoleID, permIDs); err != nil {
		return fmt.Errorf("seeder.seedAdminUser: %w", err)
	}

	fmt.Println("[seeder] Seed completed successfully.")
	return nil
}

// seedPermissions ensures all defined permissions exist and returns their ObjectIDs.
func seedPermissions(ctx context.Context, db pkgMongo.Database) (map[string]primitive.ObjectID, error) {
	col := db.Collection(colPermissions)
	now := time.Now()

	defined := []models.Permission{
		models.PermissionCreatePost,
		models.PermissionUpdatePost,
		models.PermissionDeletePost,
		models.PermissionReadPost,
	}

	ids := make(map[string]primitive.ObjectID, len(defined))

	for _, perm := range defined {
		// Check if already exists
		var existing models.Permissions
		err := col.FindOne(ctx, bson.M{"name": string(perm)}).Decode(&existing)
		if err == nil {
			// Already exists — use existing ID
			ids[string(perm)] = existing.ID
			fmt.Printf("[seeder] Permission '%s' already exists, skipping.\n", perm)
			continue
		}

		// Insert new
		newPerm := models.Permissions{
			ID:        primitive.NewObjectID(),
			Name:      perm,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if _, err := col.InsertOne(ctx, newPerm); err != nil {
			return nil, fmt.Errorf("insert permission '%s': %w", perm, err)
		}
		ids[string(perm)] = newPerm.ID
		fmt.Printf("[seeder] Created permission '%s'.\n", perm)
	}

	return ids, nil
}

// seedRoles ensures admin and user roles exist. Returns the admin role ObjectID.
func seedRoles(ctx context.Context, db pkgMongo.Database, permIDs map[string]primitive.ObjectID) (primitive.ObjectID, error) {
	col := db.Collection(colRoles)
	now := time.Now()

	// Collect all permission names for admin role
	allPerms := make([]string, 0, len(permIDs))
	for name := range permIDs {
		allPerms = append(allPerms, name)
	}

	type roleDef struct {
		name  models.Role
		perms []string
	}

	roleDefs := []roleDef{
		{name: models.RoleAdmin, perms: allPerms},
		{name: models.RoleUser, perms: []string{
			string(models.PermissionCreatePost),
			string(models.PermissionReadPost),
		}},
	}

	var adminRoleID primitive.ObjectID

	for _, rd := range roleDefs {
		var existing models.Roles
		err := col.FindOne(ctx, bson.M{"name": string(rd.name)}).Decode(&existing)
		if err == nil {
			fmt.Printf("[seeder] Role '%s' already exists, skipping.\n", rd.name)
			if rd.name == models.RoleAdmin {
				adminRoleID = existing.ID
			}
			continue
		}

		newRole := models.Roles{
			ID:          primitive.NewObjectID(),
			Name:        rd.name,
			Permissions: rd.perms,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		if _, err := col.InsertOne(ctx, newRole); err != nil {
			return primitive.NilObjectID, fmt.Errorf("insert role '%s': %w", rd.name, err)
		}

		if rd.name == models.RoleAdmin {
			adminRoleID = newRole.ID
		}
		fmt.Printf("[seeder] Created role '%s'.\n", rd.name)
	}

	return adminRoleID, nil
}

// seedAdminUser ensures the admin user exists in social_users.
func seedAdminUser(ctx context.Context, db pkgMongo.Database, adminRoleID primitive.ObjectID, permIDs map[string]primitive.ObjectID) error {
	col := db.Collection(colUsers)
	now := time.Now()

	// Check if admin already exists by phone
	var existing models.User
	err := col.FindOne(ctx, bson.M{"phone": adminPhone}).Decode(&existing)
	if err == nil {
		fmt.Printf("[seeder] Admin user (phone: %s) already exists, skipping.\n", adminPhone)
		return nil
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	// Collect all permission IDs for admin
	permObjIDs := make([]primitive.ObjectID, 0, len(permIDs))
	for _, id := range permIDs {
		permObjIDs = append(permObjIDs, id)
	}

	admin := models.User{
		ID:          primitive.NewObjectID(),
		Username:    adminUsername,
		Phone:       adminPhone,
		Password:    string(hashed),
		Roles:       []primitive.ObjectID{adminRoleID},
		Permissions: permObjIDs,
		Birthday:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if _, err := col.InsertOne(ctx, admin); err != nil {
		return fmt.Errorf("insert admin user: %w", err)
	}

	fmt.Printf("[seeder] Created admin user (phone: %s, id: %s).\n", adminPhone, admin.ID.Hex())
	return nil
}
