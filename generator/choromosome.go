package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
	"github.com/yumyum-pi/go-schoolScheduler/utils/stats"
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
type Chromosome struct {
	Genes   []models.Period // array of periods
	Name    string          // name of genes
	ErrIDs  []int           // slice of IDs that are conflicted
	ErrNo   int             // stores the no. of errors in the chromosomes
	Fitness int             // fitness of the chromosome
}

// Length return the lenght of genes
func (chrmsm *Chromosome) Length() int {
	return len((*chrmsm).Genes)
}

// Bytes return the byte value of teh chromosome
// TODO write test
func (chrmsm *Chromosome) Bytes() []byte {
	chrmsmLen := (*chrmsm).Length()
	chrmsmBS := models.PeriodByteSize * chrmsmLen // calculate the lenght of the byte
	b := make([]byte, chrmsmBS)                   // create an slice of bytes

	// loop through all the periods using the bytes
	for byteIndex := 0; byteIndex < chrmsmBS; byteIndex += models.PeriodByteSize {
		// calculate the period index from the byteIndex
		cIndex := byteIndex / models.PeriodByteSize

		// get the byte data of the periods
		pBytes := (*chrmsm).Genes[cIndex].Bytes()
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
				(*chrmsm).Genes[geneID] = p
				// increate the pID and continue the loop
				pID++
			}
		}
	}
}

// InitR return the Inital chromosome from the initail chromosomes
func (chrmsm *Chromosome) InitR() Chromosome {
	// make a new chromosome
	var newChrsm Chromosome
	geneLength := (*chrmsm).Length()
	newChrsm.Genes = make([]models.Period, geneLength)

	// loop throught each class
	for geneID := 0; geneID < geneLength; geneID += models.MaxCap {

		// create a new chromosome of a class length
		pReq := make([]models.Period, models.MaxCap)
		// copy the class period data to pReq
		copy(pReq[:], (*chrmsm).Genes[geneID:(geneID+models.MaxCap)])

		//loop throght each period
		for pID := 0; pID < models.MaxCap; pID++ {
			// check if pReq is not empty
			if len(pReq) != 0 {
				// create a random no from 0 to pReq length
				randPID := utils.GenerateRandomInt(len(pReq), 10)

				// assign a random period to the new chromosome
				// with the randomly generated no.
				(newChrsm).Genes[geneID+pID] = pReq[randPID]

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

	pID := geneIDToPID(geneID)       // calculate the periodID
	geneLength := (*chrmsm).Length() // get the length of the gene

	// loop through each class
	for gID := 0; gID < geneLength; gID += models.MaxCap {
		classGeneID := gID + pID
		// chekc if the period with the pID has the same teacher
		// check if the teacher is not empty
		if (*chrmsm).Genes[classGeneID].TeacherID == tID && tID != (models.TeacherID{}) && classGeneID != geneID {
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
	tID1, tID2 := (*chrmsm).Genes[geneID1].TeacherID, (*chrmsm).Genes[geneID2].TeacherID

	// calculate the periodID for both the teachers
	pID1, pID2 := geneIDToPID(geneID1), geneIDToPID(geneID2)
	geneLength := (*chrmsm).Length() // get the length of the gene
	// loop through each class
	for gID := 0; gID < geneLength; gID += models.MaxCap {
		// check if the period with the pID2 has the same teacher1
		// check if the teacher1 is not empty
		if (*chrmsm).Genes[gID+pID2].TeacherID == tID1 && tID1 != (models.TeacherID{}) {
			// added one to the existing counter
			n1++
		}

		// chekc if the period with the pID2 has the same teacher
		// check if the teacher is not empty
		if (*chrmsm).Genes[gID+pID1].TeacherID == tID2 && tID2 != (models.TeacherID{}) {
			// added one to the existing counter
			n2++
		}
	}
	return
}

// ErrorCheck calculates the fitness of the chromosome and
// update the ErrIDs in of the chromosome
// TODO write test
func (chrmsm *Chromosome) ErrorCheck() {
	geneLength := (*chrmsm).Length() // get the length of the gene
	var err []int

	// check a teacher is assgned a same period
	// loop throght each period
	for geneID := 0; geneID < geneLength; geneID++ {
		// get the no. times the teacher has been assigned in the given pID
		n := (*chrmsm).TeacherAllocationConflict(geneID, (*chrmsm).Genes[geneID].TeacherID)

		// if the teacher is assigned more then once
		if n > 0 {
			// add the geneID to error list
			err = append(err, geneID)
		}
	}
	(*chrmsm).ErrIDs = make([]int, len(err))

	// reassign the error
	copy((*chrmsm).ErrIDs, err)
	(*chrmsm).ErrNo = len(err)

}

// ErrorCheckM3 calculates the fitness of the chromosome by
// check the no. of required subjects in each class
// and returns a bool
// TODO write test
func (chrmsm *Chromosome) ErrorCheckM3(c *[]models.Class) bool {
	geneLength := (*chrmsm).Length() // get the length of the gene

	// loop through classes
	for geneID := 0; geneID < geneLength; geneID += models.MaxCap {
		// calculate class index
		cIndex := geneID / models.MaxCap
		// create an map for subjecID and the assigned classes
		subAssigned := make(map[models.SubjectID]int)

		// loop through each class
		for pID := 0; pID < models.MaxCap; pID++ {
			// get subejct id of the period
			sID := (*chrmsm).Genes[geneID+pID].SubjectID
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

// ErrorHandle tries to resolve the error in gene by the using all the available methods
func (chrmsm *Chromosome) ErrorHandle() {
	(*chrmsm).ErrorCheck()    // check error
	(*chrmsm).ErrorHandleM1() // use methord1
	(*chrmsm).ErrorHandleM2() // use methord2
	(*chrmsm).ErrNo = len((*chrmsm).ErrIDs)
}

// ErrorHandleM1 uses inner class error swapping methords to minimize errors
// TODO write test
func (chrmsm *Chromosome) ErrorHandleM1() {
	// get the list of error geneID
	var err []int

	// loop through all the error geneID
	for g1Index, gene1 := range (*chrmsm).ErrIDs {
		if gene1 == -1 {
			continue /// skip to the next loop
		}
		// loop through errList again to check for a replacement
		for g2Index, gene2 := range (*chrmsm).ErrIDs {
			// skip if same index || gene2 is swaped || or out of class
			if g2Index != g1Index && gene2 != -1 && inClass(gene1, gene2) {
				// if check swaping the genes will not create conlflits for both the geneIDs
				if (*chrmsm).SwapGeneSafe(gene1, gene2) {
					// changing the error geneID to -1 to be skipped in next loop
					(*chrmsm).ErrIDs[g1Index] = -1
					continue /// skip to the next loop
				}
			}
		}

		// no match is found in the loop
		err = append(err, gene1) //
	}
	(*chrmsm).ErrIDs = make([]int, len(err))

	// reassign the error
	copy((*chrmsm).ErrIDs, err)
}

// ErrorHandleM2 uses inner class swapping methords to minimize errors
func (chrmsm *Chromosome) ErrorHandleM2() {
	// create the a slice to store err geneIDs
	var err []int

	// loop through all the error geneID
	for g1Index, gene1ID := range (*chrmsm).ErrIDs {
		if gene1ID == -1 {
			continue // skip to the next loop
		}
		// find a swapable period
		gene2ID := chrmsm.SwapablePeriods(gene1ID)

		// check if the swapable period returned true id
		if gene2ID != -1 {
			// remvoe the gene1ID from the error list
			(*chrmsm).ErrIDs[g1Index] = -1

			// loop the error list and find gene2ID in the list
			for i := range (*chrmsm).ErrIDs {
				// check if the current error index has the gene2ID
				if (*chrmsm).ErrIDs[i] == gene2ID {
					// remove the gene2ID from the error list
					(*chrmsm).ErrIDs[i] = -1
					break // exit the loop
				}
			}
			// if gene2ID is not found in the list then just continue
			// Swap the genes
			(*chrmsm).SwapGene(gene1ID, gene2ID)
		} else {
			err = append(err, gene1ID)
		}
	}

	// reassign the error
	(*chrmsm).ErrIDs = make([]int, len(err))

	// reassign the error
	copy((*chrmsm).ErrIDs, err)

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
	n := (*chrmsm).TeacherAllocationConflict(g1, (*chrmsm).Genes[g2].TeacherID)

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
	(*chrmsm).Genes[g1], (*chrmsm).Genes[g2] = (*chrmsm).Genes[g2], (*chrmsm).Genes[g1]
}

// CalFitness calculate the distribution of teacher and periods in the timetable
func (chrmsm *Chromosome) CalFitness(info bool) {
	// generate teacher data
	teacherTTList := (*chrmsm).teacherTTBool()                        // get the slice of periods assigned for each teacher
	teacherPeriodDist := (*chrmsm).teachersPeriodDist(&teacherTTList) // get the list of distribution calculation

	// create a mega list of distribution data from teacher and class subject

	fitness := 0
	// loop through all the data and calcuate variance
	for tID, diffList := range teacherPeriodDist {
		v := stats.VarianceInt(diffList) * 100
		fitness += v
		if info {
			diffListLen := len(diffList)
			fp := models.MaxCap - diffListLen
			fmt.Printf("len=%v,\tfree=%v,\tsum=%v,\tv=%v\ttID=%v\n", len(diffList), fp, stats.Sum(diffList), v, tID.Bytes())
		}
	}

	(*chrmsm).Fitness = fitness
}

// teacherTTBool creates and return map data of all teacher's and their timetable.
// The timetable is an arry of bool with a fixed length of MaxCap
// bool represents an assigned class for the teacher
func (chrmsm *Chromosome) teacherTTBool() map[models.TeacherID][models.MaxCap]bool {
	// make an empty map of arry of bool with a fixed length of MaxCap
	// each element in the map is key by teacherID
	teacherTTList := make(map[models.TeacherID][models.MaxCap]bool)

	// loop through all the genes
	for geneID, gene := range (*chrmsm).Genes {
		pID := geneIDToPID(geneID) // convenrt geneID to periodID

		// check if teacherID of the gene is not empty
		if gene.TeacherID != (models.TeacherID{}) {
			// check the teacher has assigned classes to current periodID
			if teacherTTList[gene.TeacherID][pID] {
				// this this a conflict and this issue needs to be resolved
				// TODO throw a proper error
				// Use panic, may be look into it so that the current work flow does not have
				// check error each time running this method

				// For now, write an error to the console
				fmt.Printf("> Error: Teacher no empty at periods=%v\n", pID)
				// continue on to the next gene
				continue
			}
			// The teacher has no assigned period
			tt := teacherTTList[gene.TeacherID] // get the timetbale of assigned periods
			tt[pID] = true                      // set the current periods to assigned
			teacherTTList[gene.TeacherID] = tt  // reassign the list to the teacherID
		}
	}

	return teacherTTList
}

// teacherPeriodDist calculates and returns a map data of slice of int
// that contains the no. of free periods b/w each assigned periods.
// The map is keyed by TeacherID
func (chrmsm *Chromosome) teachersPeriodDist(teacherTTList *map[models.TeacherID][models.MaxCap]bool) map[models.TeacherID][]int {
	// make can empty map of slice of int
	// each element in the map is key by teacherID
	teacherPeriodDist := make(map[models.TeacherID][]int)

	// loop through each teacher
	for tID, tt := range *teacherTTList {
		// make slice of fixed capacity of MaxCap to store the no. of free periods between each assigned peirod
		diffList := make([]int, models.MaxCap)

		lastAssignedPeriod := 0 // stored the periodID of last assigned periods
		nPeriod := 0            // no. of assigned period
		nFreePeriod := 0
		// loop though all the period
		// calculate the no. of free periods between each assgined periods
		for pID, period := range tt {
			// check if period is assigned
			if period {
				// calculate the difference between the current assigned period & last assigned period
				diffList[nPeriod] = (pID - lastAssignedPeriod)
				nPeriod++                // increasing by one to calculate the next index id for diffList
				lastAssignedPeriod = pID // change last periodID to current periodID
			} else {
				nFreePeriod++
			}
		}

		diffList = diffList[:(nPeriod)]   // remove the unassigned elements
		teacherPeriodDist[tID] = diffList // assign the data to the map with the given teacherID
	}

	return teacherPeriodDist
}

// deletes an element at the index "i"
func deleteEml(s *[]models.Period, i int) {
	l := len(*s) - 1
	(*s)[i] = (*s)[l] // Copy last element to index i.
	(*s) = (*s)[:l]
}
