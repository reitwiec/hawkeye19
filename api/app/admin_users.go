package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

const perPage = 10

func (hawk *App) editUser(w http.ResponseWriter, r *http.Request) {
	userData := User{}
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		ResponseWriter(false, "User not edited.", nil, http.StatusBadRequest, w)
		return
	}
	userData.Username = strings.TrimSpace(userData.Username)
	userData.Name = strings.TrimSpace(userData.Name)
	userData.Password = strings.TrimSpace(userData.Password)
	userData.Email = strings.TrimSpace(userData.Email)
	userData.Tel = strings.TrimSpace(userData.Tel)
	userData.College = strings.TrimSpace(userData.College)
	userData.SideQuest = strings.TrimSpace(userData.SideQuest)
	userData.UnlockOrder = strings.TrimSpace(userData.UnlockOrder)

	tx := hawk.DB.Begin()
	err = tx.Where("id = ? ", userData.ID).First(&User{}).Updates(User{
		Name:        userData.Name,
		Access:      userData.Access,
		Email:       userData.Email,
		Tel:         userData.Tel,
		College:     userData.College,
		Region1:     userData.Region1,
		Region2:     userData.Region2,
		Region3:     userData.Region3,
		Region4:     userData.Region4,
		Region5:     userData.Region5,
		Banned:      userData.Banned,
		SideQuest:   userData.SideQuest,
		UnlockOrder: userData.UnlockOrder,
	}).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("User not found")
			ResponseWriter(false, "User not found", nil, http.StatusBadRequest, w)
			return
		}
		fmt.Printf("Database error %s", err.Error())
		ResponseWriter(false, "User not edited", nil, http.StatusInternalServerError, w)
		return
	}
	tx.Commit()
	ResponseWriter(true, "User edited", nil, http.StatusOK, w)

}

func (hawk *App) deleteUser(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "User not deleted.", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]

	tx := hawk.DB.Begin()
	u := User{}
	err := tx.Where("id = ?", key).First(&u).Delete(u).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "User not found", nil, http.StatusBadRequest, w)
		} else {
			fmt.Printf("Database error %s ", err.Error())
			ResponseWriter(false, "User not deleted", nil, http.StatusInternalServerError, w)
		}
		return
	}
	tx.Commit()
	ResponseWriter(true, "User deleted successfully", nil, http.StatusOK, w)
}

func (hawk *App) banUser(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "User not banned", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]
	tx := hawk.DB.Begin()
	u := User{}
	err := tx.Where("ID = ?", key).First(&u).Update("Banned", 1).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "User not found", nil, http.StatusBadRequest, w)
		} else {
			fmt.Printf("Database error %s ", err.Error())
			ResponseWriter(false, "User not banned", nil, http.StatusInternalServerError, w)
		}
		return
	}
	tx.Commit()
	ResponseWriter(true, "User banned successfully", nil, http.StatusOK, w)
}

func (hawk *App) unbanUser(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "User not unbanned", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]
	tx := hawk.DB.Begin()
	u := User{}
	err := tx.Where("ID = ?", key).First(&u).Update("Banned", 0).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "User not found", nil, http.StatusBadRequest, w)
		} else {
			fmt.Printf("Database error %s ", err.Error())
			ResponseWriter(false, "User not unbanned", nil, http.StatusInternalServerError, w)
		}
		return
	}
	tx.Commit()
	ResponseWriter(true, "User unbanned successfully", nil, http.StatusOK, w)
}

func (hawk *App) makeAdmin(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "User not made admin", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]
	tx := hawk.DB.Begin()
	u := User{}
	err := tx.Where("ID = ?", key).First(&u).Update("Access", 1).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "User not found", nil, http.StatusBadRequest, w)
		} else {
			fmt.Printf("Database error %s ", err.Error())
			ResponseWriter(false, "User not made admin", nil, http.StatusInternalServerError, w)
		}
		return
	}
	tx.Commit()
	ResponseWriter(true, "User made admin successfully", nil, http.StatusOK, w)
}
func (hawk *App) revokeAdmin(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		ResponseWriter(false, "Admin not revoked", nil, http.StatusBadRequest, w)
		return
	}
	key := keys[0]
	tx := hawk.DB.Begin()
	u := User{}
	err := tx.Where("ID = ?", key).First(&u).Update("Access", 0).Error
	if err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			ResponseWriter(false, "Admin not revoked", nil, http.StatusBadRequest, w)
		} else {
			fmt.Printf("Database error %s ", err.Error())
			ResponseWriter(false, "Admin not revoked", nil, http.StatusInternalServerError, w)
		}
		return
	}
	tx.Commit()
	ResponseWriter(true, "Admin revoked", nil, http.StatusOK, w)
}


func (hawk *App) listUsers(w http.ResponseWriter, r *http.Request) {
	pgs, ok := r.URL.Query()["page"]
	if !ok || len(pgs[0]) < 1 {
		ResponseWriter(false, "Cant list users", nil, http.StatusBadRequest, w)
		return
	}
	page, err := strconv.Atoi(pgs[0])
	var users []User
	offset := (page - 1) * perPage
	err = hawk.DB.Select("id, username, name, email, tel, college, banned, points, region1, region2, region3, region4, region5, unlock_order, side_quest").Limit(perPage).Offset(offset).Order("region5 desc").Find(&users).Error
	if err != nil {
		fmt.Println ("Error in database")
		ResponseWriter(false, "Cannot list users", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "List of users", users, http.StatusOK, w)
}


func (hawk *App) searchUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.URL.Query()["user"]
	if !ok || len(user[0]) < 1 {
		ResponseWriter(false, "Cant find user", nil, http.StatusBadRequest, w)
		return
	}
	user[0] = strings.TrimSpace(user[0])
	var users []User
	arg := "%" + user[0] + "%"
	err := hawk.DB.Select("id, username, name, email, tel, college, banned, points, region1, region2, region3, region4, region5, unlock_order, side_quest").Where("username LIKE ? OR name LIKE ?", arg, arg).Find(&users).Error
	if err != nil {
		fmt.Println("Cannot find user")
		ResponseWriter(false, "Cannot find user", nil, http.StatusInternalServerError, w)
		return
	}
	/*
	var nameMatches []User
	err = hawk.DB.Select("id, username, name, email, tel, college, banned, points, region1, region2, region3, region4, region5").Where("name LIKE ?", "%"+user[0]+"%").Find(&nameMatches).Error
	if err != nil {
		fmt.Println ("Cannot find user")
		ResponseWriter(false, "Cannot find user", nil, http.StatusInternalServerError, w)
		return
	}

	users = append(users, nameMatches...)
	*/
	ResponseWriter(true, "User found", users, http.StatusOK, w)
}

func (hawk *App) userLogs(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	if len(keys) < 2 {
		ResponseWriter(false, "Cannot log questions. Less Parame", nil, http.StatusBadRequest, w)
		return
	}
	page, err := strconv.Atoi(keys["page"][0])
	if err != nil {
		ResponseWriter(false, "Parameters not valid. Cannot log answers.", nil, http.StatusBadRequest, w)
		return
	}
	id, err := strconv.Atoi(keys["id"][0])
	if err != nil {
		ResponseWriter(false, "Parameters not valid. Cannot log answers.", nil, http.StatusBadRequest, w)
		return
	}
	var answers []Attempt
	offset := (page - 1) * perPage

	err = hawk.DB.Where("user = ? ", id).Order("id desc").Offset(offset).Limit(perPage).Find(&answers).Error
	if err != nil {
		fmt.Println ("Cannot log questions")
		ResponseWriter(false, "Cannot log questions", nil, http.StatusInternalServerError, w)
		return
	}
	ResponseWriter(true, "Requested answer logs", answers, http.StatusOK, w)
}
