package otr4

import (
	"github.com/otrv4/ed448"
)

var (
	testByteSlice = []byte{
		0xad, 0xd0, 0x35, 0x07, 0x1d, 0x09, 0x6c, 0xd3,
		0xdd, 0xf8, 0x96, 0x59, 0x39, 0x1c, 0x29, 0xa2,
		0x1e, 0x49, 0x34, 0xae, 0xc1, 0x79, 0x0e, 0x85,
		0x1c, 0x06, 0x73, 0xf2, 0xdd, 0x5d, 0x39, 0x71,
		0xf5, 0x70, 0x71, 0x4d, 0x5c, 0xca, 0x18, 0x02,
		0xaf, 0xa3, 0x85, 0x1b, 0x8a, 0x53, 0x39, 0xb7,
		0xa2, 0x33, 0x1b, 0x8a, 0x53, 0x39, 0xb7, 0xa2,
		0x33, 0x2a, 0xf4, 0xf7, 0xb6, 0x26, 0x37, 0x3e,
		0xb7, 0xd5, 0x9a, 0x1b, 0x3c, 0xf2, 0xfd, 0x63,
	}

	testPoint = ed448.NewPoint(
		[16]uint32{
			0x03aee5b7, 0x032b3f7f, 0x06176ed0, 0x056fa571,
			0x04d01a0b, 0x0b382fa3, 0x03d55289, 0x04c8e69c,
			0x0ab17cae, 0x07b995ce, 0x03263f63, 0x08d4efc8,
			0x0382c935, 0x0587fbd5, 0x03d42439, 0x0b6979c9,
		},
		[16]uint32{
			0x070e385a, 0x04b3e06c, 0x0d8e3b40, 0x04064fff,
			0x04c1ae53, 0x0348a758, 0x04fa81f9, 0x06ef7f5a,
			0x0be2c435, 0x0b6c6794, 0x0159c719, 0x0350c7e1,
			0x0a5d1620, 0x0c9e7983, 0x0c90bd0e, 0x01196a72,
		},
		[16]uint32{
			0x0174aff6, 0x0aa1703e, 0x0b3d41e7, 0x0d68f123,
			0x0fcc832b, 0x0c11adbe, 0x0faecb96, 0x0152998d,
			0x0902ed06, 0x03560403, 0x067b0008, 0x0825eef3,
			0x0bb42471, 0x0ae05b98, 0x06c93c9b, 0x0bb36c4b,
		},
		[16]uint32{
			0x0970457b, 0x0bbb1746, 0x065927c4, 0x012c45ac,
			0x03a587e6, 0x095e2a6c, 0x0ec77f11, 0x09226042,
			0x0c304e97, 0x01783fbf, 0x0f4f1dbd, 0x07628142,
			0x0981f1fd, 0x0311dbc5, 0x05c12822, 0x033e76b1,
		},
	)

	randData = []byte{
		0x40, 0x80, 0x66, 0x2d, 0xd8, 0xe7, 0xf0, 0x9c,
		0xdf, 0xb0, 0x4e, 0x1c, 0x6e, 0x12, 0x62, 0xa3,
		0x7c, 0x31, 0x9a, 0xe1, 0xe7, 0x86, 0x87, 0xcc,
		0x82, 0x05, 0x78, 0xe6, 0x44, 0x2f, 0x4f, 0x77,
		0x0e, 0xd1, 0xb4, 0x48, 0xa6, 0x05, 0x90, 0x5e,
		0xe7, 0xba, 0xfc, 0x25, 0x99, 0x99, 0xb8, 0xc3,
		0x90, 0x3e, 0xf4, 0xa3, 0x75, 0xee, 0x85, 0x32,
		0x16, 0xb1, 0x06, 0x5b, 0x81, 0xea, 0xac, 0xb3,
		0x69, 0x47, 0x6d, 0xa2, 0xaa, 0x86, 0x0b, 0xe5,
		0xcd, 0xac, 0x43, 0xd7, 0xb7, 0xe3, 0xb0, 0x85,
		0xd8, 0x66, 0xf9, 0xb6, 0x45, 0x2e, 0x81, 0x43,
		0xc2, 0x6f, 0x61, 0xc4, 0xdd, 0x65, 0x35, 0xa4,
		0xa4, 0xf9, 0x55, 0xf0, 0xf9, 0xd2, 0xf4, 0xb7,
		0xa4, 0xf9, 0x55, 0xf0, 0xf9, 0xd2, 0xf4, 0xb7,
		0x52, 0x18, 0x41, 0x48, 0x60, 0x2d, 0x67, 0x8a,
		0xd3, 0xf3, 0xd2, 0xa4, 0xfd, 0x6f, 0x64, 0xf3,
		0x72, 0x82, 0xb0, 0x6a, 0x4d, 0xea, 0x9c, 0xef,
		0x99, 0x05, 0xe1, 0x8d, 0xaf, 0x2d, 0xdb, 0x52,
		0x57, 0x00, 0xac, 0x45, 0x24, 0x24, 0xb4, 0x79,
		0x02, 0x5f, 0x99, 0x70, 0x95, 0x2a, 0x90, 0x08,
		0x02, 0x5f, 0x99, 0x70, 0x95, 0x2a, 0x90, 0x08,
		0x51, 0x5b, 0x69, 0x03, 0xd5, 0x77, 0xb0, 0x77,
		0x35, 0x1f, 0x1b, 0x2d, 0xb1, 0x26, 0xf1, 0x69,
		0x3b, 0xcc, 0x4b, 0x0a, 0x95, 0x83, 0xd7, 0xec,
		0xfa, 0x8c, 0xf7, 0x80, 0xbe, 0x9b, 0x6d, 0xb4,
		0xc3, 0x24, 0x3c, 0x94, 0x9b, 0x63, 0xbc, 0x89,
		0xbc, 0x09, 0x39, 0xb8, 0xbf, 0xa2, 0x9b, 0xf4,
		0x3a, 0xa2, 0x9b, 0xbe, 0x6e, 0x78, 0x7b, 0x11,
		0x66, 0x60, 0x01, 0xb9, 0x83, 0x10, 0xd5, 0x7d,
		0xe4, 0x86, 0x58, 0x0a, 0x42, 0xd2, 0x2a, 0x74,
		0xe9, 0x5d, 0x77, 0xc4, 0x08, 0x46, 0x31, 0xb4,
		0x75, 0x1b, 0xf2, 0x67, 0x23, 0x19, 0x5e, 0xb6,
		0xfc, 0xe8, 0xd1, 0x38, 0x81, 0xa3, 0x98, 0x41,
		0xdf, 0xdf, 0x5d, 0x8d, 0x41, 0xb4, 0x66, 0x0f,
		0x39, 0xe1, 0x6f, 0x8c, 0x89, 0xed, 0xf6, 0x11,
	}

	randAuthData = []byte{
		0xea, 0x25, 0xbc, 0x1d, 0x8d, 0x18, 0x2f, 0xe2,
		0x33, 0xe4, 0xd1, 0x8c, 0x58, 0xac, 0x3a, 0x75,
		0x32, 0x7a, 0xb0, 0x91, 0xcf, 0x85, 0x81, 0xf8,
		0x2c, 0xc5, 0xf3, 0x55, 0x3d, 0x32, 0x2d, 0x3e,
		0x8c, 0x0b, 0xf5, 0xfb, 0x6a, 0x11, 0xf9, 0x5b,
		0x35, 0xcc, 0x4b, 0xda, 0x05, 0x8a, 0x3c, 0x66,
		0x17, 0x65, 0x14, 0xf4, 0x04, 0x40, 0xd3, 0x02,

		0xad, 0x6a, 0xc8, 0x1c, 0xb1, 0x91, 0xef, 0xcb,
		0xfa, 0x5d, 0x81, 0xbe, 0xf3, 0xa2, 0x01, 0x0f,
		0xb2, 0x32, 0x97, 0x2c, 0x2c, 0x46, 0xb6, 0xd2,
		0x55, 0x19, 0xad, 0xb4, 0x57, 0x74, 0xe6, 0x61,
		0xb0, 0xe2, 0x1a, 0x22, 0x91, 0x82, 0xe4, 0x9a,
		0xf7, 0xff, 0x82, 0x5b, 0x4f, 0xeb, 0x05, 0x2b,
		0x0d, 0xcf, 0x78, 0xe0, 0x02, 0x53, 0x6d, 0x33,

		0xd7, 0xfb, 0x29, 0xa2, 0xd1, 0x3d, 0x98, 0xec,
		0x18, 0xe8, 0x91, 0x07, 0x22, 0x13, 0xbf, 0x76,
		0x5a, 0x2f, 0x93, 0xa6, 0xd4, 0xcf, 0xe7, 0x77,
		0x99, 0x73, 0xbb, 0x20, 0x4e, 0x09, 0xb1, 0xc0,
		0x68, 0x93, 0xf1, 0xf2, 0x35, 0x28, 0x5b, 0xa3,
		0x89, 0x9d, 0x76, 0x75, 0x46, 0x5b, 0xe9, 0xa1,
		0xff, 0x1c, 0x2e, 0x8e, 0xbf, 0x5e, 0x22, 0x15,

		0x2c, 0x80, 0x3f, 0x45, 0x89, 0x58, 0x2d, 0x7e,
		0xda, 0x26, 0x01, 0xa7, 0x78, 0xa6, 0xa6, 0x61,
		0xa3, 0x56, 0x5c, 0x0b, 0xe0, 0xae, 0x09, 0xea,
		0x56, 0xce, 0x09, 0x17, 0xc0, 0x98, 0xd5, 0x40,
		0xe8, 0x9b, 0x18, 0xdf, 0xb7, 0x78, 0x44, 0xc3,
		0xc9, 0x66, 0x62, 0x85, 0xae, 0xa6, 0xd7, 0x10,
		0x7b, 0x14, 0xf0, 0x40, 0xd7, 0x72, 0x92, 0x2b,

		0x02, 0x35, 0xbd, 0x52, 0x34, 0x56, 0x9b, 0x58,
		0x2c, 0x39, 0xaf, 0x2e, 0x92, 0xca, 0x6c, 0x0a,
		0x81, 0x22, 0x88, 0x38, 0xd3, 0xdd, 0x17, 0x25,
		0x27, 0xc9, 0x2d, 0xf6, 0x4d, 0xa1, 0xf2, 0x9c,
		0xbd, 0x08, 0xf4, 0xa0, 0x91, 0x08, 0x79, 0xf6,
		0x8a, 0x78, 0x3c, 0xf0, 0xac, 0x2d, 0x97, 0x03,
		0x54, 0xe3, 0xc6, 0x22, 0xb9, 0xf4, 0x55, 0x3a,
	}

	testMessage = []byte{
		0x5d, 0xf1, 0x18, 0xbf, 0x8e, 0x3f, 0xfe, 0xcd,
		0x95, 0xd3, 0x49, 0xda, 0xcd, 0xac, 0x2c, 0xdf,
		0x72, 0x5e, 0xb7, 0x61, 0x44, 0xf1, 0x93, 0xa6,
		0x70, 0x8e, 0x64, 0xff, 0x7c, 0xec, 0x6c, 0xe5,
		0xc6, 0x8d, 0x8f, 0xa0, 0x43, 0x23, 0x45, 0x33,
		0x73, 0x71, 0xe6, 0x2f, 0x57, 0xbb, 0x0f, 0x70,
		0x11, 0x8c, 0x62, 0x26, 0x9e, 0x17, 0x5d, 0x22,
	}

	testPubA = &publicKey{
		ed448.NewPoint(
			[16]uint32{
				0x0f40a704, 0x08736fbf, 0x08bf4671, 0x041c8cd0,
				0x02b9e934, 0x0a5cea06, 0x0e5626ea, 0x0a18011a,
				0x0b6a4534, 0x08790731, 0x03aedcc2, 0x0f0b6142,
				0x0ac3f958, 0x0f5d4e38, 0x0ae39cf2, 0x07c929f4,
			},

			[16]uint32{
				0x0dd36c68, 0x06c8827f, 0x010cbebf, 0x00fc1eeb,
				0x000307e4, 0x045ae00b, 0x00296716, 0x0265848f,
				0x0f227e15, 0x094cda83, 0x0c9894a0, 0x08b8a971,
				0x0acfd16b, 0x0f0923f0, 0x0046eef4, 0x01834ec8,
			},
			[16]uint32{
				0x0c37735d, 0x04359f41, 0x08dc5484, 0x007be35b,
				0x03d2cdec, 0x06838092, 0x0bd4c68e, 0x08c8858d,
				0x09fa4b07, 0x0f7534a9, 0x04ec159b, 0x0742a5b1,
				0x01d906d0, 0x037452d4, 0x09d529e9, 0x0bcf8832,
			},
			[16]uint32{
				0x08352010, 0x0fdefde8, 0x0e480543, 0x0839f9bd,
				0x0d98f832, 0x0314b8b1, 0x0df01a60, 0x0b25fcd1,
				0x0d2a033c, 0x0aa474e4, 0x0308c005, 0x0303aa3b,
				0x07da57b9, 0x08870b3a, 0x049a2edf, 0x0630f320,
			},
		),
	}

	testPrivA = &privateKey{
		ed448.NewScalar([]byte{
			0x13, 0x66, 0x00, 0x41, 0x14, 0x93, 0x97, 0x66,
			0x8a, 0x8d, 0xf2, 0xd3, 0x20, 0x77, 0xa6, 0x5e,
			0x9b, 0x5f, 0x97, 0x7c, 0x39, 0x34, 0xbe, 0xf3,
			0xd6, 0x0d, 0xcc, 0x7c, 0x42, 0x1d, 0xa3, 0x3b,
			0xb4, 0xc1, 0x31, 0x3a, 0xbc, 0x3a, 0x70, 0xec,
			0xd1, 0xb4, 0xa0, 0xee, 0x39, 0x92, 0xc1, 0x43,
			0xbd, 0x87, 0xc0, 0xbd, 0x62, 0x17, 0x54, 0x12,
		}),
	}

	testPubB = &publicKey{
		ed448.NewPoint(
			[16]uint32{
				0x048603ce, 0x08657fa3, 0x0a7ed7d1, 0x00467214,
				0x0781aeaf, 0x06202e8e, 0x0142a539, 0x0c55ab78,
				0x05405585, 0x0c7d68bc, 0x0ffe9cc1, 0x061888b5,
				0x0d994802, 0x05147e54, 0x0f533d6f, 0x023ab901,
			},
			[16]uint32{
				0x06cf7d90, 0x06269f0e, 0x0ab6c4e2, 0x09804fd9,
				0x048bab98, 0x0e33fdcf, 0x0996c34a, 0x0ab80d7e,
				0x0d59830b, 0x058a5d0b, 0x026dbc6a, 0x0e2cb254,
				0x04664218, 0x0b106106, 0x011ced8c, 0x00851398,
			},
			[16]uint32{
				0x0df8cf30, 0x07b60230, 0x076e63ae, 0x06998644,
				0x05abd81b, 0x0ecd45f8, 0x0e03891e, 0x0bb625fa,
				0x09337432, 0x0a2dec2d, 0x0cc75d04, 0x023ff0ef,
				0x0be80126, 0x0cf82eed, 0x057fd1f6, 0x0fa34c11,
			},
			[16]uint32{
				0x03800dae, 0x0c0a6db2, 0x06019eaf, 0x000dda7c,
				0x0162bcd6, 0x0327a6c6, 0x0b772fef, 0x0754b0fe,
				0x00816e60, 0x0626b4f5, 0x01cb5bea, 0x0623ffcd,
				0x0c288061, 0x025fd8c5, 0x08dca40d, 0x05c1517d,
			},
		),
	}

	testPrivB = &privateKey{
		ed448.NewScalar([]byte{
			0xb9, 0x1c, 0xa1, 0xe6, 0x54, 0xb5, 0xdc, 0x03,
			0x11, 0x0e, 0x6f, 0xa8, 0x52, 0x6b, 0x3d, 0x7c,
			0x46, 0xbd, 0xd6, 0x1b, 0x52, 0x8b, 0x18, 0xa4,
			0x46, 0x32, 0x10, 0x1a, 0x57, 0xab, 0x51, 0xbe,
			0x9a, 0x90, 0xdd, 0x76, 0xa7, 0x82, 0x72, 0x0e,
			0x5e, 0x91, 0xda, 0x1b, 0xbd, 0xca, 0x47, 0x46,
			0xfb, 0x49, 0xc1, 0x74, 0x8f, 0xa5, 0x79, 0x2e,
		},
		),
	}

	testPubC = ed448.NewPoint(
		[16]uint32{
			0xbaaead5, 0x685c976, 0xb0ca061, 0xba86cd7,
			0xea519fa, 0x57e4ab4, 0xc02f5fc, 0xb82204c,
			0xd78542c, 0x74a01e5, 0x9328d21, 0x3871d1a,
			0x5dcaa1c, 0x0156612, 0x7ea2255, 0xd9d787d,
		},
		[16]uint32{
			0x1a1f403, 0xbfea69f, 0x4a6d633, 0xa60a88d,
			0x4e36fdc, 0x8028db6, 0xc2fe5ba, 0xc58f5c6,
			0x85375d6, 0xbef9d4f, 0x2953a6a, 0xa779d7a,
			0xb729468, 0x9b47792, 0x0ac10fe, 0x434eff1,
		},
		[16]uint32{
			0x803ebe1, 0x18470ae, 0xb80caad, 0x3b777f9,
			0x67a6139, 0x6d9aa39, 0x787cf2e, 0x7c11d72,
			0x0374caf, 0x02cd168, 0x3d6858b, 0xdac0675,
			0xcaae54b, 0xffe21b4, 0x4bb4af8, 0x55a2cce,
		},
		[16]uint32{
			0x384109e, 0x40695a8, 0xa822b8e, 0x6026944,
			0x8e9ae46, 0xaad36b0, 0xa10ec79, 0x505a46c,
			0x4e7c598, 0x0b9daf8, 0xe22fb37, 0xa2aeb13,
			0x126a250, 0x874aa07, 0x0fe3b33, 0x1be9c1d,
		})

	invalidPub = &publicKey{
		ed448.NewPoint(
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
			[16]uint32{
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
				0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff,
			},
		),
	}

	tmpSerPubA = []byte{
		0x00, 0x10, 0xb3, 0x64, 0x1e, 0x39, 0x8f, 0xd8,
		0x23, 0xf2, 0x89, 0xd9, 0xfa, 0xa8, 0x9f, 0x14,
		0x83, 0xb2, 0xb9, 0x65, 0xb3, 0x1c, 0xfa, 0x84,
		0x82, 0xcb, 0xda, 0x45, 0xb6, 0x22, 0x52, 0xd7,
		0xda, 0xff, 0xd2, 0xdb, 0xb2, 0x1f, 0x7a, 0x2f,
		0x8b, 0x1a, 0x68, 0x16, 0x6f, 0x74, 0x90, 0x56,
		0xb3, 0x8a, 0x4f, 0xdd, 0x66, 0x3c, 0x3d, 0xcc,
		0xb7, 0x8b, 0x80,
	}

	serPubA = []byte{
		0x00, 0x10, 0x02, 0x5d, 0x7b, 0x83, 0xd3, 0xb8,
		0xea, 0xe6, 0xde, 0x02, 0xfd, 0x69, 0xdc, 0xe5,
		0xe7, 0xb8, 0xe6, 0x19, 0x4b, 0x35, 0x9f, 0xa4,
		0x34, 0x4a, 0xb5, 0x5a, 0xa3, 0x83, 0xfd, 0x9b,
		0xe0, 0x25, 0x87, 0x97, 0x58, 0x78, 0x38, 0xd6,
		0xce, 0xf2, 0xf8, 0x8a, 0xa3, 0xa0, 0x71, 0x97,
		0x37, 0x31, 0x39, 0xc2, 0x50, 0xa5, 0xd4, 0x10,
		0xd6, 0x60,
	}

	testSigma = &authMessage{
		ed448.NewScalar([]byte{
			0xb7, 0x45, 0x7c, 0xd2, 0x31, 0x90, 0xbf, 0xcd,
			0xcb, 0x09, 0x14, 0x76, 0x87, 0x19, 0xb8, 0xad,
			0x94, 0xe6, 0x6c, 0x1a, 0x61, 0x9d, 0x2e, 0xcf,
			0x6b, 0xcf, 0xd3, 0xc8, 0x2b, 0xa3, 0x38, 0x4f,
			0x00, 0x09, 0x06, 0xb3, 0xce, 0x50, 0xfe, 0x79,
			0x98, 0x85, 0x84, 0x6c, 0x23, 0xf6, 0x2e, 0xbb,
			0x6c, 0x13, 0xeb, 0xbe, 0x82, 0x59, 0xd3, 0x11,
		}),
		ed448.NewScalar([]byte{
			0x34, 0x51, 0x89, 0xe2, 0xf6, 0xf6, 0xfb, 0x6a,
			0xa1, 0xfd, 0xc3, 0xc2, 0x98, 0xc1, 0xe6, 0x84,
			0x6e, 0xf5, 0x91, 0x59, 0x90, 0x78, 0x7e, 0x9a,
			0xb0, 0xa7, 0x05, 0xbf, 0x41, 0x2d, 0x70, 0x41,
			0x98, 0x46, 0xb8, 0x69, 0xf0, 0x84, 0x77, 0x84,
			0x20, 0x07, 0xc2, 0xe0, 0x26, 0x01, 0x6d, 0xb3,
			0x9f, 0xc9, 0xbe, 0x99, 0x15, 0x2f, 0x47, 0x20,
		}),
		ed448.NewScalar([]byte{
			0xad, 0x6a, 0xc8, 0x1c, 0xb1, 0x91, 0xef, 0xcb,
			0xfa, 0x5d, 0x81, 0xbe, 0xf3, 0xa2, 0x01, 0x0f,
			0xb2, 0x32, 0x97, 0x2c, 0x2c, 0x46, 0xb6, 0xd2,
			0x55, 0x19, 0xad, 0xb4, 0x57, 0x74, 0xe6, 0x61,
			0xb0, 0xe2, 0x1a, 0x22, 0x91, 0x82, 0xe4, 0x9a,
			0xf7, 0xff, 0x82, 0x5b, 0x4f, 0xeb, 0x05, 0x2b,
			0x0d, 0xcf, 0x78, 0xe0, 0x2, 0x53, 0x6d, 0x33,
		}),
		ed448.NewScalar([]byte{
			0x2c, 0x80, 0x3f, 0x45, 0x89, 0x58, 0x2d, 0x7e,
			0xda, 0x26, 0x01, 0xa7, 0x78, 0xa6, 0xa6, 0x61,
			0xa3, 0x56, 0x5c, 0x0b, 0xe0, 0xae, 0x09, 0xea,
			0x56, 0xce, 0x09, 0x17, 0xc0, 0x98, 0xd5, 0x40,
			0xe8, 0x9b, 0x18, 0xdf, 0xb7, 0x78, 0x44, 0xc3,
			0xc9, 0x66, 0x62, 0x85, 0xae, 0xa6, 0xd7, 0x10,
			0x7b, 0x14, 0xf0, 0x40, 0xd7, 0x72, 0x92, 0x2b,
		}),
		ed448.NewScalar([]byte{
			0xd7, 0xfb, 0x29, 0xa2, 0xd1, 0x3d, 0x98, 0xec,
			0x18, 0xe8, 0x91, 0x07, 0x22, 0x13, 0xbf, 0x76,
			0x5a, 0x2f, 0x93, 0xa6, 0xd4, 0xcf, 0xe7, 0x77,
			0x99, 0x73, 0xbb, 0x20, 0x4e, 0x09, 0xb1, 0xc0,
			0x68, 0x93, 0xf1, 0xf2, 0x35, 0x28, 0x5b, 0xa3,
			0x89, 0x9d, 0x76, 0x75, 0x46, 0x5b, 0xe9, 0xa1,
			0xff, 0x1c, 0x2e, 0x8e, 0xbf, 0x5e, 0x22, 0x15,
		}),
		ed448.NewScalar([]byte{
			0x02, 0x35, 0xbd, 0x52, 0x34, 0x56, 0x9b, 0x58,
			0x2c, 0x39, 0xaf, 0x2e, 0x92, 0xca, 0x6c, 0x0a,
			0x81, 0x22, 0x88, 0x38, 0xd3, 0xdd, 0x17, 0x25,
			0x27, 0xc9, 0x2d, 0xf6, 0x4d, 0xa1, 0xf2, 0x9c,
			0xbd, 0x08, 0xf4, 0xa0, 0x91, 0x08, 0x79, 0xf6,
			0x8a, 0x78, 0x3c, 0xf0, 0xac, 0x2d, 0x97, 0x03,
			0x54, 0xe3, 0xc6, 0x22, 0xb9, 0xf4, 0x55, 0x3a,
		}),
	}
)
