#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>

int solve(const std::vector<int>& nums) {
    std::unordered_map<int, int> last_seen;
    last_seen.reserve(nums.size());
    int max_length = 0;
    int left = 0;

    for (int right = 0; right < static_cast<int>(nums.size()); ++right) {
        int num = nums[right];
        auto it = last_seen.find(num);
        if (it != last_seen.end() && it->second >= left) {
            left = it->second + 1;
        }

        last_seen[num] = right;
        max_length = std::max(max_length, right - left + 1);
    }

    return max_length;
}

int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);

    int n;
    std::cin >> n;

    std::vector<int> nums(n);
    for (int& num : nums) {
        std::cin >> num;
    }

    std::cout << solve(nums) << '\n';
}
