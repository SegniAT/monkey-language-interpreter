let newAdder = fn(x) {
  return fn(y) { return x + y; };
};
let addFive = newAdder(5);
let sumTen = addFive(10);

let map = fn(arr, f) {
  let iter = fn(arr, acc) {
    if (len(arr) == 0) {
      return acc;
    }
    return iter(rest(arr), push(acc, f(first(arr))));
  };
  return iter(arr, []);
};
let doubled = map([1, 2, 3, 4], fn(x) { x * 2; });

let data = {"users": [{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}]};
let aliceName = data["users"][0]["name"];
puts(aliceName);
puts(len(doubled));

let MAX = fn(_, f) { f(_); };
let config = {foo: 1, "bar": 2};
