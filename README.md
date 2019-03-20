# understanding_golang

#### 00 - Understand package and import

- package main:
  - package = collection of common source code
  - main = secret name it means that file will be runnable
- 2 types of packages:
  - executable - must have func called 'main'
  - reusable - defines a package that can be used as a dependency
- import:
  - import standard libs (fmt) and reusable packages

#### 01 - Card project

- define variables
- function
- array and slice
- for loop
  - array vs slice:
    - Array: fixed length list of things
    - Slice: array that cang grow and shrink
    - slice
      - cards := []string{"Ace", newCard()}
- declaring custom types
  ```
  type deck []string
  ```
- add method to custom type (receiver functions)

  ```
  func (d deck) print() {}

  // (d deck) - type on which this method will be attached to
  // d - working variable (initialized variable on which we're working on) (keyword 'this' in Javascript)
  // print() - function name
  ```
