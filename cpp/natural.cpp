#include <iostream>

void natural(int n) {
	if (n < 10) {
		std::cout << static_cast<char>('0' + n);
	} else {
		natural(n / 10);
		std::cout <<  static_cast<char>('0' + (n % 10));
	}
}

int main() {
	std::cout << "Print 376: ";
	natural(376);
	std::cout << std::endl;
}
