package graph

import "graphql/internal/store"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *store.Category
	CourseDB   *store.Course
}
