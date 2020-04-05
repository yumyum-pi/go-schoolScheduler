package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

// pIDToGeneID return the geneID from
// class index and periodsID
// TODO write test
func pIDToGeneID(classIndex, pID int) int {
	return (classIndex*models.MaxCap + pID)
}

// geneIDToPID returns the periodID from the geneID
// TODO write test
func geneIDToPID(geneID int) int {
	n := (geneID) / models.MaxCap
	return geneID - (n * models.MaxCap)
}

// inClass check is two geneID's are in the same class
// TODO write test
func inClass(gene1, gene2 int) bool {
	// calculate the class
	g1 := (gene1) / models.MaxCap
	g2 := (gene2) / models.MaxCap

	// compare and return the class
	return g1 == g2
}

// Chromosome is the collection of TimeTable of the whole school.
// arrange in an slice of period.
type Chromosome []models.Period

// Bytes return the byte value of teh chromosome
// TODO write test
func (chrmsm *Chromosome) Bytes() []byte {
	chrmsmLen := len(*chrmsm)
	chrmsmBS := models.PeriodByteSize * chrmsmLen // calculate the lenght of the byte
	b := make([]byte, chrmsmBS)                   // create an slice of bytes

	// loop through all the periods using the bytes
	for byteIndex := 0; byteIndex < chrmsmBS; byteIndex += models.PeriodByteSize {
		// calculate the period index from the byteIndex
		cIndex := byteIndex / models.PeriodByteSize

		// get the byte data of the periods
		pBytes := (*chrmsm)[cIndex].Bytes()
		// copy all the bytes from period byte data to chromosome byte data form the byteIndex
		copy(b[byteIndex:byteIndex+models.PeriodByteSize], pBytes[:])
	}
	return b
}

// InitS create the Inital chromosome from the classes with sequence
// TODO write test
func (chrmsm *Chromosome) InitS(cs *[]models.Class) {
	// loop through each class
	for classIndex, c := range *cs {
		pID := 0
		// loop thought each subject
		for _, s := range c.Subjects {
			// pID of the end of the same subject
			n := pID + s.ReqClasses

			// create a new period
			p := models.Period{
				ClassID:   c.ID,
				SubjectID: s.ID,
				TeacherID: s.TeacherID,
			}

			// loop until the pID == n
			for pID < n {
				// generate geneID for the period
				geneID := pIDToGeneID(classIndex, pID)
				// assign the period
				(*chrmsm)[geneID] = p
				// increate the pID and continue the loop
				pID++
			}
		}
	}
}

// InitR return the Inital chromosome from the initail chromosomes
func (chrmsm *Chromosome) InitR() Chromosome {
	// make a new chromosome var
	newChrsm := make(Chromosome, len(*chrmsm))

	// loop throught each class
	for geneID := 0; geneID < len((*chrmsm)); geneID += models.MaxCap {

		// create a new chromosome of a class length
		pReq := make(Chromosome, models.MaxCap)
		// copy the class period data to pReq
		copy(pReq[:], (*chrmsm)[geneID:(geneID+models.MaxCap)])

		//loop throght each period
		for pID := 0; pID < models.MaxCap; pID++ {
			// check if pReq is not empty
			if len(pReq) != 0 {
				// create a random no from 0 to pReq length
				randPID := utils.GenerateRandomInt(len(pReq), 10)

				// assign a random period to the new chromosome
				// with the randomly generated no.
				(newChrsm)[geneID+pID] = pReq[randPID]

				// delete the used pReq element
				deleteEml(&pReq, randPID)
			}
		}
	}
	return newChrsm
}

// TeacherAllocationConflict return the no. of allocated period of a teacher with same periodID
// TODO write test
func (chrmsm *Chromosome) TeacherAllocationConflict(geneID int, tID models.TeacherID) int {
	// create counter to store teacher's assigned period in perticuler pID
	// make the default value to 0
	n := 0

	pID := geneIDToPID(geneID) // calculate the periodID

	// loop through each class
	for gID := 0; gID < len(*chrmsm); gID += models.MaxCap {
		classGeneID := gID + pID
		// chekc if the period with the pID has the same teacher
		// check if the teacher is not empty
		if (*chrmsm)[classGeneID].TeacherID == tID && tID != (models.TeacherID{}) && classGeneID != geneID {
			// added one to the existing counter
			n++
		}
	}
	return n
}

// TeachersAllocationConflict return the no. of allocated period of a teacher with same periodID
// n1 = no. of periods allocated by teacher2 at pID1
// n2 = no. of periods allocated by teacher1 at pID2
// TODO write test and review
func (chrmsm *Chromosome) TeachersAllocationConflict(geneID1, geneID2 int) (n1, n2 int) {
	// get the ID of both the teacher assigned at the given geneIDs
	tID1, tID2 := (*chrmsm)[geneID1].TeacherID, (*chrmsm)[geneID2].TeacherID

	// calculate the periodID for both the teachers
	pID1, pID2 := geneIDToPID(geneID1), geneIDToPID(geneID2)

	// loop through each class
	for gID := 0; gID < len(*chrmsm); gID += models.MaxCap {
		// check if the period with the pID2 has the same teacher1
		// check if the teacher1 is not empty
		if (*chrmsm)[gID+pID2].TeacherID == tID1 && tID1 != (models.TeacherID{}) {
			// added one to the existing counter
			n1++
		}

		// chekc if the period with the pID2 has the same teacher
		// check if the teacher is not empty
		if (*chrmsm)[gID+pID1].TeacherID == tID2 && tID2 != (models.TeacherID{}) {
			// added one to the existing counter
			n2++
		}
	}
	return
}

// ErrorCheckM2 calculates the fitness of the chromosome and
// returns a numeric value that is comparable
// TODO write test
func (chrmsm *Chromosome) ErrorCheckM2() (errorRate float64, errGeneID []int) {
	// check a teacher is assgned a same period
	// loop throght each period
	for geneID := 0; geneID < len(*chrmsm); geneID++ {
		// get the no. times the teacher has been assigned in the given pID
		n := (*chrmsm).TeacherAllocationConflict(geneID, (*chrmsm)[geneID].TeacherID)

		// if the teacher is assigned more then once
		if n > 0 {
			// add the geneID to error list
			errGeneID = append(errGeneID, geneID)
		}
	}
	// calcuate the error rate
	errorRate = (float64(len(errGeneID)) / float64(len(*chrmsm))) * 100
	return

}

// ErrorCheckM3 calculates the fitness of the chromosome by
// check the no. of required subjects in each class
// and returns a bool
// TODO write test
func (chrmsm *Chromosome) ErrorCheckM3(c *[]models.Class) bool {
	// loop through classes
	for geneID := 0; geneID < len(*chrmsm); geneID += models.MaxCap {
		// calculate class index
		cIndex := geneID / models.MaxCap
		// create an map for subjecID and the assigned classes
		subAssigned := make(map[models.SubjectID]int)

		// loop through each class
		for pID := 0; pID < models.MaxCap; pID++ {
			// get subejct id of the period
			sID := (*chrmsm)[geneID+pID].SubjectID
			// increase the no. of assigned period of the subjectID
			subAssigned[sID]++
		}

		// check the assigned subjects
		// loop though all the subject in the class with the cIndex
		for _, sub := range (*c)[cIndex].Subjects {
			// check if the no. of assigned periods of the subject matches
			// the required no. of subeject periods
			if subAssigned[sub.ID] != sub.ReqClasses {
				// the no. of assigned periods of a subejct is != required periods
				return false
			}
		}
		// no mismatch found in the no. of assigend periods
		// and requried periods all the subjects in class
		// loop through next class
	}
	// no mismatch found in the no. of assigend periods
	// and requried periods all the subjects in all the classes
	return true

}

// ErrorHandleM1 uses inner class error swapping methords to minimize errors
// TODO write test
func (chrmsm *Chromosome) ErrorHandleM1() {
	// get the list of error geneID
	_, errList := (*chrmsm).ErrorCheckM2()

	// loop through all the error geneID
	for g1Index, gene1 := range errList {
		if gene1 == -1 {
			continue /// skip to the next loop
		}
		// loop through errList again to check for a replacement
		for g2Index, gene2 := range errList {
			// skip if same index || gene2 is swaped || or out of class
			if g2Index != g1Index && gene2 != -1 && inClass(gene1, gene2) {
				// if check swaping the genes will not create conlflits for both the geneIDs
				if (*chrmsm).SwapGeneSafe(gene1, gene2) {
					// changing the error geneID to -1 to be skipped in next loop
					errList[g1Index] = -1
				}
			}
		}
	}
}

// ErrorHandleM2 uses inner class swapping methords to minimize errors
func (chrmsm *Chromosome) ErrorHandleM2() {
	// get the list of error geneID
	_, errList := (*chrmsm).ErrorCheckM2()

	// loop through all the error geneID
	for g1Index, gene1ID := range errList {
		if gene1ID == -1 {
			continue // skip to the next loop
		}
		// find a swapable period
		gene2ID := chrmsm.SwapablePeriods(gene1ID)

		// check if the swapable period returned true id
		if gene2ID != -1 {
			// remvoe the gene1ID from the error list
			errList[g1Index] = -1

			// loop the error list and find gene2ID in the list
			for i := range errList {
				// check if the current error index has the gene2ID
				if errList[i] == gene2ID {
					// remove the gene2ID from the error list
					errList[i] = -1
				}
			}
			// if gene2ID is not found in the list then just continue
			// Swap the genes
			(*chrmsm).SwapGene(gene1ID, gene2ID)
		}
	}
}

// SwapablePeriods return the geneID of the period which can be easly repalce
// without allocation problems
func (chrmsm *Chromosome) SwapablePeriods(g1 int) int {
	classIndex := g1 / models.MaxCap       // calculate class index
	classID := classIndex * models.MaxCap  // calculate the 1st geneID of the current class
	nextClassID := classID + models.MaxCap // calculate the 1st geneID of the next class

	// loop through periods in the class
	for g2 := classID; g2 < nextClassID; g2++ {
		// get the no. of assigned period of the teacher of
		// assigned to gene1 and gene2 if they swtich placces
		n1, n2 := (*chrmsm).TeachersAllocationConflict(g1, g2)

		// check is no period is assigned to the teaher in others geneID
		if n1 == 0 && n2 == 0 {
			// return the current index
			return g2
		}
	}
	return -1
}

// SwapGeneSafe check and swaps two genes
// return true when the swap is successful
func (chrmsm *Chromosome) SwapGeneSafe(g1, g2 int) (conflict bool) {
	// check the allocation of gene2's teacher at gene 1
	n := (*chrmsm).TeacherAllocationConflict(g1, (*chrmsm)[g2].TeacherID)

	// if the gene2's teacher2 had not been allocated
	if n == 0 {
		// swap the genes
		(*chrmsm).SwapGene(g1, g2)
		return true
	}
	return false
}

// SwapGene swaps two genes
func (chrmsm *Chromosome) SwapGene(g1, g2 int) {
	(*chrmsm)[g1], (*chrmsm)[g2] = (*chrmsm)[g2], (*chrmsm)[g1]
}

// printL console log the chromosome in a list format
func (chrmsm *Chromosome) printL() {
	for i, g := range *chrmsm {
		fmt.Printf("i=%v\tcID=%v\tsID=%v\ttID%v\n", i, g.ClassID, g.SubjectID, g.TeacherID)
	}
}

// printFitnessErrs prints fitness errors to the console
func (chrmsm *Chromosome) printFitnessErrs(details bool) {
	errRate, list := (*chrmsm).ErrorCheckM2()

	if details {
		// print error list
		for _, geneID := range list {
			fmt.Printf("> Error: geneID=%v\tpid=%v\tcID=%v\tsID=%v\ttID=%v\n", geneID, geneIDToPID(geneID), (*chrmsm)[geneID].ClassID.Bytes(), (*chrmsm)[geneID].SubjectID.Bytes(), (*chrmsm)[geneID].TeacherID.Bytes())
		}
	}

	fmt.Printf("err=%v\\%v\tRate=%v\t\n", len(list), len(*chrmsm), errRate)
}

// deletes an element at the index "i"
func deleteEml(s *Chromosome, i int) {
	l := len(*s) - 1
	(*s)[i] = (*s)[l] // Copy last element to index i.
	(*s) = (*s)[:l]
}
