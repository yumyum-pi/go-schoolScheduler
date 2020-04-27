package models

// Class

// TClassIDL is a slice of ClassID for test
var TClassIDL = []ClassID{
	{Year: [YearBS]byte{7, 228}, Standard: 1, Section: 0},  // 0
	{Year: [YearBS]byte{7, 228}, Standard: 2, Section: 1},  // 1
	{Year: [YearBS]byte{7, 228}, Standard: 3, Section: 2},  // 2
	{Year: [YearBS]byte{7, 228}, Standard: 4, Section: 3},  // 3
	{Year: [YearBS]byte{7, 228}, Standard: 5, Section: 4},  // 4
	{Year: [YearBS]byte{7, 228}, Standard: 6, Section: 5},  // 5
	{Year: [YearBS]byte{7, 228}, Standard: 7, Section: 6},  // 6
	{Year: [YearBS]byte{7, 228}, Standard: 8, Section: 7},  // 7
	{Year: [YearBS]byte{7, 228}, Standard: 9, Section: 8},  // 8
	{Year: [YearBS]byte{7, 228}, Standard: 10, Section: 9}, // 9
}

// TClassIDBL is a slice of test bytes of classIDs
var TClassIDBL = []ClassIDB{
	{7, 228, 1, 0},  // 0
	{7, 228, 2, 1},  // 1
	{7, 228, 3, 2},  // 2
	{7, 228, 4, 3},  // 3
	{7, 228, 5, 4},  // 4
	{7, 228, 6, 5},  // 5
	{7, 228, 7, 6},  // 6
	{7, 228, 8, 7},  // 7
	{7, 228, 9, 8},  // 8
	{7, 228, 10, 9}, // 9
}

// TClassL is a slice of class for test
var TClassL = []Class{
	{ID: TClassIDL[0], Subjects: TSubjectL[:1], Capacity: 42}, // 0-9: 6 - 48
	{ID: TClassIDL[1], Subjects: TSubjectL[:2], Capacity: 36}, // 0-4: 12 - 48
	{ID: TClassIDL[2], Subjects: TSubjectL[:3], Capacity: 30}, // 0-5: 18 - 48
	{ID: TClassIDL[3], Subjects: TSubjectL[:4], Capacity: 24}, // 0-4: 24 - 48
	{ID: TClassIDL[4], Subjects: TSubjectL[:5], Capacity: 18}, // 0-3: 30 - 48
	{ID: TClassIDL[5], Subjects: TSubjectL[:6], Capacity: 12}, // 0-6: 36 - 48
	{ID: TClassIDL[6], Subjects: TSubjectL[:7], Capacity: 9},  // 0-5: 39 - 48
	{ID: TClassIDL[7], Subjects: TSubjectL[:8], Capacity: 6},  // 0-2: 42 - 48
	{ID: TClassIDL[8], Subjects: TSubjectL[:9], Capacity: 3},  // 0-3: 45 - 48
	{ID: TClassIDL[9], Subjects: TSubjectL[:], Capacity: 0},   // 0-9: 48 - 48
}

//Subject

// TSubjectIDL is the slice of SubjectID for test
var TSubjectIDL = []SubjectID{
	{Standard: 1, Type: 0}, // 0
	{Standard: 2, Type: 1}, // 1
	{Standard: 3, Type: 2}, // 2
	{Standard: 4, Type: 3}, // 3
	{Standard: 5, Type: 4}, // 4
	{Standard: 6, Type: 5}, // 5
	{Standard: 7, Type: 6}, // 6
	{Standard: 8, Type: 7}, // 7
	{Standard: 9, Type: 8}, // 8
	{Standard: 0, Type: 9}, // 9
}

// TSubjectIDBL list of subjectID bytes slice
var TSubjectIDBL = []SubjectIDB{
	{1, 0}, // 0
	{2, 1}, // 1
	{3, 2}, // 2
	{4, 3}, // 3
	{5, 4}, // 4
	{6, 5}, // 5
	{7, 6}, // 6
	{8, 7}, // 7
	{9, 8}, // 8
	{0, 9}, // 9
}

// TSubjectL is a slice of Subject for test
var TSubjectL = []Subject{
	{ID: TSubjectIDL[0], TeacherID: TTeacherIDL[0], Req: 6}, // 6
	{ID: TSubjectIDL[1], TeacherID: TTeacherIDL[1], Req: 6}, // 12
	{ID: TSubjectIDL[2], TeacherID: TTeacherIDL[2], Req: 6}, // 18
	{ID: TSubjectIDL[3], TeacherID: TTeacherIDL[3], Req: 6}, // 24
	{ID: TSubjectIDL[4], TeacherID: TTeacherIDL[4], Req: 6}, // 30
	{ID: TSubjectIDL[5], TeacherID: TTeacherIDL[5], Req: 6}, // 36
	{ID: TSubjectIDL[6], TeacherID: TTeacherIDL[6], Req: 3}, // 39
	{ID: TSubjectIDL[7], TeacherID: TTeacherIDL[7], Req: 3}, // 42
	{ID: TSubjectIDL[8], TeacherID: TTeacherIDL[8], Req: 3}, // 45
	{ID: TSubjectIDL[9], TeacherID: TTeacherIDL[9], Req: 3}, // 48
}

//Teacher

// TTeacherIDL is a collection of teacherID
var TTeacherIDL = []TeacherID{
	{Year: [YearBS]byte{7, 228}, JoinNo: 0}, // 0
	{Year: [YearBS]byte{7, 228}, JoinNo: 1}, // 1
	{Year: [YearBS]byte{7, 228}, JoinNo: 2}, // 2
	{Year: [YearBS]byte{7, 228}, JoinNo: 3}, // 3
	{Year: [YearBS]byte{7, 228}, JoinNo: 4}, // 4
	{Year: [YearBS]byte{7, 228}, JoinNo: 5}, // 5
	{Year: [YearBS]byte{7, 228}, JoinNo: 6}, // 6
	{Year: [YearBS]byte{7, 228}, JoinNo: 7}, // 7
	{Year: [YearBS]byte{7, 228}, JoinNo: 8}, // 8
	{Year: [YearBS]byte{7, 228}, JoinNo: 9}, // 9
}

// TTeacherIDBL is a collection of teacherID
var TTeacherIDBL = []TeacherIDB{
	{7, 228, 0}, // 0
	{7, 228, 1}, // 1
	{7, 228, 2}, // 2
	{7, 228, 3}, // 3
	{7, 228, 4}, // 4
	{7, 228, 5}, // 5
	{7, 228, 6}, // 6
	{7, 228, 7}, // 7
	{7, 228, 8}, // 8
	{7, 228, 9}, // 9
}

// TClassAssignedL is a list of assigned classes for teachers
var TClassAssignedL = []ClassAssigned{
	{TSubjectIDL[0], TClassIDL[0], 6}, // 6
	{TSubjectIDL[1], TClassIDL[1], 6}, // 12
	{TSubjectIDL[2], TClassIDL[2], 6}, // 18
	{TSubjectIDL[3], TClassIDL[3], 6}, // 24
	{TSubjectIDL[4], TClassIDL[4], 6}, // 30
	{TSubjectIDL[5], TClassIDL[5], 6}, // 36
	{TSubjectIDL[6], TClassIDL[6], 3}, // 39
	{TSubjectIDL[7], TClassIDL[7], 3}, // 42
	{TSubjectIDL[8], TClassIDL[8], 3}, // 45
	{TSubjectIDL[9], TClassIDL[9], 3}, // 48
}

// TTeacherL is a slice of teachers
var TTeacherL = []Teacher{
	{TTeacherIDL[0], TSubjectIDL[:1], TClassAssignedL[:1], 42}, //
	{TTeacherIDL[1], TSubjectIDL[:2], TClassAssignedL[:2], 36},
	{TTeacherIDL[2], TSubjectIDL[:3], TClassAssignedL[:3], 30},
	{TTeacherIDL[3], TSubjectIDL[:4], TClassAssignedL[:4], 24},
	{TTeacherIDL[4], TSubjectIDL[:5], TClassAssignedL[:5], 18},
	{TTeacherIDL[5], TSubjectIDL[:6], TClassAssignedL[:6], 12},
	{TTeacherIDL[6], TSubjectIDL[:7], TClassAssignedL[:7], 9},
	{TTeacherIDL[7], TSubjectIDL[:8], TClassAssignedL[:8], 6},
	{TTeacherIDL[8], TSubjectIDL[:9], TClassAssignedL[:9], 3},
	{TTeacherIDL[9], TSubjectIDL[:], TClassAssignedL[:], 0},
}

// Periods
