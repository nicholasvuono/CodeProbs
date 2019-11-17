import java.util.*;

public class DayOfWeek {
    
    public static String solution(String S, int k) {
        ArrayList<String> days = new ArrayList<>(
            Arrays.asList("Mon",
                          "Tue",
                          "Wed",
                          "Thu",
                          "Fri",
                          "Sat",
                          "Sun"));
        return days.get((days.indexOf(S) + k) % 7);
    }

    public static void main(String[] args) {
        System.out.println(solution("Wed", 2));
        System.out.println(solution("Sat", 23));
    }

}