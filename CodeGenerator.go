package main

import (
	"strconv"
)


func secondPass(instr Instruction, syscodes []uint16,  st SymbolTable) []uint16 {
	if(instr.isCInstruction) {
		return append(syscodes, generateCInstruction(instr , st))
	} 
	return append(syscodes, generateAInstruction(instr , st))
}

func generateCInstruction(instr Instruction ,  st SymbolTable) uint16 {
	return uint16(0xE000) | getDestination(instr.destination, st) | getValue(instr.expression, st) | getValue(instr.jump, st)
}

func getDestination(dest string, st SymbolTable) uint16 {
	ret := uint16(0)
	for _, runeValue := range dest {
        ret = ret | getDest(runeValue)
    }
    return ret
}

func getDest(r rune) uint16 {
	switch {
	case r == 'M':
		return 0x8
	case r == 'D':
		return 0x10
	case r == 'A':
		return 0x20
	}
	return 0
}

func getValue(dest string, st SymbolTable) uint16 {
	symbolval := st.getOpeartiion(dest)
	if symbolval != nil {
		return symbolval.(uint16)
	}
	return uint16(0)
}

func generateAInstruction(instr Instruction,  st SymbolTable) uint16 {
	symbolval := st.getConstant(instr.address)
	if symbolval != nil {
		return symbolval.(uint16)
	}
	i,e := strconv.Atoi(instr.address)
	if(e != nil) {
		rCounter++
		st.insertConstant(instr.address, rCounter)
		return rCounter
	}
	return uint16(i);
}
