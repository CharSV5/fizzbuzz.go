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
  Context("When i is divisible by 15", func() {
    It("returns true", func() {
      Expect(isDivisibleBy(30, 15)).Should(BeTrue())
    })
  })
})
