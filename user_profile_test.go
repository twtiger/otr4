package otr4

//import (
//	"crypto/rand"
//
//	. "gopkg.in/check.v1"
//)
//
//var (
//	testTransitionSig = &dsaSignature{
//		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
//		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
//		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
//		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
//		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
//	}
//
//	testSignature = &signature{
//		0x6f, 0xee, 0xc9, 0xeb, 0x3c, 0x4a, 0x55, 0x9d,
//		0xea, 0x51, 0x02, 0x08, 0x98, 0x76, 0x0a, 0x3b,
//		0x20, 0x98, 0x92, 0xf7, 0xd0, 0x4d, 0x55, 0xfc,
//		0x86, 0x87, 0x85, 0xbc, 0x5f, 0x9c, 0xcc, 0x6d,
//		0x1c, 0x73, 0xef, 0x71, 0x9c, 0xd4, 0xbd, 0x38,
//		0xaf, 0xa4, 0xa5, 0xdb, 0xfb, 0x5f, 0x77, 0xfc,
//		0x3c, 0x8c, 0x89, 0x75, 0xa8, 0x23, 0x65, 0x06,
//		0xa2, 0xfb, 0x40, 0xd9, 0x9a, 0x82, 0x53, 0x8c,
//		0x46, 0x81, 0x23, 0xc6, 0xdc, 0xd0, 0xf0, 0x40,
//		0xd3, 0xb5, 0x29, 0xc3, 0xc1, 0xe1, 0xc3, 0x8d,
//		0xd6, 0xb6, 0x56, 0x1e, 0x00, 0xc2, 0x07, 0xb1,
//		0x90, 0x63, 0xb6, 0xa2, 0x10, 0xa7, 0xb2, 0x0f,
//		0x0d, 0x39, 0x18, 0xb9, 0x57, 0x39, 0x69, 0xda,
//		0x9e, 0x1e, 0xbf, 0xbd, 0x29, 0x4f, 0xa4, 0x00,
//	}
//
//	expiration = int64(12)
//
//	testSerBody = []byte{
//		// versions
//		0x00, 0x00, 0x00, 0x02, 0x03, 0x04,
//		// csPubKey
//		0x00, 0x10, 0xde, 0xdc, 0xb5, 0xea, 0x6c, 0x1f,
//		0x08, 0x8a, 0x15, 0x63, 0xc1, 0x76, 0x8a, 0xae,
//		0x09, 0xff, 0x41, 0x91, 0x12, 0x68, 0x1d, 0x68,
//		0xc0, 0x1a, 0x65, 0x56, 0x15, 0xb1, 0x3e, 0xe5,
//		0x12, 0x03, 0xfc, 0x82, 0x59, 0xfa, 0xd2, 0x8f,
//		0x40, 0x42, 0xa1, 0x30, 0xe5, 0xbb, 0x1e, 0x36,
//		0xb3, 0xdb, 0x32, 0x5b, 0x98, 0xc7, 0x56, 0x15,
//		0x13, 0x09, 0xc2, 0xe7, 0xd1, 0xfc, 0xce, 0x30,
//		0xc6, 0x00, 0xbc, 0x3c, 0x69, 0x7e, 0xa9, 0x8d,
//		0x82, 0x81, 0xf8, 0xae, 0xa0, 0xbc, 0xb0, 0x1f,
//		0x70, 0x3b, 0x77, 0xc6, 0xa7, 0x6e, 0xd2, 0x51,
//		0x49, 0xad, 0x9c, 0x44, 0x60, 0x22, 0x09, 0xb1,
//		0xf0, 0x08, 0xad, 0x27, 0xed, 0xfc, 0x88, 0x52,
//		0x3f, 0x2f, 0x03, 0x2f, 0xa7, 0xf1, 0x1b, 0x56,
//		0xa1, 0x6b, 0xe0, 0x73, 0x2d, 0xf4, 0xab, 0x54,
//		0x34, 0x22, 0x20, 0x26, 0xa3, 0x5a, 0x65, 0x7e,
//		0xf0, 0xf8, 0x1b, 0xf3, 0x7b, 0x42, 0x1b, 0x58,
//		0x15, 0xc7, 0x84, 0xf1, 0x99, 0x8c, 0x53, 0x7f,
//		0x66, 0xc3, 0x5a, 0xca, 0xb3, 0xf8, 0x11, 0x0d,
//		0xbb, 0x70, 0xe3, 0x81, 0x6e, 0xfc, 0x58, 0xde,
//		0x18, 0xf3, 0x44, 0x40, 0xd3, 0x36, 0x7e, 0x93,
//		0x2f, 0x5e,
//		// expiration
//		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c,
//		// transitional sig
//		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
//		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
//		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
//		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
//		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
//	}
//)
//
//func (s *OTR4Suite) Test_NewProfileBody(c *C) {
//	con := &conversation{}
//
//	keyPair, _ := deriveCramerShoupKeys(rand.Reader)
//	profile, err := con.newProfile("34", keyPair)
//
//	c.Assert(profile.versions, DeepEquals, "34")
//	c.Assert(profile, Not(IsNil))
//	c.Assert(err, IsNil)
//}
//
//func (s *OTR4Suite) Test_CreateProfileBody(c *C) {
//	keyPair, _ := deriveCramerShoupKeys(rand.Reader)
//	exp := "4"
//	profile, err := createProfileBody("4", keyPair)
//
//	c.Assert(profile.versions, DeepEquals, exp)
//	c.Assert(profile.pub, Not(IsNil))
//	c.Assert(err, IsNil)
//
//	profile, err = createProfileBody("", keyPair)
//
//	c.Assert(profile, IsNil)
//	c.Assert(err, ErrorMatches, ".* no valid version agreement could be found")
//
//	profile, err = createProfileBody("1", keyPair)
//
//	c.Assert(profile, IsNil)
//	c.Assert(err, ErrorMatches, ".* no valid version agreement could be found")
//
//	profile, err = createProfileBody("31", keyPair)
//
//	c.Assert(profile, IsNil)
//	c.Assert(err, ErrorMatches, ".* no valid version agreement could be found")
//
//	profile, err = createProfileBody("24", keyPair)
//
//	c.Assert(profile, IsNil)
//	c.Assert(err, ErrorMatches, ".* no valid version agreement could be found")
//}
//
//func (s *OTR4Suite) Test_SignUserProfile(c *C) {
//	keyPair, _ := deriveCramerShoupKeys(fixedRand(csRandData))
//	profile, err := createProfileBody("43", keyPair)
//	profile.expiration = expiration
//
//	err = profile.sign(fixedRand(csRandData), keyPair)
//
//	c.Assert(profile.sig, DeepEquals, testSignature)
//	c.Assert(err, IsNil)
//
//	err = profile.sign(fixedRand([]byte{0x00}), keyPair)
//
//	c.Assert(err, ErrorMatches, "*.cannot source enough entropy")
//}
//
//func (s *OTR4Suite) Test_VerifyUserProfile(c *C) {
//	keyPair, _ := deriveCramerShoupKeys(fixedRand(csRandData))
//	profile, err := createProfileBody("43", keyPair)
//	profile.expiration = expiration
//
//	err = profile.sign(fixedRand(csRandData), keyPair)
//
//	c.Assert(err, IsNil)
//
//	valid, err := profile.verify(keyPair.pub)
//
//	c.Assert(valid, Equals, true)
//	c.Assert(err, IsNil)
//}
//
//func (s *OTR4Suite) Test_SingAndVerifyUserProfile(c *C) {
//	keyPair, _ := deriveCramerShoupKeys(rand.Reader)
//
//	profile, err := createProfileBody("43", keyPair)
//
//	err = profile.sign(rand.Reader, keyPair)
//
//	valid, err := profile.verify(keyPair.pub)
//
//	c.Assert(valid, Equals, true)
//	c.Assert(err, IsNil)
//}
//
//func (s *OTR4Suite) Test_SerializeUserSignature(c *C) {
//	exp := [112]byte{
//		0x6f, 0xee, 0xc9, 0xeb, 0x3c, 0x4a, 0x55, 0x9d,
//		0xea, 0x51, 0x02, 0x08, 0x98, 0x76, 0x0a, 0x3b,
//		0x20, 0x98, 0x92, 0xf7, 0xd0, 0x4d, 0x55, 0xfc,
//		0x86, 0x87, 0x85, 0xbc, 0x5f, 0x9c, 0xcc, 0x6d,
//		0x1c, 0x73, 0xef, 0x71, 0x9c, 0xd4, 0xbd, 0x38,
//		0xaf, 0xa4, 0xa5, 0xdb, 0xfb, 0x5f, 0x77, 0xfc,
//		0x3c, 0x8c, 0x89, 0x75, 0xa8, 0x23, 0x65, 0x06,
//		0xa2, 0xfb, 0x40, 0xd9, 0x9a, 0x82, 0x53, 0x8c,
//		0x46, 0x81, 0x23, 0xc6, 0xdc, 0xd0, 0xf0, 0x40,
//		0xd3, 0xb5, 0x29, 0xc3, 0xc1, 0xe1, 0xc3, 0x8d,
//		0xd6, 0xb6, 0x56, 0x1e, 0x00, 0xc2, 0x07, 0xb1,
//		0x90, 0x63, 0xb6, 0xa2, 0x10, 0xa7, 0xb2, 0x0f,
//		0x0d, 0x39, 0x18, 0xb9, 0x57, 0x39, 0x69, 0xda,
//		0x9e, 0x1e, 0xbf, 0xbd, 0x29, 0x4f, 0xa4, 0x00,
//	}
//
//	c.Assert(serializeSignature(testSignature), DeepEquals, exp)
//}
//
//func (s *OTR4Suite) Test_DeserializeUserSignature(c *C) {
//	sig := []byte{
//		0x6f, 0xee, 0xc9, 0xeb, 0x3c, 0x4a, 0x55, 0x9d,
//		0xea, 0x51, 0x02, 0x08, 0x98, 0x76, 0x0a, 0x3b,
//		0x20, 0x98, 0x92, 0xf7, 0xd0, 0x4d, 0x55, 0xfc,
//		0x86, 0x87, 0x85, 0xbc, 0x5f, 0x9c, 0xcc, 0x6d,
//		0x1c, 0x73, 0xef, 0x71, 0x9c, 0xd4, 0xbd, 0x38,
//		0xaf, 0xa4, 0xa5, 0xdb, 0xfb, 0x5f, 0x77, 0xfc,
//		0x3c, 0x8c, 0x89, 0x75, 0xa8, 0x23, 0x65, 0x06,
//		0xa2, 0xfb, 0x40, 0xd9, 0x9a, 0x82, 0x53, 0x8c,
//		0x46, 0x81, 0x23, 0xc6, 0xdc, 0xd0, 0xf0, 0x40,
//		0xd3, 0xb5, 0x29, 0xc3, 0xc1, 0xe1, 0xc3, 0x8d,
//		0xd6, 0xb6, 0x56, 0x1e, 0x00, 0xc2, 0x07, 0xb1,
//		0x90, 0x63, 0xb6, 0xa2, 0x10, 0xa7, 0xb2, 0x0f,
//		0x0d, 0x39, 0x18, 0xb9, 0x57, 0x39, 0x69, 0xda,
//		0x9e, 0x1e, 0xbf, 0xbd, 0x29, 0x4f, 0xa4, 0x00,
//	}
//
//	c.Assert(deserializeSignature(sig), DeepEquals, testSignature)
//}
//
//func (s *OTR4Suite) Test_DeserializeUserTransitionSignature(c *C) {
//	sig := []byte{
//		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
//		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
//		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
//		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
//		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
//	}
//
//	c.Assert(deserializeTransSignature(sig), DeepEquals, testTransitionSig)
//}
//
//func (s *OTR4Suite) Test_SerializeUserProfileBody(c *C) {
//	keyPair, _ := deriveCramerShoupKeys(fixedRand(csRandData))
//	profile, err := createProfileBody("34", keyPair)
//	profile.expiration = expiration
//
//	ser := serializeBody(profile)
//
//	c.Assert(ser, DeepEquals, testSerBody[:184])
//	c.Assert(err, IsNil)
//
//	profile.transitionSig = testTransitionSig
//
//	ser = serializeBody(profile)
//
//	c.Assert(ser, DeepEquals, testSerBody)
//}
//
//func (s *OTR4Suite) Test_SerializeUserProfile(c *C) {
//	r := fixedRand([]byte{
//		0x40, 0x80, 0x66, 0x2d, 0xd8, 0xe7, 0xf0, 0x9c,
//		0xdf, 0xb0, 0x4e, 0x1c, 0x6e, 0x12, 0x62, 0xa3,
//		0x7c, 0x31, 0x9a, 0xe1, 0xe7, 0x86, 0x87, 0xcc,
//		0x82, 0x05, 0x78, 0xe6, 0x44, 0x2f, 0x4f, 0x77,
//	},
//	)
//
//	con := &conversation{
//		random: r,
//	}
//
//	keyPair, _ := deriveCramerShoupKeys(fixedRand(csRandData))
//	profile, _ := con.newProfile("34", keyPair)
//	profile.transitionSig = testTransitionSig
//
//	ser := profile.serialize()
//
//	// version
//	c.Assert(ser[:6], DeepEquals, testSerBody[:6])
//	// pub key
//	c.Assert(ser[6:176], DeepEquals, testSerBody[6:176])
//	// transition sig
//	c.Assert(ser[184:216], DeepEquals, testSerBody[184:216])
//}
//
//func (s *OTR4Suite) Test_DeserializeUserProfile(c *C) {
//	ser := []byte{
//		// versions
//		0x00, 0x00, 0x00, 0x02, 0x03, 0x04,
//		// csPubKey
//		0x00, 0x10, 0xde, 0xdc, 0xb5, 0xea, 0x6c, 0x1f,
//		0x08, 0x8a, 0x15, 0x63, 0xc1, 0x76, 0x8a, 0xae,
//		0x09, 0xff, 0x41, 0x91, 0x12, 0x68, 0x1d, 0x68,
//		0xc0, 0x1a, 0x65, 0x56, 0x15, 0xb1, 0x3e, 0xe5,
//		0x12, 0x03, 0xfc, 0x82, 0x59, 0xfa, 0xd2, 0x8f,
//		0x40, 0x42, 0xa1, 0x30, 0xe5, 0xbb, 0x1e, 0x36,
//		0xb3, 0xdb, 0x32, 0x5b, 0x98, 0xc7, 0x56, 0x15,
//		0x13, 0x09, 0xc2, 0xe7, 0xd1, 0xfc, 0xce, 0x30,
//		0xc6, 0x00, 0xbc, 0x3c, 0x69, 0x7e, 0xa9, 0x8d,
//		0x82, 0x81, 0xf8, 0xae, 0xa0, 0xbc, 0xb0, 0x1f,
//		0x70, 0x3b, 0x77, 0xc6, 0xa7, 0x6e, 0xd2, 0x51,
//		0x49, 0xad, 0x9c, 0x44, 0x60, 0x22, 0x09, 0xb1,
//		0xf0, 0x08, 0xad, 0x27, 0xed, 0xfc, 0x88, 0x52,
//		0x3f, 0x2f, 0x03, 0x2f, 0xa7, 0xf1, 0x1b, 0x56,
//		0xa1, 0x6b, 0xe0, 0x73, 0x2d, 0xf4, 0xab, 0x54,
//		0x34, 0x22, 0x20, 0x26, 0xa3, 0x5a, 0x65, 0x7e,
//		0xf0, 0xf8, 0x1b, 0xf3, 0x7b, 0x42, 0x1b, 0x58,
//		0x15, 0xc7, 0x84, 0xf1, 0x99, 0x8c, 0x53, 0x7f,
//		0x66, 0xc3, 0x5a, 0xca, 0xb3, 0xf8, 0x11, 0x0d,
//		0xbb, 0x70, 0xe3, 0x81, 0x6e, 0xfc, 0x58, 0xde,
//		0x18, 0xf3, 0x44, 0x40, 0xd3, 0x36, 0x7e, 0x93,
//		0x2f, 0x5e,
//		// expiration
//		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c,
//		// transitional sig
//		0xee, 0xec, 0x0c, 0xa7, 0x39, 0x65, 0x3c, 0x35,
//		0xe2, 0x28, 0xd3, 0xc8, 0xc1, 0x07, 0x96, 0xeb,
//		0x06, 0xe8, 0x14, 0x05, 0x62, 0x52, 0xab, 0x6c,
//		0x63, 0xf1, 0x4f, 0x55, 0xb3, 0xea, 0x9b, 0x1d,
//		0xbf, 0xe7, 0xb7, 0xec, 0x8b, 0x52, 0x43, 0x46,
//		// signature
//		0x6f, 0xee, 0xc9, 0xeb, 0x3c, 0x4a, 0x55, 0x9d,
//		0xea, 0x51, 0x02, 0x08, 0x98, 0x76, 0x0a, 0x3b,
//		0x20, 0x98, 0x92, 0xf7, 0xd0, 0x4d, 0x55, 0xfc,
//		0x86, 0x87, 0x85, 0xbc, 0x5f, 0x9c, 0xcc, 0x6d,
//		0x1c, 0x73, 0xef, 0x71, 0x9c, 0xd4, 0xbd, 0x38,
//		0xaf, 0xa4, 0xa5, 0xdb, 0xfb, 0x5f, 0x77, 0xfc,
//		0x3c, 0x8c, 0x89, 0x75, 0xa8, 0x23, 0x65, 0x06,
//		0xa2, 0xfb, 0x40, 0xd9, 0x9a, 0x82, 0x53, 0x8c,
//		0x46, 0x81, 0x23, 0xc6, 0xdc, 0xd0, 0xf0, 0x40,
//		0xd3, 0xb5, 0x29, 0xc3, 0xc1, 0xe1, 0xc3, 0x8d,
//		0xd6, 0xb6, 0x56, 0x1e, 0x00, 0xc2, 0x07, 0xb1,
//		0x90, 0x63, 0xb6, 0xa2, 0x10, 0xa7, 0xb2, 0x0f,
//		0x0d, 0x39, 0x18, 0xb9, 0x57, 0x39, 0x69, 0xda,
//		0x9e, 0x1e, 0xbf, 0xbd, 0x29, 0x4f, 0xa4, 0x00,
//	}
//
//	profile := deserializeProfile(ser)
//
//	// version
//	c.Assert(profile.versions, DeepEquals, "34")
//	// pub key XXX: test the value
//	// expiration
//	c.Assert(profile.expiration, DeepEquals, int64(12))
//	// transition sig
//	c.Assert(profile.transitionSig, DeepEquals, testTransitionSig)
//	// signature
//	c.Assert(profile.sig, DeepEquals, testSignature)
//
//	ser = []byte{
//		// versions
//		0x00, 0x00, 0x00, 0x02, 0x03, 0x04,
//		// invalid csPubKey
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
//		0xff, 0xff,
//	}
//
//	profile = deserializeProfile(ser)
//
//	c.Assert(profile, IsNil)
//}
