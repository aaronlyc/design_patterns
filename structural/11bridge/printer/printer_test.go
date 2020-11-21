package printer

import (
	. "github.com/onsi/ginkgo"
	"time"
)

var _ = Describe("Printer", func() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	winComputer := &windows{}

	It("测试mac使用hp打印:", func() {
		macComputer.setPrinter(hpPrinter)
		macComputer.print()
		time.Sleep(10 * time.Second)
	}, 5)

	It("测试mac使用epson打印:", func() {
		macComputer.setPrinter(epsonPrinter)
		macComputer.print()
	})

	It("测试windows使用hp打印:", func() {
		winComputer.setPrinter(hpPrinter)
		winComputer.print()
	})

	It("测试windows使用epson打印:", func() {
		winComputer.setPrinter(epsonPrinter)
		winComputer.print()
	})
})
