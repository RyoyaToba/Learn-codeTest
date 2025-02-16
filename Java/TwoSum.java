import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> numMap = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int complement = target - nums[i];

            // マップに補数があるか確認
            if (numMap.containsKey(complement)) {
                return new int[]{numMap.get(complement), i};
            }

            // マップに現在の値を登録
            numMap.put(nums[i], i);
        }

        return new int[]{}; // 解が必ずある前提なので、ここは実行されない
    }
}
