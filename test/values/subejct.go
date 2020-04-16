package values

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// SubjectIDs is the slice of SubjectID for test
var SubjectIDs = []models.SubjectID{
	{Standerd: [2]byte{1, 0}, Type: [4]byte{1, 1, 3, 0}}, // 0
	{Standerd: [2]byte{2, 9}, Type: [4]byte{3, 2, 2, 4}}, // 1
	{Standerd: [2]byte{3, 8}, Type: [4]byte{5, 4, 3, 3}}, // 2
	{Standerd: [2]byte{4, 7}, Type: [4]byte{4, 6, 5, 4}}, // 3
	{Standerd: [2]byte{5, 6}, Type: [4]byte{5, 5, 6, 0}}, // 4
	{Standerd: [2]byte{6, 5}, Type: [4]byte{0, 6, 6, 7}}, // 5
	{Standerd: [2]byte{7, 4}, Type: [4]byte{8, 0, 7, 7}}, // 6
	{Standerd: [2]byte{8, 3}, Type: [4]byte{8, 9, 0, 8}}, // 7
	{Standerd: [2]byte{9, 2}, Type: [4]byte{8, 9, 0, 8}}, // 8
	{Standerd: [2]byte{0, 1}, Type: [4]byte{8, 9, 0, 8}}, // 9
}

// Subjects is a slice of Subject for test
var Subjects = []models.Subject{
	{ID: SubjectIDs[0], TeacherID: TeacherIDs[0], Req: 6}, // 6
	{ID: SubjectIDs[1], TeacherID: TeacherIDs[1], Req: 6}, // 12
	{ID: SubjectIDs[2], TeacherID: TeacherIDs[2], Req: 6}, // 18
	{ID: SubjectIDs[3], TeacherID: TeacherIDs[3], Req: 6}, // 24
	{ID: SubjectIDs[4], TeacherID: TeacherIDs[4], Req: 6}, // 30
	{ID: SubjectIDs[5], TeacherID: TeacherIDs[5], Req: 6}, // 36
	{ID: SubjectIDs[6], TeacherID: TeacherIDs[6], Req: 3}, // 39
	{ID: SubjectIDs[7], TeacherID: TeacherIDs[7], Req: 3}, // 42
	{ID: SubjectIDs[8], TeacherID: TeacherIDs[8], Req: 3}, // 45
	{ID: SubjectIDs[9], TeacherID: TeacherIDs[9], Req: 3}, // 48
}
