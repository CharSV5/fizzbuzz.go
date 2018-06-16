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
      Expect(isDivisibleBy(30, 15)).Should(BeTrue())
    })
    It("returns false if not divisible", func() {
      Expect(isDivisibleBy(29, 15)).Should(BeFalse())
    })
  })
})
