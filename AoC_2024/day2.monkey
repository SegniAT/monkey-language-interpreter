let content = trimSuffix(readFile("input.txt"), "\n")
let lines = splitString(content, "\n")

let prep = fn(i) {
	if (i==len(lines)){
		return {
			"colOne": [],
			"colTwo": [],
		};
	}

	let res = prep(i+1);

	let splitLine = splitString(lines[i], "   ");
	return {
		"colOne": push(res["colOne"], atoi(splitLine[0])),
		"colTwo":push(res["colTwo"], atoi(splitLine[1])),
	};
}

let res = prep(0)

let colOne = sortInts(res["colOne"]);
let colTwo = sortInts(res["colTwo"]);

let countNumInArray = fn(num, arr, i) {
	if (i == len(arr)){
		return 0;
	}

	if (num < arr[i]){
		return 0;
	}

	let res = countNumInArray(num, arr, i+1);
	return if (num == arr[i]){ res + 1 } else { res };
}

let dayTwo = fn(i){
	if (i == len(colOne)){
		return 0;
	}

	let sol = dayTwo(i+1);
	
	let num = colOne[i]
	let numCount = countNumInArray(num, colTwo, 0);

	return sol + (num * numCount);
}

puts(dayTwo(0))
