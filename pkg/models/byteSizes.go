package models

// StandardBS is the byte size of Standard
const StandardBS = 1

// SectionBS is the byte size of Section
const SectionBS = 1

// GroupBS is the byte size of Group
const GroupBS = 1

// YearBS is the byte size of Year
const YearBS = 2

// TypeBS is the byte size of the Subject Type
const TypeBS = 1

// JoinNoBS is the byte size of joining no
const JoinNoBS = 2

// ClassIDBS is the byte size of ClassID
const ClassIDBS = StandardBS + SectionBS + GroupBS + YearBS

// ClassIDB is a byte array for teacherID
type ClassIDB [ClassIDBS]byte

// SubjectIDBS is the byte size of  SubjectID
const SubjectIDBS = StandardBS + TypeBS

// SubjectIDB is a byte array for teacherID
type SubjectIDB [SubjectIDBS]byte

// TeacherIDBS is the byte size if TeacherID
const TeacherIDBS = YearBS + JoinNoBS

// TeacherIDB is a byte array for teacherID
type TeacherIDB [TeacherIDBS]byte

// PeriodBS is byte size of period
const PeriodBS = ClassIDBS + SubjectIDBS + TeacherIDBS

// PeriodB is a byte array from period
type PeriodB [PeriodBS]byte
