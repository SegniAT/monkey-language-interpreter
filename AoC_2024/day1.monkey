let content = trimSuffix(readFile("input.txt"), "\n")
let lines = splitString(content, "\n")

let prep = fn(i, colOne, colTwo) {
	if (i==len(lines)){
		return {
			"colOne": sortInts(colOne),
			"colTwo": sortInts(colTwo),
		};
	}

	let splitLine = splitString(lines[i], "   ");
	let newColOne = push(colOne, atoi(splitLine[0]))
	let newColTwo = push(colTwo, atoi(splitLine[1]))

	prep(i+1, newColOne, newColTwo);
}

let res = prep(0, [], [])

let dayOne = fn(i, j, colOne, colTwo, sum){
	if (i==len(lines)) {
		return sum
	}

	let diff = colOne[i] - colTwo[j];
	let newSum = sum + if(diff<0){ -1*diff }else{ diff };
	return dayOne (i+1, j+1, colOne, colTwo, newSum);
}

let sol = dayOne(0, 0, res["colOne"], res["colTwo"], 0)
puts(sol)

