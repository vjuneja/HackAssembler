package main

import (
	"os"
	"bufio"
	"log"
	"encoding/binary"
	"strings"
	"strconv"
	//"fmt"
)
var rCounter uint16

func main() {
    args := os.Args[1:]
    filename := args[0];
	rCounter = 15

	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	st:=newSymbolTable().initialize()

	//fmt.Println("table ", st);

    scanner := bufio.NewScanner(file)
 	pc := uint16(0)


    var instr []Instruction

	instr, pc, st = firstPass(scrub(("(OUTPUT_FIRST)")), instr, pc, st)

	instr, pc, st = firstPass(scrub(("(OUTPUT_FIRST1)")), instr, pc, st)

	//fmt.Println("table ", st);
	
	for scanner.Scan() {
		instrstr := scanner.Text()
		//fmt.Println(instrstr)

		scrubbedinstr := scrub(instrstr)
		//fmt.Println("table ", st);

    	instr, pc, _ = firstPass(scrubbedinstr, instr, pc, st)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }


	var inst16 []uint16

	for i:=0; i<len(instr); i++ {
		inst16 = secondPass(instr[i], inst16, st)
	}

	//fmt.Println(inst16);

	outputFileName := strings.Split(filename,".")[0]+"VJ.hack"

	f, _ := os.Create(outputFileName)
	defer f.Close()

	for i := range inst16 {
		binary.Write(f, binary.LittleEndian, []byte(toBinary(inst16[i])))
	}

}

func toBinary(instr uint16) string{
	s:=""
	for i:=0; i<16;i++ {
		s = strconv.Itoa(int(instr%2)) + s;
		instr=instr/2;
	}
	s+="\n"
	return s
	

}