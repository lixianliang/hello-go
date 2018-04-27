# command-line-arguments
"".main STEXT size=471 args=0x0 locals=0xb0
	0x0000 00000 (/home/letian/go/src/go-example/base/append.go:5)	TEXT	"".main(SB), $176-0
	0x0000 00000 (/home/letian/go/src/go-example/base/append.go:5)	MOVQ	(TLS), CX
	0x0009 00009 (/home/letian/go/src/go-example/base/append.go:5)	LEAQ	-48(SP), AX
	0x000e 00014 (/home/letian/go/src/go-example/base/append.go:5)	CMPQ	AX, 16(CX)
	0x0012 00018 (/home/letian/go/src/go-example/base/append.go:5)	JLS	461
	0x0018 00024 (/home/letian/go/src/go-example/base/append.go:5)	SUBQ	$176, SP
	0x001f 00031 (/home/letian/go/src/go-example/base/append.go:5)	MOVQ	BP, 168(SP)
	0x0027 00039 (/home/letian/go/src/go-example/base/append.go:5)	LEAQ	168(SP), BP
	0x002f 00047 (/home/letian/go/src/go-example/base/append.go:5)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x002f 00047 (/home/letian/go/src/go-example/base/append.go:5)	FUNCDATA	$1, gclocals·faa80ab9e72c49377c7ed5d021e03ce8(SB)
	0x002f 00047 (/home/letian/go/src/go-example/base/append.go:5)	LEAQ	""..autotmp_5+88(SP), DI
	0x0034 00052 (/home/letian/go/src/go-example/base/append.go:5)	LEAQ	"".statictmp_0(SB), SI
	0x003b 00059 (/home/letian/go/src/go-example/base/append.go:6)	DUFFCOPY	$868
	0x004e 00078 (/home/letian/go/src/go-example/base/append.go:6)	LEAQ	type.int(SB), AX
	0x0055 00085 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	AX, (SP)
	0x0059 00089 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	$1, 8(SP)
	0x0062 00098 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	$5, 16(SP)
	0x006b 00107 (/home/letian/go/src/go-example/base/append.go:7)	PCDATA	$0, $0
	0x006b 00107 (/home/letian/go/src/go-example/base/append.go:7)	CALL	runtime.makeslice(SB)
	0x0070 00112 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	24(SP), AX
	0x0075 00117 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	40(SP), CX
	0x007a 00122 (/home/letian/go/src/go-example/base/append.go:7)	MOVQ	32(SP), DX
	0x007f 00127 (/home/letian/go/src/go-example/base/append.go:8)	TESTQ	DX, DX
	0x0082 00130 (/home/letian/go/src/go-example/base/append.go:8)	JLS	454
	0x0088 00136 (/home/letian/go/src/go-example/base/append.go:8)	MOVQ	$5, (AX)
	0x008f 00143 (/home/letian/go/src/go-example/base/append.go:10)	LEAQ	4(DX), BX
	0x0093 00147 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	BX, "".s1.len+64(SP)
	0x0098 00152 (/home/letian/go/src/go-example/base/append.go:10)	CMPQ	BX, CX
	0x009b 00155 (/home/letian/go/src/go-example/base/append.go:10)	JHI	381
	0x00a1 00161 (/home/letian/go/src/go-example/base/append.go:10)	CMPQ	BX, CX
	0x00a4 00164 (/home/letian/go/src/go-example/base/append.go:10)	JHI	447
	0x00aa 00170 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	CX, "".s1.cap+72(SP)
	0x00af 00175 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	AX, ""..autotmp_18+120(SP)
	0x00b4 00180 (/home/letian/go/src/go-example/base/append.go:10)	LEAQ	(AX)(DX*8), CX
	0x00b8 00184 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	CX, (SP)
	0x00bc 00188 (/home/letian/go/src/go-example/base/append.go:10)	LEAQ	""..autotmp_5+88(SP), CX
	0x00c1 00193 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	CX, 8(SP)
	0x00c6 00198 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	$32, 16(SP)
	0x00cf 00207 (/home/letian/go/src/go-example/base/append.go:10)	PCDATA	$0, $1
	0x00cf 00207 (/home/letian/go/src/go-example/base/append.go:10)	CALL	runtime.memmove(SB)
	0x00d4 00212 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	""..autotmp_18+120(SP), AX
	0x00d9 00217 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, ""..autotmp_3+144(SP)
	0x00e1 00225 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	"".s1.len+64(SP), AX
	0x00e6 00230 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, ""..autotmp_3+152(SP)
	0x00ee 00238 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	"".s1.cap+72(SP), AX
	0x00f3 00243 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, ""..autotmp_3+160(SP)
	0x00fb 00251 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	$0, ""..autotmp_2+128(SP)
	0x0107 00263 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	$0, ""..autotmp_2+136(SP)
	0x0113 00275 (/home/letian/go/src/go-example/base/append.go:11)	LEAQ	type.[]int(SB), AX
	0x011a 00282 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, (SP)
	0x011e 00286 (/home/letian/go/src/go-example/base/append.go:11)	LEAQ	""..autotmp_3+144(SP), AX
	0x0126 00294 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, 8(SP)
	0x012b 00299 (/home/letian/go/src/go-example/base/append.go:11)	PCDATA	$0, $2
	0x012b 00299 (/home/letian/go/src/go-example/base/append.go:11)	CALL	runtime.convT2Eslice(SB)
	0x0130 00304 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	16(SP), AX
	0x0135 00309 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	24(SP), CX
	0x013a 00314 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, ""..autotmp_2+128(SP)
	0x0142 00322 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	CX, ""..autotmp_2+136(SP)
	0x014a 00330 (/home/letian/go/src/go-example/base/append.go:11)	LEAQ	""..autotmp_2+128(SP), AX
	0x0152 00338 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	AX, (SP)
	0x0156 00342 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	$1, 8(SP)
	0x015f 00351 (/home/letian/go/src/go-example/base/append.go:11)	MOVQ	$1, 16(SP)
	0x0168 00360 (/home/letian/go/src/go-example/base/append.go:11)	PCDATA	$0, $2
	0x0168 00360 (/home/letian/go/src/go-example/base/append.go:11)	CALL	fmt.Println(SB)
	0x016d 00365 (/home/letian/go/src/go-example/base/append.go:12)	MOVQ	168(SP), BP
	0x0175 00373 (/home/letian/go/src/go-example/base/append.go:12)	ADDQ	$176, SP
	0x017c 00380 (/home/letian/go/src/go-example/base/append.go:12)	RET
	0x017d 00381 (/home/letian/go/src/go-example/base/append.go:12)	MOVQ	DX, ""..autotmp_19+80(SP)
	0x0182 00386 (/home/letian/go/src/go-example/base/append.go:12)	LEAQ	type.int(SB), SI
	0x0189 00393 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	SI, (SP)
	0x018d 00397 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	AX, 8(SP)
	0x0192 00402 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	DX, 16(SP)
	0x0197 00407 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	CX, 24(SP)
	0x019c 00412 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	BX, 32(SP)
	0x01a1 00417 (/home/letian/go/src/go-example/base/append.go:10)	PCDATA	$0, $0
	0x01a1 00417 (/home/letian/go/src/go-example/base/append.go:10)	CALL	runtime.growslice(SB)
	0x01a6 00422 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	56(SP), CX
	0x01ab 00427 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	40(SP), AX
	0x01b0 00432 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	""..autotmp_19+80(SP), DX
	0x01b5 00437 (/home/letian/go/src/go-example/base/append.go:10)	MOVQ	"".s1.len+64(SP), BX
	0x01ba 00442 (/home/letian/go/src/go-example/base/append.go:10)	JMP	161
	0x01bf 00447 (/home/letian/go/src/go-example/base/append.go:10)	PCDATA	$0, $0
	0x01bf 00447 (/home/letian/go/src/go-example/base/append.go:10)	CALL	runtime.panicslice(SB)
	0x01c4 00452 (/home/letian/go/src/go-example/base/append.go:10)	UNDEF
	0x01c6 00454 (/home/letian/go/src/go-example/base/append.go:8)	PCDATA	$0, $0
	0x01c6 00454 (/home/letian/go/src/go-example/base/append.go:8)	CALL	runtime.panicindex(SB)
	0x01cb 00459 (/home/letian/go/src/go-example/base/append.go:8)	UNDEF
	0x01cd 00461 (/home/letian/go/src/go-example/base/append.go:8)	NOP
	0x01cd 00461 (/home/letian/go/src/go-example/base/append.go:5)	PCDATA	$0, $-1
	0x01cd 00461 (/home/letian/go/src/go-example/base/append.go:5)	CALL	runtime.morestack_noctxt(SB)
	0x01d2 00466 (/home/letian/go/src/go-example/base/append.go:5)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 8d 44 24 d0 48 3b  dH..%....H.D$.H;
	0x0010 41 10 0f 86 b5 01 00 00 48 81 ec b0 00 00 00 48  A.......H......H
	0x0020 89 ac 24 a8 00 00 00 48 8d ac 24 a8 00 00 00 48  ..$....H..$....H
	0x0030 8d 7c 24 58 48 8d 35 00 00 00 00 48 89 6c 24 f0  .|$XH.5....H.l$.
	0x0040 48 8d 6c 24 f0 e8 00 00 00 00 48 8b 6d 00 48 8d  H.l$......H.m.H.
	0x0050 05 00 00 00 00 48 89 04 24 48 c7 44 24 08 01 00  .....H..$H.D$...
	0x0060 00 00 48 c7 44 24 10 05 00 00 00 e8 00 00 00 00  ..H.D$..........
	0x0070 48 8b 44 24 18 48 8b 4c 24 28 48 8b 54 24 20 48  H.D$.H.L$(H.T$ H
	0x0080 85 d2 0f 86 3e 01 00 00 48 c7 00 05 00 00 00 48  ....>...H......H
	0x0090 8d 5a 04 48 89 5c 24 40 48 39 cb 0f 87 dc 00 00  .Z.H.\$@H9......
	0x00a0 00 48 39 cb 0f 87 15 01 00 00 48 89 4c 24 48 48  .H9.......H.L$HH
	0x00b0 89 44 24 78 48 8d 0c d0 48 89 0c 24 48 8d 4c 24  .D$xH...H..$H.L$
	0x00c0 58 48 89 4c 24 08 48 c7 44 24 10 20 00 00 00 e8  XH.L$.H.D$. ....
	0x00d0 00 00 00 00 48 8b 44 24 78 48 89 84 24 90 00 00  ....H.D$xH..$...
	0x00e0 00 48 8b 44 24 40 48 89 84 24 98 00 00 00 48 8b  .H.D$@H..$....H.
	0x00f0 44 24 48 48 89 84 24 a0 00 00 00 48 c7 84 24 80  D$HH..$....H..$.
	0x0100 00 00 00 00 00 00 00 48 c7 84 24 88 00 00 00 00  .......H..$.....
	0x0110 00 00 00 48 8d 05 00 00 00 00 48 89 04 24 48 8d  ...H......H..$H.
	0x0120 84 24 90 00 00 00 48 89 44 24 08 e8 00 00 00 00  .$....H.D$......
	0x0130 48 8b 44 24 10 48 8b 4c 24 18 48 89 84 24 80 00  H.D$.H.L$.H..$..
	0x0140 00 00 48 89 8c 24 88 00 00 00 48 8d 84 24 80 00  ..H..$....H..$..
	0x0150 00 00 48 89 04 24 48 c7 44 24 08 01 00 00 00 48  ..H..$H.D$.....H
	0x0160 c7 44 24 10 01 00 00 00 e8 00 00 00 00 48 8b ac  .D$..........H..
	0x0170 24 a8 00 00 00 48 81 c4 b0 00 00 00 c3 48 89 54  $....H.......H.T
	0x0180 24 50 48 8d 35 00 00 00 00 48 89 34 24 48 89 44  $PH.5....H.4$H.D
	0x0190 24 08 48 89 54 24 10 48 89 4c 24 18 48 89 5c 24  $.H.T$.H.L$.H.\$
	0x01a0 20 e8 00 00 00 00 48 8b 4c 24 38 48 8b 44 24 28   .....H.L$8H.D$(
	0x01b0 48 8b 54 24 50 48 8b 5c 24 40 e9 e2 fe ff ff e8  H.T$PH.\$@......
	0x01c0 00 00 00 00 0f 0b e8 00 00 00 00 0f 0b e8 00 00  ................
	0x01d0 00 00 e9 29 fe ff ff                             ...)...
	rel 5+4 t=16 TLS+0
	rel 55+4 t=15 "".statictmp_0+0
	rel 70+4 t=8 runtime.duffcopy+868
	rel 81+4 t=15 type.int+0
	rel 108+4 t=8 runtime.makeslice+0
	rel 208+4 t=8 runtime.memmove+0
	rel 278+4 t=15 type.[]int+0
	rel 300+4 t=8 runtime.convT2Eslice+0
	rel 361+4 t=8 fmt.Println+0
	rel 389+4 t=15 type.int+0
	rel 418+4 t=8 runtime.growslice+0
	rel 448+4 t=8 runtime.panicslice+0
	rel 455+4 t=8 runtime.panicindex+0
	rel 462+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=91 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	84
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JLS	47
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	JNE	56
	0x0031 00049 (<autogenerated>:1)	PCDATA	$0, $0
	0x0031 00049 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0036 00054 (<autogenerated>:1)	UNDEF
	0x0038 00056 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x003f 00063 (<autogenerated>:1)	PCDATA	$0, $0
	0x003f 00063 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0044 00068 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004b 00075 (<autogenerated>:1)	MOVQ	(SP), BP
	0x004f 00079 (<autogenerated>:1)	ADDQ	$8, SP
	0x0053 00083 (<autogenerated>:1)	RET
	0x0054 00084 (<autogenerated>:1)	NOP
	0x0054 00084 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0054 00084 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0059 00089 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 45 48  dH..%....H;a.vEH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 76 09 48 8b 2c 24 48 83 c4 08 c3 75  ..<.v.H.,$H....u
	0x0030 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8  ................
	0x0040 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24 48  ...........H.,$H
	0x0050 83 c4 08 c3 e8 00 00 00 00 eb a5                 ...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 50+4 t=8 runtime.throwinit+0
	rel 58+4 t=15 "".initdone·+-1
	rel 64+4 t=8 fmt.init+0
	rel 70+4 t=15 "".initdone·+-1
	rel 85+4 t=8 runtime.morestack_noctxt+0
go.info."".main SDWARFINFO size=65
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 04 73 31 2e  .............s1.
	0x0020 6c 65 6e 00 05 9c 11 88 7f 22 00 00 00 00 04 73  len......".....s
	0x0030 31 2e 63 61 70 00 05 9c 11 90 7f 22 00 00 00 00  1.cap......"....
	0x0040 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+471
	rel 42+4 t=28 go.info.int+0
	rel 60+4 t=28 go.info.int+0
go.range."".main SDWARFRANGE size=0
go.info."".init SDWARFINFO size=29
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 01 00           .............
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+91
go.range."".init SDWARFRANGE size=0
"".statictmp_0 SRODATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 02 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*[]int- SRODATA dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 00 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 48+8 t=1 type.[]int+0
type.[]int SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int-+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..hashfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	rel 0+8 t=1 runtime.memequal_varlen+0
type..alg32 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc32+0
	rel 8+8 t=1 type..eqfunc32+0
type..namedata.*[4]int- SRODATA dupok size=10
	0x0000 00 00 07 2a 5b 34 5d 69 6e 74                    ...*[4]int
type.*[4]int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f0 ba d6 e0 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[4]int-+0
	rel 48+8 t=1 type.[4]int+0
runtime.gcbits. SRODATA dupok size=0
type.[4]int SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00   ...............
	0x0010 ae 6a 57 d6 02 08 08 91 00 00 00 00 00 00 00 00  .jW.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg32+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[4]int-+0
	rel 44+4 t=6 type.*[4]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·faa80ab9e72c49377c7ed5d021e03ce8 SRODATA dupok size=11
	0x0000 03 00 00 00 06 00 00 00 00 01 0e                 ...........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
