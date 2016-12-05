package main

import "regexp"

type Instruction struct {
	isCInstruction bool
	address string
	destination string
	expression string
	jump string
}

func firstPass(instruction string, instr []Instruction, pc uint16, st SymbolTable) (i []Instruction, p uint16 , s SymbolTable) {
	if(instruction == "") {
		return 
	}
	pattern := "\\((.*)\\)|@(.*)|((.*)=)?([^;]+)(;(.*))?"
	regex, _ := regexp.Compile(pattern)
	matches := regex.FindAllStringSubmatch(instruction,1);
	match := matches[0]
	if(match[1] != "") {
		// this is a label 
		st.insertConstant(match[1], pc)
		return instr, pc, st;
	} else if (match[2] != "") {
		// this is an A instruction
		instr = append(instr, Instruction{false, match[2], "", "", ""});
		return instr, pc+1, st;
	}
	// this is an C instruction
	instr = append(instr, Instruction{true, "",  match[4],  match[5],  match[7]});
	return instr, pc+1, st;
}


