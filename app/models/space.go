package models

import (
	"mm-wiki/app/utils"
	"github.com/snail007/go-activerecord/mysql"
	"time"
)

const (
	Space_Delete_True = 1
	Space_Delete_False = 0

	Space_Root_Id = 1
	Space_Admin_Id = 2
	Space_Default_Id = 3
)

const Table_Space_Name = "space"

type Space struct {

}

var SpaceModel = Space{}

// get space by space_id
func (s *Space) GetSpaceBySpaceId(spaceId string) (space map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id":   spaceId,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	space = rs.Row()
	return
}

// space_id and name is exists
func (s *Space) HasSameName(spaceId, name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id <>": spaceId,
		"name":   name,
		"is_delete":  Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// name is exists
func (s *Space) HasSpaceName(name string) (has bool, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name":  name,
		"is_delete": Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	if rs.Len() > 0 {
		has = true
	}
	return
}

// get space by name
func (s *Space) GetSpaceByName(name string) (space map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name":  name,
		"is_delete": Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	space = rs.Row()
	return
}

// delete space by space_id
func (s *Space) Delete(spaceId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Update(Table_Space_Name, map[string]interface{}{
		"is_delete": Space_Delete_False,
		"update_time": time.Now().Unix(),
	}, map[string]interface{}{
		"space_id": spaceId,
	}))
	if err != nil {
		return
	}
	return
}

// insert space
func (s *Space) Insert(spaceValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Space_Name, spaceValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// update space by space_id
func (s *Space) Update(spaceId string, spaceValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	spaceValue["update_time"] =  time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Space_Name, spaceValue, map[string]interface{}{
		"space_id":   spaceId,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// get limit spaces by search keyword
func (s *Space) GetSpacesByKeywordAndLimit(keyword string, limit int, number int) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"is_delete":  Space_Delete_False,
	}).WhereWrap(map[string]interface{}{
		"name LIKE": "%" + keyword + "%",
	}, "AND (", "").WhereWrap(map[string]interface{}{
		"description LIKE": "%" + keyword + "%",
	}, "OR", ")").Limit(limit, number).OrderBy("space_id", "DESC")
	rs, err = db.Query(sql)

	if err != nil {
		return
	}
	spaces = rs.Rows()

	return
}

// get limit spaces
func (s *Space) GetSpacesByLimit(limit int, number int) (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			From(Table_Space_Name).
			Where(map[string]interface{}{
				"is_delete": Space_Delete_False,
			}).
			Limit(limit, number).
			OrderBy("space_id", "DESC"))
	if err != nil {
		return
	}
	spaces = rs.Rows()

	return
}

// get all spaces
func (s *Space) GetSpaces() (spaces []map[string]string, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().From(Table_Space_Name).Where(map[string]interface{}{
			"is_delete": Space_Delete_False,
		}))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// get space count
func (s *Space) CountSpaces() (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(
		db.AR().
			Select("count(*) as total").
			From(Table_Space_Name).
			Where(map[string]interface{}{
				"is_delete": Space_Delete_False,
			}))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

// get space count by keyword
func (s *Space) CountSpacesByKeyword(keyword string) (count int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("count(*) as total").From(Table_Space_Name).
		Where(map[string]interface{}{"is_delete":  Space_Delete_False}).
		WhereWrap(map[string]interface{}{"name LIKE": "%" + keyword + "%"}, "AND (", "").
		WhereWrap(map[string]interface{}{"description LIKE": "%" + keyword + "%"}, "OR", ")")
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	count = utils.Convert.StringToInt64(rs.Value("total"))
	return
}

// get space by name
func (s *Space) GetSpaceByLikeName(name string) (spaces []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"name Like": "%" + name + "%",
		"is_delete":     Space_Delete_False,
	}).Limit(0, 1))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// get space by many space_id
func (s *Space) GetSpaceBySpaceIds(spaceIds []string) (spaces []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Space_Name).Where(map[string]interface{}{
		"space_id":   spaceIds,
		"is_delete": Space_Delete_False,
	}))
	if err != nil {
		return
	}
	spaces = rs.Rows()
	return
}

// update space by name
func (s *Space) UpdateSpaceByName(space map[string]interface{}) (affect int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	space["update_time"] = time.Now().Unix()
	rs, err = db.Exec(db.AR().Update(Table_Space_Name, space, map[string]interface{}{
		"name": space["name"],
	}))
	if err != nil {
		return
	}
	affect = rs.RowsAffected
	return
}