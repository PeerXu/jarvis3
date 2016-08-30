package user

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestGenSalt(t *testing.T) {
	c.Convey("salts should not equal", t, func() {
		s1 := genSalt()
		s2 := genSalt()
		c.So(s1, c.ShouldNotEqual, s2)
	})
}

func TestNewEncryptedPasswordNotEqualWithDifferenceSalt(t *testing.T) {
	c.Convey("encrypted passwords should not equal", t, func() {
		s := genSalt()
		p1 := NewEncryptedPasswordWithSalt("hello", s)
		p2 := NewEncryptedPasswordWithSalt("world", s)
		c.So(p1.equal(p2), c.ShouldBeFalse)
	})
}

func TestValicatePassword(t *testing.T) {
	c.Convey("valicating password", t, func() {
		passwd := NewEncryptedPassword("hello")
		c.Convey("with same password should be equal", func() {
			c.So(passwd.Validate("hello"), c.ShouldBeTrue)
		})

		c.Convey("with difference password should be not equal", func() {
			c.So(passwd.Validate("world"), c.ShouldBeFalse)
		})
	})
}
