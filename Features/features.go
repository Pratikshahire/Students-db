package features

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Student struct {
	Name   string
	RollNo string
	Marks  int
}

type Class struct {
	Engineers []Student
}

func inputName() string {
	fmt.Println("Enter your first name")
	var name string
	fmt.Scanln(&name)
	return name
}

func inputRoll() string {
	fmt.Println("Enter your roll no")
	var roll string
	fmt.Scanln(&roll)
	return roll
}

func inputMarks() int {
	fmt.Println("Enter your marks")
	var marks int
	fmt.Scanln(&marks)
	return marks
}

func (c *Class) NewClass() {

	var studentsArray []Student
	file, err := os.Open("db.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		marks, _ := strconv.Atoi(fields[2])

		entry := Student{
			Name:   fields[0],
			RollNo: fields[1],
			Marks:  marks,
		}

		studentsArray = append(studentsArray, entry)
	}

	c.Engineers = studentsArray
}

func (c *Class) AddStudent() {

	var wg sync.WaitGroup

	name := inputName()

	roll := inputRoll()

	marks := inputMarks()

	entry := Student{
		Name:   name,
		RollNo: roll,
		Marks:  marks,
	}

	c.Engineers = append(c.Engineers, entry)

	fmt.Println("Student Added Successfully")
	fmt.Println()

	wg.Add(1)
	go c.writeToFile(&wg)
	wg.Wait()
}

func (c *Class) ShowStudents() {
	for _, student := range c.Engineers {
		name, roll, marks := student.Name, student.RollNo, student.Marks
		fmt.Printf("Name: %s\nRoll No: %s\nMarks: %d\n", name, roll, marks)
		fmt.Println()
	}
}

func (c Class) UpdateStudent() {

	var wg sync.WaitGroup

	roll := inputRoll()

	index := 0

	for ind, stu := range c.Engineers {
		if stu.RollNo == roll {
			index = ind
			break
		}
	}

	newName := inputName()
	newMarks := inputMarks()

	c.Engineers[index] = Student{
		RollNo: roll,
		Name:   newName,
		Marks:  newMarks,
	}

	wg.Add(1)
	go c.writeToFile(&wg)
	wg.Wait()

}

func (c *Class) DeleteStudent() {

	var wg sync.WaitGroup

	roll := inputRoll()
	index := 0

	for idx, student := range c.Engineers {
		if student.RollNo == roll {
			index = idx
		}
	}

	c.Engineers = append(c.Engineers[:index], c.Engineers[index+1:]...)

	wg.Add(1)
	go c.writeToFile(&wg)
	wg.Wait()
}

func (c *Class) writeToFile(wg *sync.WaitGroup) {
	defer wg.Done()
	var result []string

	for _, val := range c.Engineers {
		current := fmt.Sprintf("%v %v %v", val.Name, val.RollNo, val.Marks)
		result = append(result, current)
	}

	err := os.WriteFile("db.txt", []byte(strings.Join(result, "\n")), 0664)
	if err != nil {
		fmt.Println(err.Error())
	}
}