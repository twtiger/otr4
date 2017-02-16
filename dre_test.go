package otr4

import (
	"crypto/rand"
	"errors"

	"testing"

	"github.com/twstrike/ed448"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

// XXX: move to appropiate place
type OTR4Suite struct{}

var _ = Suite(&OTR4Suite{})

func (s *OTR4Suite) Test_HashToScalar(c *C) {
	scalar := hashToScalar(testByteSlice)

	exp := ed448.NewDecafScalar([]byte{
		0x1e, 0xda, 0x47, 0xce, 0x5a, 0x2a, 0x10, 0xdb,
		0x67, 0x8a, 0x38, 0x2c, 0xe2, 0x70, 0x2f, 0xea,
		0x92, 0x8d, 0x6a, 0x4c, 0x11, 0x27, 0xfd, 0x7c,
		0xb0, 0x6f, 0x1a, 0x0b, 0x71, 0x82, 0x6b, 0x90,
		0xe3, 0x6b, 0xdd, 0x7d, 0x17, 0xab, 0xfd, 0x9e,
		0xad, 0xf2, 0x04, 0x0d, 0x97, 0x19, 0x46, 0x09,
		0x3d, 0xb3, 0xa3, 0x67, 0xca, 0x01, 0x8d, 0x95,
	})

	c.Assert(scalar, DeepEquals, exp)
}

func (s *OTR4Suite) Test_Concat(c *C) {
	empty := []byte{}
	bytes := []byte{
		0x04, 0x2a, 0xf3, 0xcc, 0x69, 0xbb, 0xa1, 0x50,
	}

	exp := []byte{
		0x04, 0x2a, 0xf3, 0xcc, 0x69, 0xbb, 0xa1, 0x50,
		0x71, 0x7b, 0x24, 0xd5, 0xd4, 0x98, 0x0c, 0xfe,
		0xce, 0x60, 0xe7, 0x97, 0x84, 0xf4, 0x1c, 0x72,
		0x01, 0x07, 0xb8, 0x24, 0xa8, 0x43, 0x0e, 0x81,
		0x25, 0xca, 0xb4, 0xa0, 0xda, 0xf5, 0xfa, 0xf6,
		0x0c, 0x90, 0x99, 0x7f, 0x1e, 0xed, 0x83, 0xde,
		0xbe, 0xe7, 0xef, 0x8e, 0xea, 0xeb, 0xc8, 0x5d,
		0x67, 0x5b, 0x3b, 0x04, 0x55, 0x0a, 0x36, 0x2f,
		0x06, 0xea, 0x48, 0xc4, 0x23, 0x28, 0xe1, 0x99,
		0x08, 0xa5, 0x88, 0x8f, 0xad, 0x7f, 0x39, 0xdf,
		0x56, 0xa3, 0xaa, 0x4d, 0x59, 0x66, 0xec, 0xd5,
		0x6c, 0x38, 0x02, 0x8c, 0x80, 0x96, 0xd2, 0xd4,
		0x54, 0x24, 0x76, 0x70, 0xda, 0x99, 0xc5, 0xd6,
		0x81, 0x40, 0x49, 0xcd, 0x76, 0xb1, 0x05, 0xc4,
		0xa8, 0x42, 0x17, 0x09, 0x51, 0xc2, 0xa9, 0x2e,
	}

	c.Assert(func() { concat() }, Panics, "missing concat arguments")
	c.Assert(func() { concat(bytes) }, Panics, "missing concat arguments")
	c.Assert(func() { concat("not a valid input", bytes) }, Panics, "invalid input")
	c.Assert(concat(empty, bytes, testSec, testPubA), DeepEquals, exp)
}

func (s *OTR4Suite) Test_Auth(c *C) {
	message := []byte("our message")
	out, err := auth(fixedRand(randData), testPubA, testPubB, testPubC, testSec, message)

	c.Assert(out, DeepEquals, testSigma)
	c.Assert(err, IsNil)

	r := make([]byte, 56*5-1)
	out, err = auth(fixedRand(r), testPubA, testPubB, testPubC, testSec, message)

	c.Assert(err, DeepEquals, errors.New("unexpected EOF: not enough bytes"))
	c.Assert(out, IsNil)

	r = make([]byte, 56)
	out, err = auth(fixedRand(r), testPubA, testPubB, testPubC, testSec, message)

	c.Assert(err, DeepEquals, errors.New("unexpected EOF: not enough bytes"))
	c.Assert(out, IsNil)
}

func (s *OTR4Suite) Test_Verify(c *C) {
	message := []byte("our message")

	b := verify(testPubA, testPubB, testPubC, testSigma, message)

	c.Assert(b, Equals, true)
}

func (s *OTR4Suite) Test_VerifyAndAuth(c *C) {
	message := []byte("hello, I am a message")
	sigma, _ := auth(rand.Reader, testPubA, testPubB, testPubC, testSec, message)
	ver := verify(testPubA, testPubB, testPubC, sigma, message)
	c.Assert(ver, Equals, true)

	fakeMessage := []byte("fake message")
	ver = verify(testPubA, testPubB, testPubC, sigma, fakeMessage)
	c.Assert(ver, Equals, false)

	ver = verify(testPubB, testPubB, testPubC, sigma, message)
	c.Assert(ver, Equals, false)

	ver = verify(testPubA, testPubA, testPubC, sigma, message)
	c.Assert(ver, Equals, false)

	ver = verify(testPubA, testPubB, testPubB, sigma, message)
	c.Assert(ver, Equals, false)

	ver = verify(testPubA, testPubB, testPubC, testSigma, message)
	c.Assert(ver, Equals, false)
}
