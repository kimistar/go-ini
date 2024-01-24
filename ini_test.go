package ini

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	TestIni     = "testdata/test.ini"
	NotFoundIni = "404.ini"
	MoreIni     = "testdata/more.ini"

	DevelopEnv = "development"
	TestingEnv = "testing"
	ProductEnv = "production"
)

func TestLoad(t *testing.T) {
	Convey("Test load", t, func() {
		Convey("ini file not found", func() {
			_, err := Load(NotFoundIni)
			So(err, ShouldNotBeNil)
		})
		Convey("ini file found", func() {
			_, err := Load(TestIni)
			So(err, ShouldBeNil)
		})
	})
}

func TestConfig_Read(t *testing.T) {
	Convey("Test read", t, func() {
		cfg, _ := Load(TestIni, MoreIni)

		Convey("comment begins with # or ;", func() {
			v1 := cfg.Read(DevelopEnv, "key2")
			So(v1, ShouldEqual, "")
			v2 := cfg.Read(DevelopEnv, "key3")
			So(v2, ShouldEqual, "")
		})
		Convey("a line doesnot contain =", func() {
			v := cfg.Read(DevelopEnv, "key4")
			So(v, ShouldEqual, "")
		})
		Convey("a key's value is empty", func() {
			v := cfg.Read(DevelopEnv, "key5")
			So(v, ShouldEqual, "")
		})
		Convey(`comment with substr '\t#'`, func() {
			v := cfg.Read(DevelopEnv, "key6")
			So(v, ShouldEqual, "key6")
		})
		Convey(`comment with substr ' #'`, func() {
			v := cfg.Read(DevelopEnv, "key7")
			So(v, ShouldEqual, "key7")
		})
		Convey(`comment with substr ' ;'`, func() {
			v := cfg.Read(DevelopEnv, "key8")
			So(v, ShouldEqual, "key8")
		})
		Convey(`comment with substr '\t//'`, func() {
			v := cfg.Read(DevelopEnv, "key9")
			So(v, ShouldEqual, "key9")
		})
		Convey(`comment with substr ' //'`, func() {
			v := cfg.Read(DevelopEnv, "key10")
			So(v, ShouldEqual, "key10")
		})
		Convey("normal ", func() {
			v := cfg.Read(DevelopEnv, "key11")
			So(v, ShouldEqual, "key11#$*()@###")
		})
		Convey("different value with same key in different sections", func() {
			v1 := cfg.Read(DevelopEnv, "key1")
			So(v1, ShouldEqual, "development")
			v2 := cfg.Read(TestingEnv, "key1")
			So(v2, ShouldEqual, "testing")
			v3 := cfg.Read(ProductEnv, "key1")
			So(v3, ShouldEqual, "production")
		})
		Convey("values in multiple files", func() {
			v1 := cfg.Read(DevelopEnv, "key1")
			So(v1, ShouldEqual, "development")
			v2 := cfg.Read(DevelopEnv, "name")
			So(v2, ShouldEqual, "star")
			v3 := cfg.Read(TestingEnv, "name")
			So(v3, ShouldEqual, "kimi")
			v4 := cfg.Read(ProductEnv, "name")
			So(v4, ShouldEqual, "Kimi.Wang")
		})
	})
}
