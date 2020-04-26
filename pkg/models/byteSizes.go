package models

// StandardBS is the byte size of Standard
const StandardBS = 2

// SectionBS is the byte size of Section
const SectionBS = 2

// GroupBS is the byte size of Group
const GroupBS = 2

// YearBS is the byte size of Year
const YearBS = 4

// TypeBS is the byte size of the Subject Type
const TypeBS = 4

// JoinNoBS is the byte size of joining no
const JoinNoBS = 4

// ClassIDBS is the byte size of ClassID
const ClassIDBS = StandardBS + SectionBS + GroupBS + YearBS

// SubjectIDBS is the byte size of  SubjectID
const SubjectIDBS = StandardBS + TypeBS

// TeacherIDBS is the byte size if TeacherID
const TeacherIDBS = YearBS + JoinNoBS

// PeriodByteSize is byte size of period
const PeriodByteSize = ClassIDBS + SubjectIDBS + TeacherIDBS
