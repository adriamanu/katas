function iqTest(numbers) {
  const numberArray = numbers
    .split(" ")
    .map((numberAsString) => parseInt(numberAsString));
  const evenNumbers = numberArray
    .filter((number) => number % 2 === 0)
    .map((number) => number);
  const oddNumbers = numberArray
    .filter((number) => number % 2 !== 0)
    .map((number) => number);

  return evenNumbers.length > oddNumbers.length
    ? numberArray.findIndex((number) => number % 2 !== 0) + 1
    : numberArray.findIndex((number) => number % 2 === 0) + 1;
}

console.log(iqTest("2 4 7 8 10"));
console.log(iqTest("1 2 2"));
