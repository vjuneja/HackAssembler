package main

import "github.com/tchap/go-patricia/patricia"

type SymbolTable struct {
	constants *patricia.Trie
	opeartiions *patricia.Trie

}

func newSymbolTable() SymbolTable {
	return SymbolTable{patricia.NewTrie(), patricia.NewTrie()};
}

func (t SymbolTable) initialize() SymbolTable {
	t.constants.Insert(patricia.Prefix("R0"),uint16(0));
	t.constants.Insert(patricia.Prefix("R1"),uint16(1));
	t.constants.Insert(patricia.Prefix("R2"),uint16(2));
	t.constants.Insert(patricia.Prefix("R3"),uint16(3));
	t.constants.Insert(patricia.Prefix("R4"),uint16(4));
	t.constants.Insert(patricia.Prefix("R5"),uint16(5));
	t.constants.Insert(patricia.Prefix("R6"),uint16(6));
	t.constants.Insert(patricia.Prefix("R7"),uint16(7));
	t.constants.Insert(patricia.Prefix("R8"),uint16(8));
	t.constants.Insert(patricia.Prefix("R9"),uint16(9));
	t.constants.Insert(patricia.Prefix("R10"),uint16(10));
	t.constants.Insert(patricia.Prefix("R11"),uint16(11));
	t.constants.Insert(patricia.Prefix("R12"),uint16(12));
	t.constants.Insert(patricia.Prefix("R13"),uint16(13));
	t.constants.Insert(patricia.Prefix("R14"),uint16(14));
	t.constants.Insert(patricia.Prefix("R15"),uint16(15));

	t.constants.Insert(patricia.Prefix("SP"),uint16(0));
	t.constants.Insert(patricia.Prefix("LCL"),uint16(1));
	t.constants.Insert(patricia.Prefix("ARG"),uint16(2));
	t.constants.Insert(patricia.Prefix("THIS"),uint16(3));
	t.constants.Insert(patricia.Prefix("THAT"),uint16(4));
	t.constants.Insert(patricia.Prefix("SCREEN"),uint16(16384));
	t.constants.Insert(patricia.Prefix("KBD"),uint16(24576));

	t.opeartiions.Insert(patricia.Prefix("0"),uint16(0xA80));
	t.opeartiions.Insert(patricia.Prefix("1"),uint16(0xFC0));
	t.opeartiions.Insert(patricia.Prefix("-1"),uint16(0xE80));
	t.opeartiions.Insert(patricia.Prefix("D"),uint16(0x300));
	t.opeartiions.Insert(patricia.Prefix("A"),uint16(0xC00));
	t.opeartiions.Insert(patricia.Prefix("!D"),uint16(0x340));
	t.opeartiions.Insert(patricia.Prefix("!A"),uint16(0xC40));
	t.opeartiions.Insert(patricia.Prefix("-D"),uint16(0x3C0));
	t.opeartiions.Insert(patricia.Prefix("-A"),uint16(0xCC0));
	t.opeartiions.Insert(patricia.Prefix("D+1"),uint16(0x7C0));
	t.opeartiions.Insert(patricia.Prefix("A+1"),uint16(0xDC0));
	t.opeartiions.Insert(patricia.Prefix("D-1"),uint16(0x380));
	t.opeartiions.Insert(patricia.Prefix("A-1"),uint16(0xC80));
	t.opeartiions.Insert(patricia.Prefix("D+A"),uint16(0x80));
	t.opeartiions.Insert(patricia.Prefix("D-A"),uint16(0x4C0));
	t.opeartiions.Insert(patricia.Prefix("A-D"),uint16(0x1C0));
	t.opeartiions.Insert(patricia.Prefix("D&A"),uint16(0x0));
	t.opeartiions.Insert(patricia.Prefix("D|A"),uint16(0x540));
	t.opeartiions.Insert(patricia.Prefix("M"),uint16(0x1C00));
	t.opeartiions.Insert(patricia.Prefix("!M"),uint16(0x1C40));
	t.opeartiions.Insert(patricia.Prefix("-M"),uint16(0x1CC0));
	t.opeartiions.Insert(patricia.Prefix("M+1"),uint16(0x1DC0));
	t.opeartiions.Insert(patricia.Prefix("M-1"),uint16(0x1C80));
	t.opeartiions.Insert(patricia.Prefix("D+M"),uint16(0x1080));
	t.opeartiions.Insert(patricia.Prefix("D-M"),uint16(0x14C0));
	t.opeartiions.Insert(patricia.Prefix("M-D"),uint16(0x11C0));
	t.opeartiions.Insert(patricia.Prefix("D&M"),uint16(0x1000));
	t.opeartiions.Insert(patricia.Prefix("D|M"),uint16(0x1540));

	t.opeartiions.Insert(patricia.Prefix("JGT"),uint16(1));
	t.opeartiions.Insert(patricia.Prefix("JEQ"),uint16(2));
	t.opeartiions.Insert(patricia.Prefix("JGE"),uint16(3));
	t.opeartiions.Insert(patricia.Prefix("JLT"),uint16(4));
	t.opeartiions.Insert(patricia.Prefix("JNE"),uint16(5));
	t.opeartiions.Insert(patricia.Prefix("JLE"),uint16(6));
	t.opeartiions.Insert(patricia.Prefix("JMP"),uint16(7));


	return t;
}

func (t SymbolTable) insertConstant(symbol string, value uint16) SymbolTable {
	t.constants.Insert(patricia.Prefix(symbol),value);
	return t;
}

func (t SymbolTable) getConstant(symbol string) patricia.Item {
	return t.constants.Get(patricia.Prefix(symbol));
}

func (t SymbolTable) getOpeartiion(symbol string) patricia.Item {
	return t.opeartiions.Get(patricia.Prefix(symbol));
}
