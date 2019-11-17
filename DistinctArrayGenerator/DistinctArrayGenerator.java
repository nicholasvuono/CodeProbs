import java.util.Arrays;

public class DistinctArrayGenerator {

    public static int[] solution(int N) {
        int[] array = new int[N];
        int total = 0;
        for(int i = 0; i < N-1; i++) {
            array[i] = i;
            total = total + i;
        }
        array[N-1] = -(total);
        return array;
    }

    public static void main(String[] args) {
        System.out.println(Arrays.toString(solution(4)));
        System.out.println("Sum: " + Arrays.stream(solution(4)).sum());
        System.out.println(Arrays.toString(solution(3)));
        System.out.println("Sum: " + Arrays.stream(solution(3)).sum());
    }
}