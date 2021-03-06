// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package locales

import "golang.org/x/text/message/catalog"

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}
func init() {
	locales = map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"fr": &dictionary{index: frIndex, data: frData},
	}
}

var messageKeyToIndex = map[string]int{
	"eras.wide.0":                     0,
	"eras.wide.1":                     1,
	"eras.narrow.0":                   2,
	"eras.narrow.1":                   3,
	"eras.abbr.0":                     4,
	"eras.abbr.1":                     5,
	"quarters.abbr_sa.1":              6,
	"quarters.abbr_sa.2":              7,
	"quarters.abbr_sa.3":              8,
	"quarters.abbr_sa.4":              9,
	"quarters.wide.1":                 10,
	"quarters.wide.2":                 11,
	"quarters.wide.3":                 12,
	"quarters.wide.4":                 13,
	"quarters.narrow.1":               14,
	"quarters.narrow.2":               15,
	"quarters.narrow.3":               16,
	"quarters.narrow.4":               17,
	"quarters.abbr.1":                 18,
	"quarters.abbr.2":                 19,
	"quarters.abbr.3":                 20,
	"quarters.abbr.4":                 21,
	"quarters.wide_sa.1":              22,
	"quarters.wide_sa.2":              23,
	"quarters.wide_sa.3":              24,
	"quarters.wide_sa.4":              25,
	"quarters.narrow_sa.1":            26,
	"quarters.narrow_sa.2":            27,
	"quarters.narrow_sa.3":            28,
	"quarters.narrow_sa.4":            29,
	"months.abbr.1":                   30,
	"months.abbr.2":                   31,
	"months.abbr.3":                   32,
	"months.abbr.4":                   33,
	"months.abbr.5":                   34,
	"months.abbr.6":                   35,
	"months.abbr.7":                   36,
	"months.abbr.8":                   37,
	"months.abbr.9":                   38,
	"months.abbr.10":                  39,
	"months.abbr.11":                  40,
	"months.abbr.12":                  41,
	"months.wide_sa.1":                42,
	"months.wide_sa.2":                43,
	"months.wide_sa.3":                44,
	"months.wide_sa.4":                45,
	"months.wide_sa.5":                46,
	"months.wide_sa.6":                47,
	"months.wide_sa.7":                48,
	"months.wide_sa.8":                49,
	"months.wide_sa.9":                50,
	"months.wide_sa.10":               51,
	"months.wide_sa.11":               52,
	"months.wide_sa.12":               53,
	"months.narrow_sa.1":              54,
	"months.narrow_sa.2":              55,
	"months.narrow_sa.3":              56,
	"months.narrow_sa.4":              57,
	"months.narrow_sa.5":              58,
	"months.narrow_sa.6":              59,
	"months.narrow_sa.7":              60,
	"months.narrow_sa.8":              61,
	"months.narrow_sa.9":              62,
	"months.narrow_sa.10":             63,
	"months.narrow_sa.11":             64,
	"months.narrow_sa.12":             65,
	"months.abbr_sa.1":                66,
	"months.abbr_sa.2":                67,
	"months.abbr_sa.3":                68,
	"months.abbr_sa.4":                69,
	"months.abbr_sa.5":                70,
	"months.abbr_sa.6":                71,
	"months.abbr_sa.7":                72,
	"months.abbr_sa.8":                73,
	"months.abbr_sa.9":                74,
	"months.abbr_sa.10":               75,
	"months.abbr_sa.11":               76,
	"months.abbr_sa.12":               77,
	"months.wide.1":                   78,
	"months.wide.2":                   79,
	"months.wide.3":                   80,
	"months.wide.4":                   81,
	"months.wide.5":                   82,
	"months.wide.6":                   83,
	"months.wide.7":                   84,
	"months.wide.8":                   85,
	"months.wide.9":                   86,
	"months.wide.10":                  87,
	"months.wide.11":                  88,
	"months.wide.12":                  89,
	"months.narrow.1":                 90,
	"months.narrow.2":                 91,
	"months.narrow.3":                 92,
	"months.narrow.4":                 93,
	"months.narrow.5":                 94,
	"months.narrow.6":                 95,
	"months.narrow.7":                 96,
	"months.narrow.8":                 97,
	"months.narrow.9":                 98,
	"months.narrow.10":                99,
	"months.narrow.11":                100,
	"months.narrow.12":                101,
	"days.short_sa.sun":               102,
	"days.short_sa.mon":               103,
	"days.short_sa.tue":               104,
	"days.short_sa.wed":               105,
	"days.short_sa.thu":               106,
	"days.short_sa.fri":               107,
	"days.short_sa.sat":               108,
	"days.wide.sun":                   109,
	"days.wide.mon":                   110,
	"days.wide.tue":                   111,
	"days.wide.wed":                   112,
	"days.wide.thu":                   113,
	"days.wide.fri":                   114,
	"days.wide.sat":                   115,
	"days.narrow.sun":                 116,
	"days.narrow.mon":                 117,
	"days.narrow.tue":                 118,
	"days.narrow.wed":                 119,
	"days.narrow.thu":                 120,
	"days.narrow.fri":                 121,
	"days.narrow.sat":                 122,
	"days.abbr.sun":                   123,
	"days.abbr.mon":                   124,
	"days.abbr.tue":                   125,
	"days.abbr.wed":                   126,
	"days.abbr.thu":                   127,
	"days.abbr.fri":                   128,
	"days.abbr.sat":                   129,
	"days.short.sun":                  130,
	"days.short.mon":                  131,
	"days.short.tue":                  132,
	"days.short.wed":                  133,
	"days.short.thu":                  134,
	"days.short.fri":                  135,
	"days.short.sat":                  136,
	"days.wide_sa.sun":                137,
	"days.wide_sa.mon":                138,
	"days.wide_sa.tue":                139,
	"days.wide_sa.wed":                140,
	"days.wide_sa.thu":                141,
	"days.wide_sa.fri":                142,
	"days.wide_sa.sat":                143,
	"days.narrow_sa.sun":              144,
	"days.narrow_sa.mon":              145,
	"days.narrow_sa.tue":              146,
	"days.narrow_sa.wed":              147,
	"days.narrow_sa.thu":              148,
	"days.narrow_sa.fri":              149,
	"days.narrow_sa.sat":              150,
	"days.abbr_sa.sun":                151,
	"days.abbr_sa.mon":                152,
	"days.abbr_sa.tue":                153,
	"days.abbr_sa.wed":                154,
	"days.abbr_sa.thu":                155,
	"days.abbr_sa.fri":                156,
	"days.abbr_sa.sat":                157,
	"dayPeriods.wide.am":              158,
	"dayPeriods.wide.pm":              159,
	"dayPeriods.wide.morning1":        160,
	"dayPeriods.wide.noon":            161,
	"dayPeriods.wide.afternoon1":      162,
	"dayPeriods.wide.evening1":        163,
	"dayPeriods.wide.midnight":        164,
	"dayPeriods.wide.night1":          165,
	"dayPeriods.narrow.am":            166,
	"dayPeriods.narrow.pm":            167,
	"dayPeriods.narrow.morning1":      168,
	"dayPeriods.narrow.noon":          169,
	"dayPeriods.narrow.afternoon1":    170,
	"dayPeriods.narrow.evening1":      171,
	"dayPeriods.narrow.midnight":      172,
	"dayPeriods.narrow.night1":        173,
	"dayPeriods.abbr.am":              174,
	"dayPeriods.abbr.pm":              175,
	"dayPeriods.abbr.morning1":        176,
	"dayPeriods.abbr.noon":            177,
	"dayPeriods.abbr.afternoon1":      178,
	"dayPeriods.abbr.evening1":        179,
	"dayPeriods.abbr.midnight":        180,
	"dayPeriods.abbr.night1":          181,
	"dayPeriods.wide_sa.am":           182,
	"dayPeriods.wide_sa.pm":           183,
	"dayPeriods.wide_sa.morning1":     184,
	"dayPeriods.wide_sa.noon":         185,
	"dayPeriods.wide_sa.afternoon1":   186,
	"dayPeriods.wide_sa.evening1":     187,
	"dayPeriods.wide_sa.midnight":     188,
	"dayPeriods.wide_sa.night1":       189,
	"dayPeriods.narrow_sa.am":         190,
	"dayPeriods.narrow_sa.pm":         191,
	"dayPeriods.narrow_sa.morning1":   192,
	"dayPeriods.narrow_sa.noon":       193,
	"dayPeriods.narrow_sa.afternoon1": 194,
	"dayPeriods.narrow_sa.evening1":   195,
	"dayPeriods.narrow_sa.midnight":   196,
	"dayPeriods.narrow_sa.night1":     197,
	"dayPeriods.abbr_sa.am":           198,
	"dayPeriods.abbr_sa.pm":           199,
	"dayPeriods.abbr_sa.morning1":     200,
	"dayPeriods.abbr_sa.noon":         201,
	"dayPeriods.abbr_sa.afternoon1":   202,
	"dayPeriods.abbr_sa.evening1":     203,
	"dayPeriods.abbr_sa.midnight":     204,
	"dayPeriods.abbr_sa.night1":       205,
}

var enIndex = []uint32{ // 207 elements
	// Entry 0 - 1F
	0x00000000, 0x0000000d, 0x00000018, 0x00000019,
	0x0000001a, 0x0000001c, 0x0000001e, 0x00000020,
	0x00000022, 0x00000024, 0x00000026, 0x00000031,
	0x0000003c, 0x00000047, 0x00000052, 0x00000053,
	0x00000054, 0x00000055, 0x00000056, 0x00000058,
	0x0000005a, 0x0000005c, 0x0000005e, 0x00000069,
	0x00000074, 0x0000007f, 0x0000008a, 0x0000008b,
	0x0000008c, 0x0000008d, 0x0000008e, 0x00000091,
	// Entry 20 - 3F
	0x00000094, 0x00000097, 0x0000009a, 0x0000009d,
	0x000000a0, 0x000000a3, 0x000000a6, 0x000000a9,
	0x000000ac, 0x000000af, 0x000000b2, 0x000000b9,
	0x000000c1, 0x000000c6, 0x000000cb, 0x000000ce,
	0x000000d2, 0x000000d6, 0x000000dc, 0x000000e5,
	0x000000ec, 0x000000f4, 0x000000fc, 0x000000fd,
	0x000000fe, 0x000000ff, 0x00000100, 0x00000101,
	0x00000102, 0x00000103, 0x00000104, 0x00000105,
	// Entry 40 - 5F
	0x00000106, 0x00000107, 0x00000108, 0x0000010b,
	0x0000010e, 0x00000111, 0x00000114, 0x00000117,
	0x0000011a, 0x0000011d, 0x00000120, 0x00000123,
	0x00000126, 0x00000129, 0x0000012c, 0x00000133,
	0x0000013b, 0x00000140, 0x00000145, 0x00000148,
	0x0000014c, 0x00000150, 0x00000156, 0x0000015f,
	0x00000166, 0x0000016e, 0x00000176, 0x00000177,
	0x00000178, 0x00000179, 0x0000017a, 0x0000017b,
	// Entry 60 - 7F
	0x0000017c, 0x0000017d, 0x0000017e, 0x0000017f,
	0x00000180, 0x00000181, 0x00000182, 0x00000184,
	0x00000186, 0x00000188, 0x0000018a, 0x0000018c,
	0x0000018e, 0x00000190, 0x00000196, 0x0000019c,
	0x000001a3, 0x000001ac, 0x000001b4, 0x000001ba,
	0x000001c2, 0x000001c3, 0x000001c4, 0x000001c5,
	0x000001c6, 0x000001c7, 0x000001c8, 0x000001c9,
	0x000001cc, 0x000001cf, 0x000001d2, 0x000001d5,
	// Entry 80 - 9F
	0x000001d8, 0x000001db, 0x000001de, 0x000001e0,
	0x000001e2, 0x000001e4, 0x000001e6, 0x000001e8,
	0x000001ea, 0x000001ec, 0x000001f2, 0x000001f8,
	0x000001ff, 0x00000208, 0x00000210, 0x00000216,
	0x0000021e, 0x0000021f, 0x00000220, 0x00000221,
	0x00000222, 0x00000223, 0x00000224, 0x00000225,
	0x00000228, 0x0000022b, 0x0000022e, 0x00000231,
	0x00000234, 0x00000237, 0x0000023a, 0x0000023c,
	// Entry A0 - BF
	0x0000023e, 0x0000024c, 0x00000250, 0x00000260,
	0x0000026e, 0x00000276, 0x0000027e, 0x0000027f,
	0x00000280, 0x0000028e, 0x0000028f, 0x0000029f,
	0x000002ad, 0x000002af, 0x000002b7, 0x000002b9,
	0x000002bb, 0x000002c9, 0x000002cd, 0x000002dd,
	0x000002eb, 0x000002f3, 0x000002fb, 0x000002fd,
	0x000002ff, 0x00000306, 0x0000030a, 0x00000313,
	0x0000031a, 0x00000322, 0x00000327, 0x00000329,
	// Entry C0 - DF
	0x0000032b, 0x00000332, 0x00000336, 0x0000033f,
	0x00000346, 0x0000034e, 0x00000353, 0x00000355,
	0x00000357, 0x0000035e, 0x00000362, 0x0000036b,
	0x00000372, 0x0000037a, 0x0000037f,
} // Size: 852 bytes

const enData string = "" + // Size: 895 bytes
	"Before ChristAnno DominiBABCADQ1Q2Q3Q41st quarter2nd quarter3rd quarter4" +
	"th quarter1234Q1Q2Q3Q41st quarter2nd quarter3rd quarter4th quarter1234Ja" +
	"nFebMarAprMayJunJulAugSepOctNovDecJanuaryFebruaryMarchAprilMayJuneJulyAu" +
	"gustSeptemberOctoberNovemberDecemberJFMAMJJASONDJanFebMarAprMayJunJulAug" +
	"SepOctNovDecJanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNo" +
	"vemberDecemberJFMAMJJASONDSuMoTuWeThFrSaSundayMondayTuesdayWednesdayThur" +
	"sdayFridaySaturdaySMTWTFSSunMonTueWedThuFriSatSuMoTuWeThFrSaSundayMonday" +
	"TuesdayWednesdayThursdayFridaySaturdaySMTWTFSSunMonTueWedThuFriSatAMPMin" +
	" the morningnoonin the afternoonin the eveningmidnightat nightapin the m" +
	"orningnin the afternoonin the eveningmiat nightAMPMin the morningnoonin " +
	"the afternoonin the eveningmidnightat nightAMPMmorningnoonafternooneveni" +
	"ngmidnightnightAMPMmorningnoonafternooneveningmidnightnightAMPMmorningno" +
	"onafternooneveningmidnightnight"

var frIndex = []uint32{ // 207 elements
	// Entry 0 - 1F
	0x00000000, 0x00000013, 0x00000027, 0x00000030,
	0x00000039, 0x00000042, 0x0000004b, 0x0000004d,
	0x0000004f, 0x00000051, 0x00000053, 0x00000060,
	0x0000006c, 0x00000078, 0x00000084, 0x00000085,
	0x00000086, 0x00000087, 0x00000088, 0x0000008a,
	0x0000008c, 0x0000008e, 0x00000090, 0x0000009d,
	0x000000a9, 0x000000b5, 0x000000c1, 0x000000c2,
	0x000000c3, 0x000000c4, 0x000000c5, 0x000000ca,
	// Entry 20 - 3F
	0x000000d0, 0x000000d4, 0x000000d8, 0x000000db,
	0x000000df, 0x000000e4, 0x000000e9, 0x000000ee,
	0x000000f2, 0x000000f6, 0x000000fb, 0x00000102,
	0x0000010a, 0x0000010e, 0x00000113, 0x00000116,
	0x0000011a, 0x00000121, 0x00000126, 0x0000012f,
	0x00000136, 0x0000013e, 0x00000147, 0x00000148,
	0x00000149, 0x0000014a, 0x0000014b, 0x0000014c,
	0x0000014d, 0x0000014e, 0x0000014f, 0x00000150,
	// Entry 40 - 5F
	0x00000151, 0x00000152, 0x00000153, 0x00000158,
	0x0000015e, 0x00000162, 0x00000166, 0x00000169,
	0x0000016d, 0x00000172, 0x00000177, 0x0000017c,
	0x00000180, 0x00000184, 0x00000189, 0x00000190,
	0x00000198, 0x0000019c, 0x000001a1, 0x000001a4,
	0x000001a8, 0x000001af, 0x000001b4, 0x000001bd,
	0x000001c4, 0x000001cc, 0x000001d5, 0x000001d6,
	0x000001d7, 0x000001d8, 0x000001d9, 0x000001da,
	// Entry 60 - 7F
	0x000001db, 0x000001dc, 0x000001dd, 0x000001de,
	0x000001df, 0x000001e0, 0x000001e1, 0x000001e3,
	0x000001e5, 0x000001e7, 0x000001e9, 0x000001eb,
	0x000001ed, 0x000001ef, 0x000001f7, 0x000001fc,
	0x00000201, 0x00000209, 0x0000020e, 0x00000216,
	0x0000021c, 0x0000021d, 0x0000021e, 0x0000021f,
	0x00000220, 0x00000221, 0x00000222, 0x00000223,
	0x00000227, 0x0000022b, 0x0000022f, 0x00000233,
	// Entry 80 - 9F
	0x00000237, 0x0000023b, 0x0000023f, 0x00000241,
	0x00000243, 0x00000245, 0x00000247, 0x00000249,
	0x0000024b, 0x0000024d, 0x00000255, 0x0000025a,
	0x0000025f, 0x00000267, 0x0000026c, 0x00000274,
	0x0000027a, 0x0000027b, 0x0000027c, 0x0000027d,
	0x0000027e, 0x0000027f, 0x00000280, 0x00000281,
	0x00000285, 0x00000289, 0x0000028d, 0x00000291,
	0x00000295, 0x00000299, 0x0000029d, 0x0000029f,
	// Entry A0 - BF
	0x000002a1, 0x000002a9, 0x000002ad, 0x000002bf,
	0x000002c6, 0x000002cc, 0x000002d4, 0x000002d6,
	0x000002d8, 0x000002dc, 0x000002e0, 0x000002e5,
	0x000002e9, 0x000002ef, 0x000002f3, 0x000002f5,
	0x000002f7, 0x000002fb, 0x000002ff, 0x00000304,
	0x00000308, 0x0000030e, 0x00000312, 0x00000314,
	0x00000316, 0x0000031b, 0x0000031f, 0x0000032a,
	0x0000032e, 0x00000334, 0x00000338, 0x0000033a,
	// Entry C0 - DF
	0x0000033c, 0x00000340, 0x00000344, 0x00000349,
	0x0000034d, 0x00000353, 0x00000357, 0x00000359,
	0x0000035b, 0x0000035f, 0x00000363, 0x00000368,
	0x0000036c, 0x00000372, 0x00000376,
} // Size: 852 bytes

const frData string = "" + // Size: 886 bytes
	"avant Jésus-Christaprès Jésus-Christav. J.-C.ap. J.-C.av. J.-C.ap. J.-C." +
	"T1T2T3T41er trimestre2e trimestre3e trimestre4e trimestre1234T1T2T3T41er" +
	" trimestre2e trimestre3e trimestre4e trimestre1234janv.févr.marsavr.maij" +
	"uinjuil.aoûtsept.oct.nov.déc.janvierfévriermarsavrilmaijuinjuilletaoûtse" +
	"ptembreoctobrenovembredécembreJFMAMJJASONDjanv.févr.marsavr.maijuinjuil." +
	"aoûtsept.oct.nov.déc.janvierfévriermarsavrilmaijuinjuilletaoûtseptembreo" +
	"ctobrenovembredécembreJFMAMJJASONDdilumamejevesadimanchelundimardimercre" +
	"dijeudivendredisamediDLMMJVSdim.lun.mar.mer.jeu.ven.sam.dilumamejevesadi" +
	"manchelundimardimercredijeudivendredisamediDLMMJVSdim.lun.mar.mer.jeu.ve" +
	"n.sam.AMPMdu matinmidide l’après-mididu soirminuitdu matinAMPMmat.midiap" +
	".m.soirminuitnuitAMPMmat.midiap.m.soirminuitnuitAMPMmatinmidiaprès-midis" +
	"oirminuitnuitAMPMmat.midiap.m.soirminuitnuitAMPMmat.midiap.m.soirminuitn" +
	"uit"

	// Total table size 3485 bytes (3KiB); checksum: ADF7BBED
