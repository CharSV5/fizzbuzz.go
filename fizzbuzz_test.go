package main

import (
    "testing"

    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

func TestCart(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "FizzBuzz")
}

var _ = Describe("isDivisibleBy", func() {
  Context("When i is divided by 15", func() {
    It("returns true if divisible by 15", func() {
      Expect(isDivisibleBy(30, 15)).To(BeTrue())
    })
    It("returns false if not divisible by 15", func() {
      Expect(isDivisibleBy(29, 15)).To(BeFalse())
    })
  })
  Context("When i is divided by 3", func() {
    It("returns true if divisible by 3", func() {
      Expect(isDivisibleBy(6, 3)).To(BeTrue())
    })
    It("returns false if not divisible by 3", func() {
      Expect(isDivisibleBy(7, 3)).To(BeFalse())
    })
  })
  Context("When i is divided by 5", func() {
    It("returns true if divisible by 5", func() {
      Expect(isDivisibleBy(10, 5)).To(BeTrue())
    })
    It("returns false if not divisible by 5", func() {
      Expect(isDivisibleBy(11, 5)).To(BeFalse())
    })
  })
})

// Describe("main", func() {
//   Context("When ran", func() {
//     It("replaces numbers divisible by 15, 3 and 5 with FizzBuzz, Fizz and Buzz", func() {
//       Expect(main())Should
//     })
//   })
// })
