package schemas

type CreateStudentSchema struct {
	Name   string  `json:"name"  binding:"required"`
	Grade  uint8   `json:"grade" binding:"required"`
	Rating float32 `json:"rating" binding:"required"`
}

type UpdateStudentSchema struct {
	Name  string `json:"name"`
	Grade uint8  `json:"grade"`
}
