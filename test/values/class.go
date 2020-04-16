package values

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// ClassIDs is a slice of ClassID for test
var ClassIDs = []models.ClassID{
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 1}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 0
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 2}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 1
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 3}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 2
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 4}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 3
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 5}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 4
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 6}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 5
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 7}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 6
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 8}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 7
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 9}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 8
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{1, 0}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 9
}

// Classes is a slice of class for test
var Classes = models.Classes{
	models.Class{ID: ClassIDs[0], Subjects: Subjects[:1], Capacity: 42}, // 0-9: 6 - 48
	models.Class{ID: ClassIDs[1], Subjects: Subjects[:2], Capacity: 36}, // 0-4: 12 - 48
	models.Class{ID: ClassIDs[2], Subjects: Subjects[:3], Capacity: 30}, // 0-5: 18 - 48
	models.Class{ID: ClassIDs[3], Subjects: Subjects[:4], Capacity: 24}, // 0-4: 24 - 48
	models.Class{ID: ClassIDs[4], Subjects: Subjects[:5], Capacity: 18}, // 0-3: 30 - 48
	models.Class{ID: ClassIDs[5], Subjects: Subjects[:6], Capacity: 12}, // 0-6: 36 - 48
	models.Class{ID: ClassIDs[6], Subjects: Subjects[:7], Capacity: 9},  // 0-5: 39 - 48
	models.Class{ID: ClassIDs[7], Subjects: Subjects[:8], Capacity: 6},  // 0-2: 42 - 48
	models.Class{ID: ClassIDs[8], Subjects: Subjects[:9], Capacity: 3},  // 0-3: 45 - 48
	models.Class{ID: ClassIDs[9], Subjects: Subjects[:], Capacity: 0},   // 0-9: 48 - 48
}
