package gen

import (
	"io"
	"strings"

	"github.com/mmcloughlin/avo/internal/inst"
)

type LoaderTest struct{}

func (l LoaderTest) Generate(w io.Writer, is []*inst.Instruction) error {
	p := &printer{w: w}

	p.printf("TEXT loadertest(SB), 0, $0\n")

	for _, i := range is {
		p.printf("\t// %s %s\n", i.Opcode, i.Summary)
		if strings.HasPrefix(i.Opcode, "RET") {
			p.printf("\t// SKIP: early RET instruction would cause assembler error")
			continue
		}

		for _, f := range i.Forms {
			as := args(f.Operands)
			p.printf("\t// %#v\n", f.Operands)
			if as == nil {
				p.printf("\t// SKIP:\n")
				continue
			}
			p.printf("\t%s\t%s\n", i.Opcode, strings.Join(as, ", "))
		}
		p.printf("\n")
	}

	p.printf("\tRET\n")

	return p.Err()
}

func args(ops []inst.Operand) []string {
	as := make([]string, len(ops))
	for i, op := range ops {
		a := arg(op.Type)
		if a == "" {
			return nil
		}
		as[i] = a
	}
	return as
}

// arg generates an argument for an operand of the given type.
func arg(t string) string {
	m := map[string]string{
		// <xs:enumeration value="1" />
		// <xs:enumeration value="3" />
		// <xs:enumeration value="imm4" />
		// <xs:enumeration value="imm8" />
		// <xs:enumeration value="imm16" />
		// <xs:enumeration value="imm32" />
		// <xs:enumeration value="imm64" />
		// <xs:enumeration value="al" />
		// <xs:enumeration value="cl" />
		// <xs:enumeration value="r8" />
		// <xs:enumeration value="r8l" />
		// <xs:enumeration value="ax" />
		// <xs:enumeration value="r16" />
		// <xs:enumeration value="r16l" />
		// <xs:enumeration value="eax" />
		// <xs:enumeration value="r32" />
		// <xs:enumeration value="r32l" />
		// <xs:enumeration value="rax" />
		"r64": "R15", // <xs:enumeration value="r64" />
		// <xs:enumeration value="mm" />
		// <xs:enumeration value="xmm0" />
		"xmm": "X7", // <xs:enumeration value="xmm" />
		// <xs:enumeration value="xmm{k}" />
		// <xs:enumeration value="xmm{k}{z}" />
		// <xs:enumeration value="ymm" />
		// <xs:enumeration value="ymm{k}" />
		// <xs:enumeration value="ymm{k}{z}" />
		// <xs:enumeration value="zmm" />
		// <xs:enumeration value="zmm{k}" />
		// <xs:enumeration value="zmm{k}{z}" />
		// <xs:enumeration value="k" />
		// <xs:enumeration value="k{k}" />
		// <xs:enumeration value="moffs32" />
		// <xs:enumeration value="moffs64" />
		// <xs:enumeration value="m" />
		// <xs:enumeration value="m8" />
		// <xs:enumeration value="m16" />
		// <xs:enumeration value="m16{k}{z}" />
		// <xs:enumeration value="m32" />
		// <xs:enumeration value="m32{k}" />
		// <xs:enumeration value="m32{k}{z}" />
		// <xs:enumeration value="m64" />
		// <xs:enumeration value="m64{k}" />
		// <xs:enumeration value="m64{k}{z}" />
		// <xs:enumeration value="m128" />
		// <xs:enumeration value="m128{k}{z}" />
		// <xs:enumeration value="m256" />
		// <xs:enumeration value="m256{k}{z}" />
		// <xs:enumeration value="m512" />
		// <xs:enumeration value="m512{k}{z}" />
		// <xs:enumeration value="m64/m32bcst" />
		// <xs:enumeration value="m128/m32bcst" />
		// <xs:enumeration value="m256/m32bcst" />
		// <xs:enumeration value="m512/m32bcst" />
		// <xs:enumeration value="m128/m64bcst" />
		// <xs:enumeration value="m256/m64bcst" />
		// <xs:enumeration value="m512/m64bcst" />
		// <xs:enumeration value="vm32x" />
		// <xs:enumeration value="vm32x{k}" />
		// <xs:enumeration value="vm64x" />
		// <xs:enumeration value="vm64x{k}" />
		// <xs:enumeration value="vm32y" />
		// <xs:enumeration value="vm32y{k}" />
		// <xs:enumeration value="vm64y" />
		// <xs:enumeration value="vm64y{k}" />
		// <xs:enumeration value="vm32z" />
		// <xs:enumeration value="vm32z{k}" />
		// <xs:enumeration value="vm64z" />
		// <xs:enumeration value="vm64z{k}" />
		// <xs:enumeration value="rel8" />
		// <xs:enumeration value="rel32" />
		// <xs:enumeration value="{er}" />
		// <xs:enumeration value="{sae}" />
	}
	return m[t]
}