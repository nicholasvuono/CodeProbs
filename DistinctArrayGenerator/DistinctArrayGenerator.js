function solution(N) {
    let array = new Array(N);
    let total = 0;
    for(var i = 0; i < N-1; i++){
        array[i] = i;
        total = total + i;
    }
    array[N-1] = -(total);
    return array;
}

console.log(solution(4));
console.log("Sum: " + solution(4).reduce((accumulator, currentValue)=>accumulator+currentValue,0));
console.log(solution(3));
console.log("Sum: " + solution(3).reduce((accumulator, currentValue)=>accumulator+currentValue,0));