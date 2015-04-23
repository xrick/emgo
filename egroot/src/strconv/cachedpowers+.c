// Powers of ten taken from https://github.com/floitsch/double-conversion

static const
uint64 significands[] = {
	0xfa8fd5a0081c0288,	// -348
	0xbaaee17fa23ebf76,	// -340
	0x8b16fb203055ac76,	// -332
	0xcf42894a5dce35ea,	// -324
	0x9a6bb0aa55653b2d,	// -316
	0xe61acf033d1a45df,	// -308
	0xab70fe17c79ac6ca,	// -300
	0xff77b1fcbebcdc4f,	// -292
	0xbe5691ef416bd60c,	// -284
	0x8dd01fad907ffc3c,	// -276
	0xd3515c2831559a83,	// -268
	0x9d71ac8fada6c9b5,	// -260
	0xea9c227723ee8bcb,	// -252
	0xaecc49914078536d,	// -244
	0x823c12795db6ce57,	// -236
	0xc21094364dfb5637,	// -228
	0x9096ea6f3848984f,	// -220
	0xd77485cb25823ac7,	// -212
	0xa086cfcd97bf97f4,	// -204
	0xef340a98172aace5,	// -196
	0xb23867fb2a35b28e,	// -188
	0x84c8d4dfd2c63f3b,	// -180
	0xc5dd44271ad3cdba,	// -172
	0x936b9fcebb25c996,	// -164
	0xdbac6c247d62a584,	// -156
	0xa3ab66580d5fdaf6,	// -148
	0xf3e2f893dec3f126,	// -140
	0xb5b5ada8aaff80b8,	// -132
	0x87625f056c7c4a8b,	// -124
	0xc9bcff6034c13053,	// -116
	0x964e858c91ba2655,	// -108
	0xdff9772470297ebd,	// -100
	0xa6dfbd9fb8e5b88f,	// -92
	0xf8a95fcf88747d94,	// -84
	0xb94470938fa89bcf,	// -76
	0x8a08f0f8bf0f156b,	// -68
	0xcdb02555653131b6,	// -60
	0x993fe2c6d07b7fac,	// -52
	0xe45c10c42a2b3b06,	// -44
	0xaa242499697392d3,	// -36
	0xfd87b5f28300ca0e,	// -28
	0xbce5086492111aeb,	// -20
	0x8cbccc096f5088cc,	// -12
	0xd1b71758e219652c,	// -4
	0x9c40000000000000,	// 4
	0xe8d4a51000000000,	// 12
	0xad78ebc5ac620000,	// 20
	0x813f3978f8940984,	// 28
	0xc097ce7bc90715b3,	// 36
	0x8f7e32ce7bea5c70,	// 44
	0xd5d238a4abe98068,	// 52
	0x9f4f2726179a2245,	// 60
	0xed63a231d4c4fb27,	// 68
	0xb0de65388cc8ada8,	// 76
	0x83c7088e1aab65db,	// 84
	0xc45d1df942711d9a,	// 92
	0x924d692ca61be758,	// 100
	0xda01ee641a708dea,	// 108
	0xa26da3999aef774a,	// 116
	0xf209787bb47d6b85,	// 124
	0xb454e4a179dd1877,	// 132
	0x865b86925b9bc5c2,	// 140
	0xc83553c5c8965d3d,	// 148
	0x952ab45cfa97a0b3,	// 156
	0xde469fbd99a05fe3,	// 164
	0xa59bc234db398c25,	// 172
	0xf6c69a72a3989f5c,	// 180
	0xb7dcbf5354e9bece,	// 188
	0x88fcf317f22241e2,	// 196
	0xcc20ce9bd35c78a5,	// 204
	0x98165af37b2153df,	// 212
	0xe2a0b5dc971f303a,	// 220
	0xa8d9d1535ce3b396,	// 228
	0xfb9b7cd9a4a7443c,	// 236
	0xbb764c4ca7a44410,	// 244
	0x8bab8eefb6409c1a,	// 252
	0xd01fef10a657842c,	// 260
	0x9b10a4e5e9913129,	// 268
	0xe7109bfba19c0c9d,	// 276
	0xac2820d9623bf429,	// 284
	0x80444b5e7aa7cf85,	// 292
	0xbf21e44003acdd2d,	// 300
	0x8e679c2f5e44ff8f,	// 308
	0xd433179d9c8cb841,	// 316
	0x9e19db92b4e31ba9,	// 324
	0xeb96bf6ebadf77d9,	// 332
	0xaf87023b9bf0ee6b	// 340
};

static const
int16 exponents[] = {
	-1220,	// -348
	-1193,	// -340
	-1166,	// -332
	-1140,	// -324
	-1113,	// -316
	-1087,	// -308
	-1060,	// -300
	-1034,	// -292
	-1007,	// -284
	-980,	// -276
	-954,	// -268
	-927,	// -260
	-901,	// -252
	-874,	// -244
	-847,	// -236
	-821,	// -228
	-794,	// -220
	-768,	// -212
	-741,	// -204
	-715,	// -196
	-688,	// -188
	-661,	// -180
	-635,	// -172
	-608,	// -164
	-582,	// -156
	-555,	// -148
	-529,	// -140
	-502,	// -132
	-475,	// -124
	-449,	// -116
	-422,	// -108
	-396,	// -100
	-369,	// -92
	-343,	// -84
	-316,	// -76
	-289,	// -68
	-263,	// -60
	-236,	// -52
	-210,	// -44
	-183,	// -36
	-157,	// -28
	-130,	// -20
	-103,	// -12
	-77,	// -4
	-50,	// 4
	-24,	// 12
	3,		// 20
	30,		// 28
	56,		// 36
	83,		// 44
	109,	// 52
	136,	// 60
	162,	// 68
	189,	// 76
	216,	// 84
	242,	// 92
	269,	// 100
	295,	// 108
	322,	// 116
	348,	// 124
	375,	// 132
	402,	// 140
	428,	// 148
	455,	// 156
	481,	// 164
	508,	// 172
	534,	// 180
	561,	// 188
	588,	// 196
	614,	// 204
	641,	// 212
	667,	// 220
	694,	// 228
	720,	// 236
	747,	// 244
	774,	// 252
	800,	// 260
	827,	// 268
	853,	// 276
	880,	// 284
	907,	// 292
	933,	// 300
	960,	// 308
	986,	// 316
	1013,	// 324
	1039,	// 332
	1066	// 340
};

static inline
uint64 strconv$cachedFrac(int i) {
	if ((uint)i > sizeof(significands)/sizeof(significands[0])) {
		panicIndex();
	} 
	return significands[i];
}

static inline
int strconv$cachedExp(int i) {
	if ((uint)i > sizeof(exponents)/sizeof(exponents[0])) {
		panicIndex();
	} 
	return exponents[i];
}

static const
uint32 tens[] = {
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000
};

static inline
uint32 strconv$cachedTens(int i) {
	if ((uint)i > sizeof(tens)/sizeof(tens[0])) {
		panicIndex();
	} 
	return tens[i];
}
