// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Bookmarks", testBookmarks)
	t.Run("Ogps", testOgps)
	t.Run("Stats", testStats)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Bookmarks", testBookmarksDelete)
	t.Run("Ogps", testOgpsDelete)
	t.Run("Stats", testStatsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Bookmarks", testBookmarksQueryDeleteAll)
	t.Run("Ogps", testOgpsQueryDeleteAll)
	t.Run("Stats", testStatsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Bookmarks", testBookmarksSliceDeleteAll)
	t.Run("Ogps", testOgpsSliceDeleteAll)
	t.Run("Stats", testStatsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Bookmarks", testBookmarksExists)
	t.Run("Ogps", testOgpsExists)
	t.Run("Stats", testStatsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Bookmarks", testBookmarksFind)
	t.Run("Ogps", testOgpsFind)
	t.Run("Stats", testStatsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Bookmarks", testBookmarksBind)
	t.Run("Ogps", testOgpsBind)
	t.Run("Stats", testStatsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Bookmarks", testBookmarksOne)
	t.Run("Ogps", testOgpsOne)
	t.Run("Stats", testStatsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Bookmarks", testBookmarksAll)
	t.Run("Ogps", testOgpsAll)
	t.Run("Stats", testStatsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Bookmarks", testBookmarksCount)
	t.Run("Ogps", testOgpsCount)
	t.Run("Stats", testStatsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Bookmarks", testBookmarksHooks)
	t.Run("Ogps", testOgpsHooks)
	t.Run("Stats", testStatsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Bookmarks", testBookmarksInsert)
	t.Run("Bookmarks", testBookmarksInsertWhitelist)
	t.Run("Ogps", testOgpsInsert)
	t.Run("Ogps", testOgpsInsertWhitelist)
	t.Run("Stats", testStatsInsert)
	t.Run("Stats", testStatsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Bookmarks", testBookmarksReload)
	t.Run("Ogps", testOgpsReload)
	t.Run("Stats", testStatsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Bookmarks", testBookmarksReloadAll)
	t.Run("Ogps", testOgpsReloadAll)
	t.Run("Stats", testStatsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Bookmarks", testBookmarksSelect)
	t.Run("Ogps", testOgpsSelect)
	t.Run("Stats", testStatsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Bookmarks", testBookmarksUpdate)
	t.Run("Ogps", testOgpsUpdate)
	t.Run("Stats", testStatsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Bookmarks", testBookmarksSliceUpdateAll)
	t.Run("Ogps", testOgpsSliceUpdateAll)
	t.Run("Stats", testStatsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
