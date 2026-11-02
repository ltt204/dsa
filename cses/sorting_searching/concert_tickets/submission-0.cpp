#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

vector<int> solve(int n, int m, vector<int> prices, vector<int> members) {
    vector<int> result;
    vector<bool> sold(prices.size(), false);

    sort(prices.begin(), prices.end());

    for (int i = 0; i < m; i++) {
        int index = static_cast<int>(upper_bound(
            prices.begin(), prices.end(), members[i]) - prices.begin());

        while (index >= 0) {
            if (index >= n) {
                --index;
                continue;
            }

            if (!sold[index] && prices[index] <= members[i]) {
                result.push_back(prices[index]);
                sold[index] = true;
                break;
            }

            --index;
        }

        if (index == -1) {
            result.push_back(-1);
        }
    }
    
    return result;
}

int main() {
    int n, m;
    cin >> n >> m;

    vector<int> a(n);
    for (int i = 0; i < n; ++i) {
        cin >> a[i];
    }

    vector<int> b(m);
    for (int i = 0; i < m; ++i) {
        cin >> b[i];
    }

    vector<int> result = solve(n, m, a, b);

    for (int price : result) {
        cout << price << '\n';
    }

    return 0;
}