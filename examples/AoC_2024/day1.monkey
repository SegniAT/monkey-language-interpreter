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

let dayOne = fn(i, j){
	if (i==len(lines)) {
		return 0
	}

	let sum = dayOne(i+1, j+1);

	let diff = colOne[i] - colTwo[j];
	return sum + if(diff<0){ -1*diff }else{ diff };
}

puts(dayOne(0, 0))

