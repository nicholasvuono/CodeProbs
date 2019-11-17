function solution(S, k) {
    let days = new Array("Mon",
                         "Tue",
                         "Wed",
                         "Thu",
                         "Fri",
                         "Sat",
                         "Sun");
    return days[(days.indexOf(S) + k) % 7];
}

console.log(solution("Wed", 2));
console.log(solution("Sat", 23));