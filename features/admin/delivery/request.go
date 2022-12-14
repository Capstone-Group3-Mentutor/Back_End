package delivery

import "be12/mentutor/features/admin"

type RegisterFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IdClass  uint   `json:"id_class" form:"id_class"`
	Role     string `json:"role" form:"role"`
}

type UpdateUserFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IdClass  uint   `json:"id_class" form:"id_class"`
	Images   string `json:"images" form:"images"`
}

type AddClassFormat struct {
	Class   string `json:"class_name" form:"class_name"`
}

type UpdateClassFormat struct {
	IdClass uint
	Class   string `json:"class_name" form:"class_name"`
	Status  string `json:"status" form:"status"`
}

func ToDomain(data RegisterFormat) admin.UserCore {
	return admin.UserCore{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		IdClass:  data.IdClass,
		Role:     data.Role,
	}
}

func ToDomainClass(data AddClassFormat) admin.ClassCore {
	return admin.ClassCore{
		ClassName: data.Class,
	}
}

func ToDomainUpdateUser(data UpdateUserFormat) admin.UserCore {
	return admin.UserCore{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		IdClass:  data.IdClass,
		Images:   data.Images,
	}
}

func ToDomainUpdateClass(data UpdateClassFormat) admin.ClassCore {
	return admin.ClassCore{
		IdClass:   data.IdClass,
		ClassName: data.Class,
		Status:    data.Status,
	}
}
