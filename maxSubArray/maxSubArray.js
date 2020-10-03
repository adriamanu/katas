function maxSequence(arr) {
  if (arr.length === 0) {
    return 0;
  }

  let current = 0;
  let max = 0;

  arr.forEach((element) => {
    current = Math.max(element, current + element);
    max = Math.max(max, current);
  });
  return max;
}

console.log(maxSequence([-2, 1, -3, 4, -1, 2, 1, -5, 4]));
console.log(maxSequence([]));
