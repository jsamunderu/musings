#include <iostream>

int gcd_mod(int a, int b) {
	auto r = std::max(a, b) % std::min(a, b);
	if (r == 0) {
		return std::min(a, b);
	}
	return gcd_mod(std::min(a, b), r);
}

int gcd_sub(int a, int b) {
	auto r = std::abs(a - b);
	if (r == 0) {
		return std::min(a, b);
	}
	return gcd_sub(std::min(a, b), r);
}

int main() {
	std::cout << "gcd_mod(3, 5)=" << gcd_mod(5, 3) << std::endl;
	std::cout << "gcd_sub(3, 5)=" << gcd_sub(3, 5) << std::endl;
}
