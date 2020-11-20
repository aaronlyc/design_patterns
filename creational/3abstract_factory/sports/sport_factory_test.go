package sports

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SportFactory", func() {
	It("测试正确的情况: ", func() {
		adidasFactory, err1 := getSprotsFactory("adidas")
		Expect(err1).Should(BeNil(), "不支持的类型")

		nikeFactory, err2 := getSprotsFactory("nike")
		Expect(err2).Should(BeNil(), "不支持的类型")

		adidasShoe := adidasFactory.makeShoe()
		adidasShirt := adidasFactory.makeShirt()

		nikeShoe := nikeFactory.makeShoe()
		nikeShirt := nikeFactory.makeShirt()

		printShirtDetails(nikeShirt)
		printShoeDetails(nikeShoe)

		printShirtDetails(adidasShirt)
		printShoeDetails(adidasShoe)
	})

	It("测试不正确的情况: ", func() {
		adidasFactory, err1 := getSprotsFactory("adiXXXXdas")   //输入错误的类型
		Expect(err1).Should(BeNil(), "不支持的类型")

		nikeFactory, err2 := getSprotsFactory("nike")
		Expect(err2).Should(BeNil(), "不支持的类型")

		adidasShoe := adidasFactory.makeShoe()
		adidasShirt := adidasFactory.makeShirt()

		nikeShoe := nikeFactory.makeShoe()
		nikeShirt := nikeFactory.makeShirt()

		printShirtDetails(nikeShirt)
		printShoeDetails(nikeShoe)

		printShirtDetails(adidasShirt)
		printShoeDetails(adidasShoe)
	})
})

func printShoeDetails(s iShoe) {
	fmt.Printf("Shoe Logo: %s\n", s.getLogo())
	fmt.Printf("Shoe Size: %d\n", s.getSize())
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Shirt Logo: %s\n", s.getLogo())
	fmt.Printf("Shirt Size: %d\n", s.getSize())
}